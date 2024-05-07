package tracker

import (
	"database/sql"
	"log"

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

func (t *trackerRepoImpl) GetByUsername(username string) (*model.UserTrackedCoin, error) {
	userTrackedCoin := &model.UserTrackedCoin{}
	rows, err := t.db.Query("SELECT utc.coin_id, utc.coin_name, utc.coin_symbol FROM users u "+
		"JOIN user_tracked_coins utc ON u.id = utc.user_id WHERE u.username = ?", username)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}
	defer rows.Close()

	trackedCoins := make([]*model.Coin, 0)
	for rows.Next() {
		var trackedCoin model.Coin
		if err := rows.Scan(
			&trackedCoin.CoinID,
			&trackedCoin.CoinName,
			&trackedCoin.CoinSymbol); err != nil {
			return nil, err
		}
		trackedCoins = append(trackedCoins, &trackedCoin)
	}

	if err := rows.Err(); err != nil {
		log.Println("errorRepository: ", err.Error())
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
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}

func (t *trackerRepoImpl) Delete(userID int, coinID string) error {
	_, err := t.db.Exec("DELETE FROM user_tracked_coins WHERE user_id = ? AND coin_id = ?", userID, coinID)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}

func (t *trackerRepoImpl) GetByUserIDAndCoinID(userID int, coinID string) (*model.TrackerEntity, error) {
	trackedCoin := &model.TrackerEntity{}
	resp := t.db.QueryRow("SELECT utc.id, utc.user_id, utc.coin_id, utc.coin_name, utc.coin_symbol " +
		"FROM user_tracked_coins utc WHERE utc.user_id = ? AND utc.coin_id = ?", userID, coinID)

	if err := resp.Scan(
		&trackedCoin.ID,
		&trackedCoin.UserID,
		&trackedCoin.CoinID,
		&trackedCoin.CoinName,
		&trackedCoin.CoinSymbol); err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	return trackedCoin, nil
}
