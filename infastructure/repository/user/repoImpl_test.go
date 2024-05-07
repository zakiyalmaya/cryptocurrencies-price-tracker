package user

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	req := model.UserEntity{
		Name:     "John Doe",
		Email:    "john@example.com",
		Username: "john.doe",
		Password: "john123",
	}

	testCases := []struct {
		name    string
		request *model.UserEntity
		mock    func()
		wantErr bool
	}{
		{
			name:    "Given valid request when Create user then return success",
			request: &req,
			mock: func() {
				mock.ExpectExec("INSERT INTO users (name, email, username, password) VALUES (?, ?, ?, ?)").
					WithArgs(req.Name, req.Email, req.Username, req.Password).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name:    "Given error when Create user then return error",
			request: &req,
			mock: func() {
				mock.ExpectExec("INSERT INTO users (name, email, username, password) VALUES (?, ?, ?, ?)").
					WithArgs(req.Name, req.Email, req.Username, req.Password).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewUserRepository(db)

			tc.mock()
			err := repo.Create(tc.request)
			if err != nil && !tc.wantErr {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	testCases := []struct {
		name     string
		username string
		mock     func()
		want     *model.UserEntity
		wantErr  bool
	}{
		{
			name:     "Given valid request when Get user then return success",
			username: "username",
			mock: func() {
				mock.ExpectQuery("SELECT id, name, email, username, password FROM users WHERE username = ?").
					WithArgs("username").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "username", "password"}).
						AddRow("1", "John Doe", "john@example.com", "john.doe", "john123"))
			},
			wantErr: false,
		},
		{
			name:     "Given error when Get user then return error",
			username: "username",
			mock: func() {
				mock.ExpectQuery("SELECT id, name, email, username, password FROM users WHERE username = ?").
					WithArgs("username").
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewUserRepository(db)

			tc.mock()
			_, err := repo.Get(tc.username)
			if err != nil && !tc.wantErr {
				t.Fatalf(err.Error())
			}
		})
	}
}
