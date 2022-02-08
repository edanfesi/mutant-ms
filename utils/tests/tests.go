package tests

import (
	"io"
	"mutant-ms/constants"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func SetupServerTest(method string, url string, payload io.Reader) (*echo.Echo, *http.Request, *httptest.ResponseRecorder, echo.Context) {
	sw := echo.New()

	req := httptest.NewRequest(
		method,
		url,
		payload,
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(constants.Commons.Headers.ContentType, constants.Commons.Headers.ApplicationJSON)
	rec := httptest.NewRecorder()

	c := sw.NewContext(req, rec)

	return sw, req, rec, c
}
