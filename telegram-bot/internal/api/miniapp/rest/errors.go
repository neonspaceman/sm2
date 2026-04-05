package rest

import (
	"net/http"
)

type Details = map[string]any

type Option func(*HttpError)

type HttpError struct {
	StatusCode int     `json:"-"`
	Code       string  `json:"code,omitempty"`
	Message    string  `json:"message,omitempty"`
	Details    Details `json:"details,omitempty"`
}

func WithCode(code string) Option {
	return func(err *HttpError) {
		err.Code = code
	}
}

func WithMessage(msg string) Option {
	return func(err *HttpError) {
		err.Message = msg
	}
}

func WithDetails(details Details) Option {
	return func(err *HttpError) {
		err.Details = details
	}
}

func NewHttpError(statusCode int, opts ...Option) *HttpError {
	err := &HttpError{
		StatusCode: statusCode,
	}

	for _, opt := range opts {
		opt(err)
	}

	return err
}

func (e *HttpError) Error() string {
	return e.Message
}

var BadRequestErr = NewHttpError(http.StatusBadRequest, WithMessage("Bad request"))
