package controllers

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	testUtils "mutant-ms/utils/tests"
)

func TestHealthCheck_OK(t *testing.T) {
	uri := "/health-check"

	_, _, rec, c := testUtils.SetupServerTest(http.MethodGet, uri, bytes.NewReader([]byte{}))

	healthCheckController := NewHealthCheckController()
	healthCheckController.GetHealthCheck(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}
