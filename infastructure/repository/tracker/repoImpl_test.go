package tracker

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

func TestGetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	testCases := []struct {
		name     string
		username string
		mock     func()
		want     *model.TrackerEntity
		wantErr  bool
	}{
		{
			name:     "Given valid request when GetByUsername then return success",
			username: "username",
			mock: func() {
				mock.ExpectQuery("SELECT utc.coin_id, utc.coin_name, utc.coin_symbol FROM users u " +
					"JOIN user_tracked_coins utc ON u.id = utc.user_id WHERE u.username = ?").
					WithArgs("username").
					WillReturnRows(sqlmock.NewRows([]string{"coin_id", "coin_name", "coin_symbol"}).
						AddRow("bitcoin", "Bitcoin", "BTC").
						AddRow("bitcoin-cash", "Bitcoin Cash", "BTH"))
			},
			wantErr: false,
		},
		{
			name:     "Given error when GetByUsername then return error",
			username: "username",
			mock: func() {
				mock.ExpectQuery("SELECT utc.coin_id, utc.coin_name, utc.coin_symbol FROM users u " +
					"JOIN user_tracked_coins utc ON u.id = utc.user_id WHERE u.username = ?").
					WithArgs("username").
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewTrackerRepository(db)

			tc.mock()
			_, err := repo.GetByUsername(tc.username)
			if err != nil && !tc.wantErr {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	req := model.TrackerEntity{
		UserID: 1,
		Coin: model.Coin{
			CoinID:     "bitcoin",
			CoinSymbol: "BTC",
			CoinName:   "Bitcoin",
		},
	}

	testCases := []struct {
		name    string
		request *model.TrackerEntity
		mock    func()
		want    *model.TrackerEntity
		wantErr bool
	}{
		{
			name:    "Given valid request when Create then return success",
			request: &req,
			mock: func() {
				mock.ExpectExec("INSERT INTO user_tracked_coins (user_id, coin_id, coin_symbol, coin_name) VALUES (?, ?, ?, ?)").
					WithArgs(req.UserID, req.CoinID, req.CoinSymbol, req.CoinName).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name:    "Given error when Create then return error",
			request: &req,
			mock: func() {
				mock.ExpectExec("INSERT INTO user_tracked_coins (user_id, coin_id, coin_symbol, coin_name) VALUES (?, ?, ?, ?)").
					WithArgs(req.UserID, req.CoinID, req.CoinSymbol, req.CoinName).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewTrackerRepository(db)

			tc.mock()
			err := repo.Create(tc.request)
			if err != nil && !tc.wantErr {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	testCases := []struct {
		name    string
		userID  int
		coinID  string
		mock    func()
		want    *model.TrackerEntity
		wantErr bool
	}{
		{
			name:   "Given valid request when Delete then return success",
			userID: 1,
			coinID: "bitcoin",
			mock: func() {
				mock.ExpectExec("DELETE FROM user_tracked_coins WHERE user_id = ? AND coin_id = ?").
					WithArgs(1, "bitcoin").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name:   "Given error when Delete then return error",
			userID: 1,
			coinID: "bitcoin",
			mock: func() {
				mock.ExpectExec("DELETE FROM user_tracked_coins WHERE user_id = ? AND coin_id = ?").
					WithArgs(1, "bitcoin").
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewTrackerRepository(db)

			tc.mock()
			err := repo.Delete(tc.userID, tc.coinID)
			if err != nil && !tc.wantErr {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestGetByUserIDAndCoinID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	testCases := []struct {
		name    string
		userID  int
		coinID  string
		mock    func()
		want    *model.TrackerEntity
		wantErr bool
	}{
		{
			name:   "Given valid request when GetByUserIDAndCoinID then return success",
			coinID: "bitcoin",
			userID: 1,
			mock: func() {
				mock.ExpectQuery("SELECT utc.id, utc.user_id, utc.coin_id, utc.coin_name, utc.coin_symbol FROM user_tracked_coins utc WHERE utc.user_id = ? AND utc.coin_id = ?").
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "coin_id", "coin_name", "coin_symbol"}).
						AddRow("1", "1", "bitcoin", "Bitcoin", "BTC"))
			},
			wantErr: false,
		},
		{
			name:   "Given error when GetByUserIDAndCoinID then return error",
			coinID: "bitcoin",
			userID: 1,
			mock: func() {
				mock.ExpectQuery("SELECT utc.id, utc.user_id, utc.coin_id, utc.coin_name, utc.coin_symbol FROM user_tracked_coins utc WHERE utc.user_id = ? AND utc.coin_id = ?").
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewTrackerRepository(db)

			tc.mock()
			_, err := repo.GetByUserIDAndCoinID(tc.userID, tc.coinID)
			if err != nil && !tc.wantErr {
				t.Fatalf(err.Error())
			}
		})
	}
}
