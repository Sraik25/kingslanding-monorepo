package collaborator

type Owner struct {
	Collaborator
}

func NewOwner(emailAddress string, identity string, name string) Owner {
	return Owner{NewCollaborator(emailAddress, identity, name)}
}

func (v Owner) HasPrimeValue() int {
	return 59
}
