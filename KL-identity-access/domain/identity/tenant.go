package identity

// TODO:
// Tenants allow for the registration of many users by invitation.
// Tenants may be active or be deactivated.
// Users of a system must be authenticated but can be authenticated only if the Tenant is active.

type Tenant struct {
	name   string
	active bool
}

//#region Constructor

//#endregion

//#region Getter and Setter

func (t *Tenant) Name() string {
	return t.name
}

func (t *Tenant) SetName(name string) {
	t.name = name
}

//#endregion

// #region Methods

func (t *Tenant) Active() error {
	// TODO: implement me
	return nil
}

func (t *Tenant) Deactivate() error {
	// TODO: implement me
	return nil
}

func (t *Tenant) IsActive() bool {
	return t.active == true
}

func (t *Tenant) RegisterUser() (*User, error) {
	// TODO: implement me
	return nil, nil
}

//#endregion
