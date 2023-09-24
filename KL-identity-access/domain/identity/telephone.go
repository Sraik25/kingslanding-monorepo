package identity

import (
	"errors"
	"regexp"

	common "github.com/Sraik25/KL-common"
)

// Value Object

type Telephone struct {
	number string
	common.Assertion
}

//#region Constructor

func NewTelephone(number string) (Telephone, error) {
	nTelephone := Telephone{}
	if err := nTelephone.SetNumber(number); err != nil {
		return Telephone{}, err
	}
	return nTelephone, nil
}

//#endregion

//#region Getter and Setter

func (v *Telephone) Number() string {
	return v.number
}

func (v *Telephone) SetNumber(number string) error {
	if err := v.AssertArgumentNotEmpty(number, errors.New("telephone number is required")); err != nil {
		return err
	}
	if err := v.AssertArgumentLength(number, 9,
		errors.New("telephone number may not be more than 9 characters")); err != nil {
		return err
	}
	if matchString, err := regexp.MatchString(`((\\(\\d{3}\\))|(\\d{3}-))\\d{3}-\\d{4}`, number); !matchString {
		return err
	}
	v.number = number
	return nil
}

//#endregion
