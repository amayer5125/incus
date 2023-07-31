package acl

import (
	"github.com/cyphar/incus/lxd/cluster/request"
	"github.com/cyphar/incus/lxd/state"
	"github.com/cyphar/incus/shared/api"
)

// NetworkACL represents a Network ACL.
type NetworkACL interface {
	// Initialise.
	init(state *state.State, id int64, projectName string, aclInfo *api.NetworkACL)

	// Info.
	ID() int64
	Project() string
	Info() *api.NetworkACL
	Etag() []any
	UsedBy() ([]string, error)

	// GetLog.
	GetLog(clientType request.ClientType) (string, error)

	// Internal validation.
	validateName(name string) error
	validateConfig(config *api.NetworkACLPut) error

	// Modifications.
	Update(config *api.NetworkACLPut, clientType request.ClientType) error
	Rename(newName string) error
	Delete() error
}