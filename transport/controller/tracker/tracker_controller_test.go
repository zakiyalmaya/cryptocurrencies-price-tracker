package tracker

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mocksApp "github.com/zakiyalmaya/cryptocurrencies-price-tracker/mocks/application"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

var (
	mockTrackerSvc *mocksApp.MockTrackerService
	trackerCtrl    *TrackerController
)

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func Setup(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockTrackerSvc = mocksApp.NewMockTrackerService(mockCtl)

	trackerCtrl = NewTrackerController(mockTrackerSvc)
}

func TestGetUserTrackedCoins(t *testing.T) {
	Setup(t)

	testcases := []struct {
		name               string
		mock               func()
		expectedHttpStatus int
	}{
		{
			name: "Given valid request when hit endpoint GetUserTrackedCoin then return 200 OK",
			mock: func() {
				mockTrackerSvc.EXPECT().GetUserTrackedList("john.doe").Return(&model.UserTrackedCoin{
					Username: "john.doe",
				}, nil)
			},
			expectedHttpStatus: http.StatusOK,
		},
		{
			name: "Given error when hit GetUserTrackedCoin then return 500 Internal Server Error",
			mock: func() {
				mockTrackerSvc.EXPECT().GetUserTrackedList("john.doe").Return(nil, errors.New("error"))
			},
			expectedHttpStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			w := httptest.NewRecorder()
			ctx := GetTestGinContext(w)
			ctx.Request.Method = "GET"
			ctx.Request.Header.Set("Content-Type", "application/json")
			ctx.Set("username", "john.doe")

			trackerCtrl.GetUserTrackedList(ctx)

			if tc.expectedHttpStatus != w.Code {
				t.Errorf("error test case")
			}
		})
	}
}

func TestDeleteUserTrackedCoin(t *testing.T) {
	Setup(t)

	testcases := []struct {
		name               string
		url                string
		mock               func()
		expectedHttpStatus int
	}{
		{
			name: "Given invalid request with empty coinId when hit endpoint DeleteUserTrackedCoin then return 400 BadRequest",
			mock: func() {
				mockTrackerSvc.EXPECT().DeleteUserTrackedCoin(1, "bitcoin").Return(nil)
			},
			expectedHttpStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			w := httptest.NewRecorder()
			ctx := GetTestGinContext(w)
			ctx.Request.Method = "DELETE"
			ctx.Request.Header.Set("Content-Type", "application/json")
			ctx.Set("userId", "1")

			trackerCtrl.DeleteUserTrackedCoin(ctx)

			if tc.expectedHttpStatus != w.Code {
				t.Errorf("error test case")
			}
		})
	}
}