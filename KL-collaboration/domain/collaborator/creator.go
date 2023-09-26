package collaborator

type Creator struct {
	Collaborator
}

func NewCreator(emailAddress string, identity string, name string) Creator {
	return Creator{NewCollaborator(emailAddress, identity, name)}
}

func (v Creator) HasPrimeValue() int {
	return 43
}
