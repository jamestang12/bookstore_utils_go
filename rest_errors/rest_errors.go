package rest_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	status  int           `json:"status"`
	message string        `json:"message"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %y ]", e.message, e.status, e.error, e.causes)
}

func (e restErr) Message() string {
	return e.message
}
func (e restErr) Status() int {
	return e.status
}
func (e restErr) Causes() []interface{} {
	return e.causes
}

func NewRestErrpr(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{
		status:  status,
		message: message,
		error:   err,
		causes:  causes,
	}
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		status:  http.StatusBadRequest,
		message: message,
		error:   "bad request",
	}
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}

	return apiErr, nil
}

func NewFoundError(msg string) error {
	return errors.New(msg)
}

func NewBadNotFoundError(message string) RestErr {
	return restErr{
		status:  http.StatusNotFound,
		message: message,
		error:   "not found",
	}
}

func newUnauthorizedError(message string) RestErr {
	return restErr{
		message: "unable to retrieve user infornation from given access_token",
		status:  http.StatusUnauthorized,
		error:   "unauthorized",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		status:  http.StatusInternalServerError,
		message: message,
		error:   "Internal server error",
		causes:  []interface{}{err},
	}

	if err != nil {
		result.causes = append(result.causes, err.Error())
	}

	return result
}
