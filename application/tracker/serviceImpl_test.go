package tracker

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/repository"
	mocksClient "github.com/zakiyalmaya/cryptocurrencies-price-tracker/mocks/infrastructure/client"
	mocksRepo "github.com/zakiyalmaya/cryptocurrencies-price-tracker/mocks/infrastructure/repository"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

var (
	mockTrackerRepo     *mocksRepo.MockTrackerRepository
	mockCoinCapSvc      *mocksClient.MockCoinCapService
	mockExchangeRateSvc *mocksClient.MockExchangeRateService
	trackerSvc          TrackerService
)

func Setup(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockTrackerRepo = mocksRepo.NewMockTrackerRepository(mockCtl)
	mockCoinCapSvc = mocksClient.NewMockCoinCapService(mockCtl)
	mockExchangeRateSvc = mocksClient.NewMockExchangeRateService(mockCtl)

	trackerSvc = NewTrackerService(mockCoinCapSvc, mockExchangeRateSvc, &repository.Repositories{
		Tracker: mockTrackerRepo,
	})
}

func TestGetAssetList(t *testing.T) {
	Setup(t)
	request := &model.AssetRequest{}

	testCases := []struct {
		name    string
		request *model.AssetRequest
		mock    func()
		want    *model.AssetsResponse
		wantErr bool
	}{
		{
			name:    "Given valid request when GetAssetList then return success",
			request: request,
			mock: func() {
				mockCoinCapSvc.EXPECT().GetAssets(request).Return(&model.AssetsResponse{}, nil)
			},
			wantErr: false,
		},
		{
			name:    "Given error when GetAssetList then return error",
			request: request,
			mock: func() {
				mockCoinCapSvc.EXPECT().GetAssets(request).Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			_, err := trackerSvc.GetAssetList(tc.request)
			if err != nil && !tc.wantErr {
				t.Fatal(err)
			}
		})
	}

}

func TestDeleteUserTrackedCoin(t *testing.T) {
	Setup(t)

	testCases := []struct {
		name    string
		userID  int
		coinID  string
		mock    func()
		wantErr bool
	}{
		{
			name:   "Given valid request when DeleteUserTrackedCoin then return success",
			userID: 1,
			coinID: "bitcoin",
			mock: func() {
				mockTrackerRepo.EXPECT().Delete(1, "bitcoin").Return(nil)
			},
			wantErr: false,
		},
		{
			name:   "Given error when DeleteUserTrackedCoin then return error",
			userID: 1,
			coinID: "bitcoin",
			mock: func() {
				mockTrackerRepo.EXPECT().Delete(1, "bitcoin").Return(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			err := trackerSvc.DeleteUserTrackedCoin(tc.userID, tc.coinID)
			if err != nil && !tc.wantErr {
				t.Fatal(err)
			}
		})
	}

}

func TestAddUserTrackedCoin(t *testing.T) {
	Setup(t)

	req := &model.AddUserTrackedCoinRequest{
		UserID: 1,
		CoinID: "bitcoin",
	}

	testCases := []struct {
		name    string
		request *model.AddUserTrackedCoinRequest
		mock    func()
		wantErr bool
	}{
		{
			name:    "Given valid request when AddUserTrackedCoin then return success",
			request: req,
			mock: func() {
				mockCoinCapSvc.EXPECT().GetAsset(req.CoinID).Return(&model.AssetResponse{
					Data: model.Response{ID: req.CoinID}}, nil)
				mockTrackerRepo.EXPECT().GetByUserIDAndCoinID(req.UserID, req.CoinID).Return(nil, sql.ErrNoRows)
				mockTrackerRepo.EXPECT().Create(&model.TrackerEntity{
					UserID: req.UserID,
					Coin: model.Coin{
						CoinID: req.CoinID,
					},
				}).Return(nil)
			},
			wantErr: false,
		},
		{
			name:    "Given error when create tracker to db then return error",
			request: req,
			mock: func() {
				mockCoinCapSvc.EXPECT().GetAsset(req.CoinID).Return(&model.AssetResponse{
					Data: model.Response{ID: req.CoinID}}, nil)
				mockTrackerRepo.EXPECT().GetByUserIDAndCoinID(req.UserID, req.CoinID).Return(nil, sql.ErrNoRows)
				mockTrackerRepo.EXPECT().Create(&model.TrackerEntity{
					UserID: req.UserID,
					Coin: model.Coin{
						CoinID: req.CoinID,
					},
				}).Return(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name:    "Given coin is already exist when get by user id and coin id then return error",
			request: req,
			mock: func() {
				mockCoinCapSvc.EXPECT().GetAsset(req.CoinID).Return(&model.AssetResponse{
					Data: model.Response{ID: req.CoinID}}, nil)
				mockTrackerRepo.EXPECT().GetByUserIDAndCoinID(req.UserID, req.CoinID).Return(&model.TrackerEntity{}, nil)
			},
			wantErr: true,
		},
		{
			name:    "Given error when get by user id and coin id then return error",
			request: req,
			mock: func() {
				mockCoinCapSvc.EXPECT().GetAsset(req.CoinID).Return(&model.AssetResponse{
					Data: model.Response{ID: req.CoinID}}, nil)
				mockTrackerRepo.EXPECT().GetByUserIDAndCoinID(req.UserID, req.CoinID).Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
		{
			name:    "Given coin is not exist in asset list then return error",
			request: req,
			mock: func() {
				mockCoinCapSvc.EXPECT().GetAsset(req.CoinID).Return(&model.AssetResponse{
					Data: model.Response{ID: ""}}, nil)
			},
			wantErr: true,
		},
		{
			name:    "Given error when get asset by coin id then return error",
			request: req,
			mock: func() {
				mockCoinCapSvc.EXPECT().GetAsset(req.CoinID).Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			err := trackerSvc.AddUserTrackedCoin(tc.request)
			if err != nil && !tc.wantErr {
				t.Fatal(err)
			}
		})
	}

}

func TestGetUserTrackedList(t *testing.T) {
	Setup(t)

	testCases := []struct {
		name     string
		username string
		mock     func()
		wantErr  bool
	}{
		{
			name:     "Given valid request when GetUserTrackedList then return success",
			username: "username",
			mock: func() {
				mockTrackerRepo.EXPECT().GetByUsername("username").Return(&model.UserTrackedCoin{
					Username: "username",
					TrackedCoins: []*model.Coin{
						{CoinID: "bitcoin"},
						{CoinID: "ethereum"},
					},
				}, nil)

				ids := "bitcoin,ethereum"
				mockCoinCapSvc.EXPECT().GetAssets(&model.AssetRequest{
					Ids: &ids,
				}).Return(&model.AssetsResponse{
					Data: []model.Response{
						{ID: "bitcoin", PriceUsd: "1"},
						{ID: "ethereum", PriceUsd: "1"},
					},
				}, nil)

				rateIDR := float64(10)
				mockExchangeRateSvc.EXPECT().GetLatest("USD", "IDR").Return(rateIDR, nil)
			},
			wantErr: false,
		},
		{
			name:     "Given error when get exchange rate to client then return error",
			username: "username",
			mock: func() {
				mockTrackerRepo.EXPECT().GetByUsername("username").Return(&model.UserTrackedCoin{
					Username: "username",
					TrackedCoins: []*model.Coin{
						{CoinID: "bitcoin"},
						{CoinID: "ethereum"},
					},
				}, nil)

				ids := "bitcoin,ethereum"
				mockCoinCapSvc.EXPECT().GetAssets(&model.AssetRequest{
					Ids: &ids,
				}).Return(&model.AssetsResponse{
					Data: []model.Response{
						{ID: "bitcoin", PriceUsd: "1"},
						{ID: "ethereum", PriceUsd: "1"},
					},
				}, nil)

				mockExchangeRateSvc.EXPECT().GetLatest("USD", "IDR").Return(float64(0), errors.New("error"))
			},
			wantErr: true,
		},
		{
			name:     "Given error when get list asset to client then return error",
			username: "username",
			mock: func() {
				mockTrackerRepo.EXPECT().GetByUsername("username").Return(&model.UserTrackedCoin{
					Username: "username",
					TrackedCoins: []*model.Coin{
						{CoinID: "bitcoin"},
						{CoinID: "ethereum"},
					},
				}, nil)

				ids := "bitcoin,ethereum"
				mockCoinCapSvc.EXPECT().GetAssets(&model.AssetRequest{
					Ids: &ids,
				}).Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
		{
			name:     "Given empty user tracked coin then return success",
			username: "username",
			mock: func() {
				mockTrackerRepo.EXPECT().GetByUsername("username").Return(&model.UserTrackedCoin{
					Username: "username"}, nil)
			},
			wantErr: false,
		},
		{
			name:     "Given error when get user tracked coin to db then return error",
			username: "username",
			mock: func() {
				mockTrackerRepo.EXPECT().GetByUsername("username").Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			_, err := trackerSvc.GetUserTrackedList(tc.username)
			if err != nil && !tc.wantErr {
				t.Fatal(err)
			}
		})
	}
}
