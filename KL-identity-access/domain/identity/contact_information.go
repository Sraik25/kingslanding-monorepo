package identity

// Value Object

type ContactInformation struct {
	emailAddress       EmailAddress
	postalAddress      PostalAddress
	primaryTelephone   Telephone
	secondaryTelephone Telephone
}
