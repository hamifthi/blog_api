package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Router interface {
	Get(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type Blog interface {
	Router
}

type User interface {
	Router
}

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ApiError) Error() string {
	return e.Message
}

func NotFoundError(message string) ApiError {
	return ApiError{
		Code:    http.StatusNotFound,
		Message: message + " not found",
	}
}

func InternalServerError(message string) ApiError {
	return ApiError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func BadRequestError(message string) ApiError {
	return ApiError{
		Code:    http.StatusBadRequest,
		Message: message + " bad request",
	}
}
