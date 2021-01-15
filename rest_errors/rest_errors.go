package rest_errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusBadRequest,
		Message: message,
		Error:   "bad request",
	}
}

func NewFoundError(msg string) error {
	return errors.New(msg)
}

func NewBadNotFoundError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusNotFound,
		Message: message,
		Error:   "not found",
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Status:  http.StatusInternalServerError,
		Message: message,
		Error:   "Internal server error",
	}

	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}

	return result
}
