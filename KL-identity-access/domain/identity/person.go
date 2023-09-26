package identity

// Entity

type Person struct {
	contactInformation ContactInformation
	tenantID           TenantID
	name               FullName
	user               *User
}

//#region Constructor

func NewPerson(contactInformation ContactInformation, tenantID TenantID, name FullName, user *User) *Person {

	nPerson := &Person{}

	nPerson.SetContactInformation(contactInformation)
	nPerson.SetTenantID(tenantID)
	nPerson.SetName(name)
	nPerson.SetUser(user)

	return nPerson
}

//#endregion

//#region Methods

func (p *Person) ChangeContactInformation(contactInformation ContactInformation) {
	p.SetContactInformation(contactInformation)
}

//#endregion

//#region Getters and Setters

func (p *Person) ContactInformation() ContactInformation {
	return p.contactInformation
}

func (p *Person) TenantID() TenantID {
	return p.tenantID
}

func (p *Person) Name() FullName {
	return p.name
}

func (p *Person) User() *User {
	return p.user
}

func (p *Person) SetContactInformation(contactInformation ContactInformation) {
	p.contactInformation = contactInformation
}

func (p *Person) SetTenantID(tenantID TenantID) {
	p.tenantID = tenantID
}

func (p *Person) SetName(name FullName) {
	p.name = name
}

func (p *Person) SetUser(user *User) {
	p.user = user
}

//#endregion
