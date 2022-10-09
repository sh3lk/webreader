package model

import (
	"webreader/ent"
)

// AuthInput represents a mutation input for creating users.
type AuthInput = ent.CreateUserInput

type Credential struct {
	Target *ent.User `json:"target"`
	Token  string    `json:"token"`
}
