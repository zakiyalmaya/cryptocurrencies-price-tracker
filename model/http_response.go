package model

type ResponseSystem struct {
	ResponseMessage string      `json:"responseMessage"`
	Data            interface{} `json:"data,omitempty"`
}

func HTTPSuccessResponse(res interface{}) ResponseSystem {
	return ResponseSystem{
		ResponseMessage: "Success",
		Data:            res,
	}
}

func HTTPErrorResponse(errMsg string) ResponseSystem {
	return ResponseSystem{
		ResponseMessage: errMsg,
	}
}
