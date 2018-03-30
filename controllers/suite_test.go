package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/mattn/go-vue-example/controllers"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type S struct {
	Server *echo.Echo
}

var _ = Suite(&S{})

func (s *S) SetUpSuite(c *C) {
	s.Server = echo.New()
	controllers.Setup(s.Server)
}

func (s *S) TearDownSuite(c *C) {
}

func (s *S) SetUpTest(c *C) {
}

func (s *S) TearDownTest(c *C) {
}

func (s *S) PerformRequest(method string, path string, params url.Values) *httptest.ResponseRecorder {
	paramsEncoded := params.Encode()
	reader := strings.NewReader(paramsEncoded)
	if method == "GET" || method == "HEAD" {
		path += "?" + paramsEncoded
	}

	request, err := http.NewRequest(method, path, reader)
	if err != nil {
		panic(err)
	}

	if method == "POST" || method == "PUT" {
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	response := httptest.NewRecorder()

	s.Server.ServeHTTP(response, request)
	return response
}

// vi:syntax=go
