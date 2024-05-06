package rest_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorInterface interface {
	GetCode() int
	GetMessage() string
	GetDescription() string
}

type Error struct {
	error
	code        int
	message     string
	description string
}

func (e *Error) GetCode() int {
	return e.code
}

func (e *Error) GetMessage() string {
	return e.message
}

func (e *Error) GetDescription() string {
	return e.description
}
func (e *Error) SetMessages(messages []string) {
	if len(messages) > 0 {
		e.message = messages[0]
	}
	if len(messages) > 1 {
		e.description = messages[1]
	}
}
func newError(statusCode int, err error, messages ...string) ErrorInterface {
	errorInstance := &Error{code: statusCode, error: err, description: http.StatusText(statusCode)}
	if len(messages) > 0 {
		errorInstance.SetMessages(messages)
	} else {
		errorInstance.SetMessages([]string{err.Error()})
	}
	return errorInstance
}

func NewUnauthorizedError(err error, messages ...string) ErrorInterface {
	return newError(http.StatusUnauthorized, err, messages...)
}

func NewNotFoundError(err error, messages ...string) ErrorInterface {
	return newError(http.StatusNotFound, err, messages...)
}

func NewBadRequestError(err error, messages ...string) ErrorInterface {
	return newError(http.StatusBadRequest, err, messages...)
}

func NewInternalServerError(err error, messages ...string) ErrorInterface {
	return newError(http.StatusInternalServerError, err, messages...)
}

func NewNotImplementedError(err error, messages ...string) ErrorInterface {
	return newError(http.StatusNotImplemented, err, messages...)
}
func ErrorHandler(err ErrorInterface, c *gin.Context) {
	c.AbortWithStatusJSON(err.GetCode(), gin.H{"message": err.GetMessage(), "description": err.GetDescription()})
}
