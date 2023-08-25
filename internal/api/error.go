package api

import "net/http"

type Error interface {
	StatusCode() int
	error
}

type baseError struct {
	Msg string `json:"error"`
}

func (e *baseError) Error() string {
	return e.Msg
}

type BadRequestError struct {
	baseError
}

func NewBadRequestError(msg string) *BadRequestError {
	return &BadRequestError{
		baseError: baseError{Msg: msg},
	}
}

func (e *BadRequestError) StatusCode() int {
	return http.StatusBadRequest
}

type NotFoundError struct {
	baseError
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{
		baseError: baseError{Msg: msg},
	}
}

func (e *NotFoundError) StatusCode() int {
	return http.StatusNotFound
}

type ConflictError struct {
	baseError
}

func NewConflictError(msg string) *ConflictError {
	return &ConflictError{
		baseError: baseError{Msg: msg},
	}
}

func (e *ConflictError) StatusCode() int {
	return http.StatusConflict
}
