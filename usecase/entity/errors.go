package entity

import "errors"

type ValidationErr struct {
	Err string
}

func (ve *ValidationErr) Error() string {
	if ve == nil {
		return ""
	}

	return ve.Err
}

var (
	ErrDuplicateId        = errors.New("err. duplicate id")
	ErrValidateSearchReq  = errors.New("err. wrong search request")
	ErrForbiddenCharacter = errors.New("err. forbidden character")
)
