package KL_common

import "strings"

type Assertion struct{}

func (a Assertion) AssertArgumentNotEmpty(val string, err error) error {
	if val == "" || len(strings.TrimSpace(val)) == 0 {
		return err
	}
	return nil
}

func (a Assertion) AssertArgumentLength(val string, max int, err error) error {
	length := len(strings.TrimSpace(val))
	if length > max {
		return err
	}
	return nil
}

func (a Assertion) AssertArgumentLengthMinAndMax(val string, min, max int, err error) error {
	length := len(strings.TrimSpace(val))
	if length < min || length > max {
		return err
	}
	return nil
}

func (a Assertion) AssertArgumentNotNull(val any, err error) error {
	if val == nil {
		return err
	}
	return nil
}
