package domain

import "errors"

// TODO: Requests user
// Users exist in association with and under the control of a tenancy.
// Users of a system must be authenticated.
// Users possess personal information, including a name and contact information.
// User personal information may be changed by the users themselves or by a manager.
// User security credentials (passwords) may be changed.

var (
	ErrPasswordNotEmpty  = errors.New("the password may not be set to empty")
	ErrPersonNotEmpty    = errors.New("the person may not be set to null")
	ErrTenantIDNotEmpty  = errors.New("the tenantId may not be set to null")
	ErrUsernameNotChange = errors.New("the username may not be changed")
	ErrUsernameNotEmpty  = errors.New("the username may not be set to null")
)

type User struct {
	username string
	password string
	person   Person
	tenantID TenantID
}

//#region Constructor

func NewUser(tenantID TenantID, username, password string, person Person) (*User, error) {
	newUser := User{}
	err := newUser.SetPassword(password)
	if err != nil {
		return nil, err
	}
	err = newUser.SetPerson(person)
	if err != nil {
		return nil, err
	}
	err = newUser.SetTenantID(tenantID)
	if err != nil {
		return nil, err
	}
	err = newUser.SetUsername(username)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

//#endregion

//#region Getter and Setter

func (u *User) Username() string {
	return u.username
}

func (u *User) SetUsername(username string) error {
	if u.username != "" {
		return ErrUsernameNotChange
	}

	if username == "" {
		return ErrUsernameNotEmpty
	}
	u.username = username
	return nil
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetPassword(password string) error {
	if password == "" {
		return ErrPasswordNotEmpty
	}
	u.password = password
	return nil
}

func (u *User) Person() Person {
	return u.person
}

func (u *User) SetPerson(person Person) error {
	if person == (Person{}) {
		return ErrPersonNotEmpty
	}
	u.person = person
	return nil
}

func (u *User) TenantID() TenantID {
	return u.tenantID
}

func (u *User) SetTenantID(tenantID TenantID) error {
	if tenantID == (TenantID{}) {
		return ErrPersonNotEmpty
	}
	u.tenantID = tenantID
	return nil
}

//#endregion

// #region Methods User

func (u *User) ChangePassword(currentPassword, changedPassword string) error {
	// TODO: implement me
	return nil
}

func (u *User) ChangePersonalName(changedName string) error {
	// TODO: implement me
	return nil
}

func (u *User) ChangePersonalContactInformation(changedContactInformation any) error {
	// TODO: implement me
	return nil
}

//#endregion
