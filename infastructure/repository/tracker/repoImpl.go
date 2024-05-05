package tracker

import (
	"database/sql"

	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

type trackerRepoImpl struct {
	db *sql.DB
}

func NewTrackerRepository(db *sql.DB) TrackerRepository {
	return &trackerRepoImpl{
		db: db,
	}
}

func (t *trackerRepoImpl) GetUserTrackedCoins(username string) (*model.UserTrackedCoin, error) {
	userTrackedCoin := &model.UserTrackedCoin{}
	rows, err := t.db.Query("SELECT utc.coin_id, utc.coin_name, utc.coin_symbol FROM users u "+
		"JOIN user_tracked_coins utc ON u.id = utc.user_id WHERE u.username = ?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	trackedCoins := make([]model.Coin, 0)
	for rows.Next() {
		var trackedCoin model.Coin
		if err := rows.Scan(
			&trackedCoin.CoinID,
			&trackedCoin.CoinName,
			&trackedCoin.CoinSymbol); err != nil {
			return nil, err
		}
		trackedCoins = append(trackedCoins, trackedCoin)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	userTrackedCoin.Username = username
	userTrackedCoin.TrackedCoins = trackedCoins
	return userTrackedCoin, nil
}

func (t *trackerRepoImpl) Create(req *model.TrackerEntity) error {
	_, err := t.db.Exec("INSERT INTO user_tracked_coins (user_id, coin_id, coin_symbol, coin_name) VALUES (?, ?, ?, ?)",
		req.UserID, req.CoinID, req.CoinSymbol, req.CoinName)
	if err != nil {
		return err
	}

	return nil
}

func (t *trackerRepoImpl) GetList() (*[]model.TrackerEntity, error) {
	rows, err := t.db.Query("SELECT utc.id, utc.user_id, utc.coin_id, utc.coin_name, utc.coin_symbol "+
		"FROM user_tracked_coins utc")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	trackedCoins := make([]model.TrackerEntity, 0)
	for rows.Next() {
		var trackedCoin model.TrackerEntity
		if err := rows.Scan(
			&trackedCoin.ID,
			&trackedCoin.UserID,
			&trackedCoin.CoinID,
			&trackedCoin.CoinName,
			&trackedCoin.CoinSymbol); err != nil {
			return nil, err
		}
		trackedCoins = append(trackedCoins, trackedCoin)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &trackedCoins, nil
}
