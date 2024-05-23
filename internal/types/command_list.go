package types

import "github.com/google/uuid"

type CommandList struct {
	UUID uuid.UUID
	Cmd  string
	Desc string
}
