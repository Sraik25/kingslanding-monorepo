package collaborator

type Author struct {
	Collaborator
}

func NewAuthor(emailAddress string, identity string, name string) Author {
	return Author{NewCollaborator(emailAddress, identity, name)}
}

func (v Author) HasPrimeValue() int {
	return 59
}
