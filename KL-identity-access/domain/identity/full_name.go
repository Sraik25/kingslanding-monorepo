package identity

import (
	"errors"
	"regexp"
)

// Value Object

var (
	ErrFirstNameCannotBeBlank = errors.New("first name is required")
)

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(firstName string, lastName string) (FullName, error) {
	nFullName := FullName{}
	err := nFullName.SetFirstName(firstName)
	if err != nil {
		return FullName{}, err
	}
	nFullName.SetLastName(lastName)
	return nFullName, nil
}

//#region Getter and Setter

func (f *FullName) FirstName() string {
	return f.firstName
}

func (f *FullName) LastName() string {
	return f.lastName
}

func (f *FullName) SetFirstName(firstName string) error {
	if firstName == "" {
		return ErrFirstNameCannotBeBlank
	}

	if len(firstName) < 1 || len(firstName) > 50 {
		return errors.New("first name must be 50 characters or less")
	}

	matchString, _ := regexp.MatchString("[A-Z][a-z]*", firstName)
	if !matchString {
		return errors.New("last name must be at least one character in length")
	}

	f.firstName = firstName
	return nil
}

func (f *FullName) SetLastName(lastName string) error {
	if lastName == "" {
		return ErrFirstNameCannotBeBlank
	}

	if len(lastName) < 1 || len(lastName) > 50 {
		return errors.New("first name must be 50 characters or less")
	}

	matchString, _ := regexp.MatchString("^[a-zA-Z'][ a-zA-Z'-]*[a-zA-Z']?", lastName)
	if !matchString {
		return errors.New("last name must be at least one character in length")
	}

	f.lastName = lastName
	return nil
}

//#endregion

//#region Methods

func (f *FullName) AsFormattedName() string {
	return f.FirstName() + " " + f.LastName()
}

func (f *FullName) withChangedFirstName(firstName string) (FullName, error) {
	return NewFullName(firstName, f.LastName())
}
func (f *FullName) withChangedLastName(lastName string) (FullName, error) {
	return NewFullName(f.FirstName(), lastName)
}

//#endregion
