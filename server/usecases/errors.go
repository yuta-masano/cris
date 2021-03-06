package usecases

import "errors"

type ClientError interface {
	IsClientError() bool
}

type errEmptyContent struct{ error }

func (e errEmptyContent) IsClientError() bool { return true }

type errTemperingDetected struct{ error }

func (e errTemperingDetected) IsClientError() bool { return true }

var (
	ErrEmptyContent      = errEmptyContent{errors.New("received content length is zero")}
	ErrTemperingDetected = errTemperingDetected{errors.New("tempering of the resource was detected")}
)
