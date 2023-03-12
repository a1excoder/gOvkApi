package gOvkApi

import (
	"errors"
)

func (err *ErrorReturned) GetError() (bool, error) {
	if err.Err != nil {
		return true, err.Err
	}

	if err.OvkError != nil {
		return false, errors.New(err.OvkError.ErrorMsg)
	}

	return false, nil
}
