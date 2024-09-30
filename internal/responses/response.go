package responses

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func BadRequestResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, Response{
		Status:  http.StatusBadRequest,
		Message: "bad request",
		Data:    err.Error(),
	})
}

func InternalServerErrorResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, Response{
		Status:  http.StatusInternalServerError,
		Message: "internal server error",
		Data:    err.Error(),
	})
}
