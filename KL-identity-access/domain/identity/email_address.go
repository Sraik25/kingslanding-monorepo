package identity

import (
	"errors"
	"regexp"
)

// Value Object

var (
	ErrEmailAddressNotEmpty  = errors.New("the email address is required")
	ErrEmailAddressIsTooLong = errors.New("email address must be 100 characters or less")
	ErrEmailAddressIsInvalid = errors.New("email address and/or its format is invalid")
)

type EmailAddress struct {
	address string
}

func NewEmailAddress(address string) (EmailAddress, error) {
	nEmailAddress := EmailAddress{}
	err := nEmailAddress.SetAddress(address)
	if err != nil {
		return EmailAddress{}, err
	}
	return nEmailAddress, nil
}

func (e *EmailAddress) Address() string {
	return e.address
}

func (e *EmailAddress) SetAddress(address string) error {
	if address == "" {
		return ErrEmailAddressNotEmpty
	}

	if len(address) > 100 {
		return ErrEmailAddressIsTooLong
	}

	matchString, _ := regexp.MatchString(`\\w+([-+.']\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*`, address)
	if !matchString {
		return ErrEmailAddressIsInvalid
	}

	e.address = address
	return nil
}
