package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorType string

const (
	NotFoundError       ErrorType = "NOT FOUND"
	NotAvailableError   ErrorType = "NOT AVAILABLE"
	DataExistError      ErrorType = "DATA EXIST"
	BadRequestError     ErrorType = "BAD REQUEST"
	InternalServerError ErrorType = "INTERNAL SERVER ERROR"
)

func NewCustomError(errorType ErrorType) error {
	return errors.New(string(errorType))
}

func CustomErrorHandler(c *gin.Context, err error) {
	errString := err.Error()

	var httpCode int
	var message string

	switch errString {
	case string(NotFoundError):
		httpCode = http.StatusNotFound
		message = string(NotFoundError)
		break
	case string(NotAvailableError):
		httpCode = http.StatusNotAcceptable
		message = string(NotAvailableError)
		break
	case string(BadRequestError):
		httpCode = http.StatusBadRequest
		message = string(BadRequestError)
		break
	case string(DataExistError):
		httpCode = http.StatusBadRequest
		message = string(DataExistError)
		break
	case string(InternalServerError):
	default:
		httpCode = http.StatusInternalServerError
		message = string(InternalServerError)
	}

	c.AbortWithStatusJSON(httpCode, gin.H{
		"msg": message,
	})
}
