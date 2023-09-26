package identity

import (
	"github.com/google/uuid"
)

// Value Object:

type TenantID struct {
	id string
}

func (t *TenantID) Id() string {
	return t.id
}

func (t *TenantID) SetId(id string) error {
	parse, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	t.id = parse.String()
	return nil
}
