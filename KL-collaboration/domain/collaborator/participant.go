package collaborator

type Participant struct {
	Collaborator
}

func NewParticipant(emailAddress string, identity string, name string) Participant {
	return Participant{NewCollaborator(emailAddress, identity, name)}
}

func (v Participant) HasPrimeValue() int {
	return 59
}
