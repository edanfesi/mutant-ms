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

	if checkStatusForbidden(err) {
		response.Code = http.StatusForbidden
		return response.Code, response
	}

	return response.Code, response
}

func checkStatusForbidden(err error) bool {
	listForbidden := map[string]struct{}{
		"forbidden": {},
	}

	return check(err.Error(), listForbidden)
}

func check(err string, mapErrors map[string]struct{}) bool {
	if _, ok := mapErrors[err]; ok {
		return ok
	}

	return false
}
