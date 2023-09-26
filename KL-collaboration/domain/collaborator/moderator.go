package collaborator

// Value Object

type Moderator struct {
	Collaborator
}

func NewModerator(emailAddress string, identity string, name string) Moderator {
	return Moderator{NewCollaborator(emailAddress, identity, name)}
}

func (v Moderator) HasPrimeValue() int {
	return 19
}
