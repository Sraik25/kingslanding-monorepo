package identity

// Entity

type Person struct {
	contactInformation ContactInformation
	tenantID           TenantID
	name               FullName
	user               *User
}
