package domain

// TODO: Requests user
// Users exist in association with and under the control of a tenancy.
// Users of a system must be authenticated.
// Users possess personal information, including a name and contact information.
// User personal information may be changed by the users themselves or by a manager.
// User security credentials (passwords) may be changed.

type User struct {
	username string
	password string
	person   Person
}

//#region Getter and Setter

func (u *User) Username() string {
	return u.username
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
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
