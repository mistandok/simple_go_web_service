package clients

import (
	"net/http"
)

type Controller struct {
	logger Logger
	logic  Logic
}

func (c Controller) SayHello(writer http.ResponseWriter, request *http.Request) {
	c.logger.Log("In SayHello")
	userID := request.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Write([]byte(message))
}

func NewController(logger Logger, logic Logic) Controller {
	return Controller{
		logger: logger,
		logic:  logic,
	}
}
