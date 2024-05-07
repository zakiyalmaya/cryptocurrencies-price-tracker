package user

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/repository"
	mocksRepo "github.com/zakiyalmaya/cryptocurrencies-price-tracker/mocks/infrastructure/repository"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

var (
	mockUserRepo    *mocksRepo.MockUserRepository
	userSvc         UserService
	mockRedisClient *redis.Client
)

func Setup(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockRedisServer, err := miniredis.Run()
	if err != nil {
		t.Fatalf(err.Error())
	}

	mockRedisClient = redis.NewClient(&redis.Options{
		Addr: mockRedisServer.Addr(),
	})

	mockUserRepo = mocksRepo.NewMockUserRepository(mockCtl)
	userSvc = NewUserService(&repository.Repositories{
		Redcl: mockRedisClient,
		User:  mockUserRepo,
	})
}

func TestRegister(t *testing.T) {
	Setup(t)

	req := &model.UserEntity{
		Name:     "John Doe",
		Username: "john.doe",
		Email:    "john@example.com",
		Password: "John123!",
	}

	testCases := []struct {
		name    string
		request *model.UserEntity
		mock    func()
		wantErr bool
	}{
		{
			name:    "Given valid request when register user then return success",
			request: req,
			mock: func() {
				mockUserRepo.EXPECT().Create(req).Return(nil)
			},
			wantErr: false,
		},
		{
			name:    "Given error when register user then return error",
			request: req,
			mock: func() {
				mockUserRepo.EXPECT().Create(req).Return(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			err := userSvc.Register(tc.request)
			if err != nil && !tc.wantErr {
				t.Fatal(err)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	Setup(t)

	req := &model.AuthRequest{
		Username: "john.doe",
		Password: "John123!",
	}

	testCases := []struct {
		name    string
		request *model.AuthRequest
		mock    func()
		wantErr bool
	}{
		{
			name:    "Given valid request when login then return success",
			request: req,
			mock: func() {
				mockUserRepo.EXPECT().Get(req.Username).Return(&model.UserEntity{
					Username: req.Username,
					Password: req.Password,
				}, nil)
			},
			wantErr: false,
		},
		{
			name:    "Given invalid username when login then return error",
			request: req,
			mock: func() {
				mockUserRepo.EXPECT().Get(req.Username).Return(nil, sql.ErrNoRows)
			},
			wantErr: true,
		},
		{
			name:    "Given error get data user when login then return error",
			request: req,
			mock: func() {
				mockUserRepo.EXPECT().Get(req.Username).Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
		{
			name:    "Given invalid password when login then return error",
			request: req,
			mock: func() {
				mockUserRepo.EXPECT().Get(req.Username).Return(&model.UserEntity{
					Username: req.Username,
					Password: "123",
				}, nil)
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			_, err := userSvc.Login(tc.request)
			if err != nil && !tc.wantErr {
				t.Fatal(err)
			}
		})
	}
}

func TestLogout(t *testing.T) {
	Setup(t)

	testCases := []struct {
		name     string
		username string
		wantErr  bool
	}{
		{
			name:     "Given valid request when logout then return success",
			username: "username",
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := userSvc.Logout(tc.username)
			if err != nil && !tc.wantErr {
				t.Fatal(err)
			}
		})
	}
}
