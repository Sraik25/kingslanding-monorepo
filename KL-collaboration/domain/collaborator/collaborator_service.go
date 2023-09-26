package collaborator

import "github.com/Sraik25/KL-identity-access/domain/identity"

type CollaboratorService interface {
	AuthorFrom(tenant identity.Tenant, identity string) Author
	CreatorFrom(tenant identity.Tenant, identity string) Creator
	ModeratorFrom(tenant identity.Tenant, identity string) Moderator
	OwnerFrom(tenant identity.Tenant, identity string) Owner
	ParticipantFrom(tenant identity.Tenant, identity string) Participant
}
