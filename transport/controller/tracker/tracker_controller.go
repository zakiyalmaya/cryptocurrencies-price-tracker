package tracker

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/application/tracker"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

type TrackerController struct {
	trackerSvc tracker.TrackerService
}

func NewTrackerController(trackerSvc tracker.TrackerService) *TrackerController {
	return &TrackerController{
		trackerSvc: trackerSvc,
	}
}

func (t *TrackerController) GetUserTrackedList(c *gin.Context) {
	defer c.Request.Body.Close()

	req := &model.AssetRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse("Username not found in context"))
		return
	}

	res, err := t.trackerSvc.GetUserTrackedList(username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (t *TrackerController) AddUserTrackedCoin(c *gin.Context) {
	defer c.Request.Body.Close()

	req := &model.TrackerEntity{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse(err.Error()))
		return
	}

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse("UserID not found in context"))
		return
	}
	req.UserID = int(userID.(float64))

	err := t.trackerSvc.AddUserTrackedCoin(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.HTTPSuccessResponse(nil))
}

func (t *TrackerController) DeleteUserTrackedCoin(c *gin.Context) {
	defer c.Request.Body.Close()

	coinID := c.Param("coinId")

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse("UserID not found in context"))
		return
	}

	err := t.trackerSvc.DeleteUserTrackedCoin(int(userID.(float64)), coinID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.HTTPSuccessResponse(nil))

}

func (t *TrackerController) GetList(c *gin.Context) {
	defer c.Request.Body.Close()

	res, err := t.trackerSvc.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.HTTPSuccessResponse(res))
}
