package errors

import (
	"fmt"
	"net/http"
	"strings"
)

type errorResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func HandlerError(err error) (int, errorResponse) {
	message := strings.Builder{}
	message.WriteString(fmt.Sprint(err.Error()))

	response := errorResponse{
		Message: message.String(),
		Code:    http.StatusBadRequest,
	}

	return response.Code, response
}
