package helpers

import (
	"fmt"
	"github.com/labstack/echo"
)

type ResponseError struct {
	Code    int
	Message string
}

type ResponseMap map[string]interface{}
type Responsible interface {
	ToResponseMap() ResponseMap
}

func NewResponseError(httpCode int, message string) *ResponseError {
	err := &ResponseError{Code: httpCode, Message: message}

	return err
}

func (err *ResponseError) Error() string {
	return fmt.Sprintf("Response error: %d %s", err.Code, err.Message)
}

func JSONResponseObject(c echo.Context, statusCode int, object Responsible) error {
	body := object.ToResponseMap()
	return JSONResponse(c, statusCode, body)
}

func JSONResponse(c echo.Context, statusCode int, body ResponseMap) error {
	if statusCode != 200 {
		body["success"] = false
	} else {
		body["success"] = true
	}
	return c.JSON(statusCode, body)
}

func JSONResponseArray(c echo.Context, statusCode int, collection []ResponseMap) error {
	body := ResponseMap{}
	body["items"] = collection
	body["count"] = len(collection)

	return JSONResponse(c, statusCode, body)
}

func JSONResponseError(c echo.Context, err *ResponseError) error {
	body := ResponseMap{"success": false, "message": err.Message}

	return JSONResponse(c, err.Code, body)
}

// vi:syntax=go
