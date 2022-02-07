package errors

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerError_BadRequest(t *testing.T) {
	err := fmt.Errorf("[%w]", assert.AnError)
	statusCode, errorStruct := HandlerError(err)

	assert.Equal(t, http.StatusBadRequest, errorStruct.Code)
	assert.Equal(t, http.StatusBadRequest, statusCode)
}

func TestHandlerError_Forbidden(t *testing.T) {
	err := fmt.Errorf("forbidden")
	statusCode, errorStruct := HandlerError(err)

	assert.Equal(t, http.StatusForbidden, errorStruct.Code)
	assert.Equal(t, http.StatusForbidden, statusCode)
}
