package collaborator

// Value Object

type Collaborator struct {
	emailAddress string
	identity     string
	name         string
}

// #region Constructor

func NewCollaborator(emailAddress string, identity string, name string) Collaborator {
	nCollaborator := Collaborator{}

	nCollaborator.setEmailAddress(emailAddress)
	nCollaborator.setIdentity(identity)
	nCollaborator.setName(name)

	return nCollaborator
}

// #endregion

// #region Getters and Setters

func (v *Collaborator) EmailAddress() string {
	return v.emailAddress
}

func (v *Collaborator) Identity() string {
	return v.identity
}

func (v *Collaborator) Name() string {
	return v.name
}

func (v *Collaborator) setEmailAddress(emailAddress string) {
	v.emailAddress = emailAddress
}

func (v *Collaborator) setIdentity(identity string) {

	v.identity = identity
}

func (v *Collaborator) setName(name string) {
	v.name = name
}

// #endregion
