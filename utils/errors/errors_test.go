package errors

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerError(t *testing.T) {
	err := fmt.Errorf("[%w]", assert.AnError)
	statusCode, errorStruct := HandlerError(err)

	assert.Equal(t, http.StatusBadRequest, errorStruct.Code)
	assert.Equal(t, http.StatusBadRequest, statusCode)
}
