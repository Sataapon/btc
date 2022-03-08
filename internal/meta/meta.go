package meta

import "errors"

type Error struct {
	httpStatusCode int
	err            error
}

func NewError(httpStatusCode int, err error) *Error {
	return &Error{
		httpStatusCode: httpStatusCode,
		err:            err,
	}
}

func (e *Error) HTTPStatus() int {
	return e.httpStatusCode
}

func (e *Error) Error() string {
	return e.err.Error()
}

func IsError(err error) (*Error, bool) {
	var metaErr *Error
	if errors.As(err, &metaErr) {
		return metaErr, true
	}
	return nil, false
}
