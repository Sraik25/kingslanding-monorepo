package identity

import (
	"errors"
	common "github.com/Sraik25/KL-common"
)

// Value Object

type PostalAddress struct {
	city          string
	countryCode   string
	postalCode    string
	stateProvince string
	streetAddress string
	common.Assertion
}

//#region Constructor

func NewPostalAddress(city string, countryCode string, postalCode string, stateProvince string, streetAddress string) (PostalAddress, error) {
	nPostalAddress := PostalAddress{}

	if err := nPostalAddress.SetCity(city); err != nil {
		return PostalAddress{}, err
	}
	if err := nPostalAddress.SetCountryCode(countryCode); err != nil {
		return PostalAddress{}, err
	}
	if err := nPostalAddress.SetPostalCode(postalCode); err != nil {
		return PostalAddress{}, err
	}
	if err := nPostalAddress.SetStateProvince(stateProvince); err != nil {
		return PostalAddress{}, err
	}
	if err := nPostalAddress.SetStreetAddress(streetAddress); err != nil {
		return PostalAddress{}, err
	}

	return nPostalAddress, nil
}

//#endregion

//#region Getter and Setter

func (v *PostalAddress) City() string {
	return v.city
}

func (v *PostalAddress) CountryCode() string {
	return v.countryCode
}

func (v *PostalAddress) PostalCode() string {
	return v.postalCode
}

func (v *PostalAddress) StateProvince() string {
	return v.stateProvince
}

func (v *PostalAddress) StreetAddress() string {
	return v.streetAddress
}

func (v *PostalAddress) SetCity(city string) error {
	if err := v.AssertArgumentNotEmpty(city, errors.New("the city is required")); err != nil {
		return err
	}
	if err := v.AssertArgumentLengthMinAndMax(city, 1, 100,
		errors.New("the city must be 100 characters or less")); err != nil {
		return err
	}
	v.city = city
	return nil
}

func (v *PostalAddress) SetCountryCode(countryCode string) error {
	if err := v.AssertArgumentNotEmpty(countryCode, errors.New("the country is required")); err != nil {
		return err
	}
	if err := v.AssertArgumentLengthMinAndMax(countryCode, 2, 2,
		errors.New("the country code must be two characters")); err != nil {
		return err
	}
	v.countryCode = countryCode
	return nil
}

func (v *PostalAddress) SetPostalCode(postalCode string) error {
	if err := v.AssertArgumentNotEmpty(postalCode, errors.New("the postal code is required")); err != nil {
		return err
	}
	if err := v.AssertArgumentLengthMinAndMax(postalCode, 5, 12,
		errors.New("the postal code must be 12 characters or less")); err != nil {
		return err
	}
	v.postalCode = postalCode
	return nil
}

func (v *PostalAddress) SetStateProvince(stateProvince string) error {
	if err := v.AssertArgumentNotEmpty(stateProvince, errors.New("the state/province is required")); err != nil {
		return err
	}
	if err := v.AssertArgumentLengthMinAndMax(stateProvince, 2, 100,
		errors.New("the state/province must be 100 characters or less")); err != nil {
		return err
	}
	v.stateProvince = stateProvince
	return nil
}

func (v *PostalAddress) SetStreetAddress(streetAddress string) error {
	if err := v.AssertArgumentNotEmpty(streetAddress, errors.New("the city is required")); err != nil {
		return err
	}
	if err := v.AssertArgumentLengthMinAndMax(streetAddress, 1, 100,
		errors.New("the street address must be 100 characters or less")); err != nil {
		return err
	}
	v.streetAddress = streetAddress
	return nil
}

//#endregion

//#region Methods

//#endregion
