package identity

import (
	"errors"
	common "github.com/Sraik25/KL-common"
)

// Value Object

type ContactInformation struct {
	emailAddress       EmailAddress
	postalAddress      PostalAddress
	primaryTelephone   Telephone
	secondaryTelephone Telephone
	common.Assertion
}

func NewContactInformation(emailAddress EmailAddress, postalAddress PostalAddress, primaryTelephone Telephone,
	secondaryTelephone Telephone) (ContactInformation, error) {
	nContactInformation := ContactInformation{}
	if err := nContactInformation.SetEmailAddress(emailAddress); err != nil {
		return ContactInformation{}, err
	}
	if err := nContactInformation.SetPostalAddress(postalAddress); err != nil {
		return ContactInformation{}, err
	}
	if err := nContactInformation.SetPrimaryTelephone(primaryTelephone); err != nil {
		return ContactInformation{}, err
	}
	if err := nContactInformation.SetSecondaryTelephone(secondaryTelephone); err != nil {
		return ContactInformation{}, err
	}

	return nContactInformation, nil
}

//#region Methods

func (v *ContactInformation) ChangeEmailAddress(emailAddress EmailAddress) (ContactInformation, error) {
	return NewContactInformation(
		emailAddress,
		v.PostalAddress(),
		v.PrimaryTelephone(),
		v.SecondaryTelephone(),
	)
}

func (v *ContactInformation) ChangePostalAddress(postalAddress PostalAddress) (ContactInformation, error) {
	return NewContactInformation(
		v.EmailAddress(),
		postalAddress,
		v.PrimaryTelephone(),
		v.SecondaryTelephone(),
	)
}

func (v *ContactInformation) ChangePrimaryTelephone(telephone Telephone) (ContactInformation, error) {
	return NewContactInformation(
		v.EmailAddress(),
		v.PostalAddress(),
		telephone,
		v.SecondaryTelephone(),
	)
}

func (v *ContactInformation) ChangeSecondaryTelephone(telephone Telephone) (ContactInformation, error) {
	return NewContactInformation(
		v.EmailAddress(),
		v.PostalAddress(),
		v.PrimaryTelephone(),
		telephone,
	)
}

//#endregion

//#region Getters and Setters

func (v *ContactInformation) EmailAddress() EmailAddress {
	return v.emailAddress
}

func (v *ContactInformation) PostalAddress() PostalAddress {
	return v.postalAddress
}

func (v *ContactInformation) PrimaryTelephone() Telephone {
	return v.primaryTelephone
}

func (v *ContactInformation) SecondaryTelephone() Telephone {
	return v.secondaryTelephone
}

func (v *ContactInformation) SetEmailAddress(emailAddress EmailAddress) error {
	if err := v.AssertArgumentNotNull(emailAddress, errors.New("")); err != nil {
		return err
	}
	v.emailAddress = emailAddress
	return nil
}

func (v *ContactInformation) SetPostalAddress(postalAddress PostalAddress) error {
	if err := v.AssertArgumentNotNull(postalAddress, errors.New("")); err != nil {
		return err
	}
	v.postalAddress = postalAddress
	return nil
}

func (v *ContactInformation) SetPrimaryTelephone(primaryTelephone Telephone) error {
	if err := v.AssertArgumentNotNull(primaryTelephone, errors.New("")); err != nil {
		return err
	}
	v.primaryTelephone = primaryTelephone
	return nil
}

func (v *ContactInformation) SetSecondaryTelephone(secondaryTelephone Telephone) error {
	if err := v.AssertArgumentNotNull(secondaryTelephone, errors.New("")); err != nil {
		return err
	}
	v.secondaryTelephone = secondaryTelephone
	return nil
}

//#endregion
