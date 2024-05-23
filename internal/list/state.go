package list

import "github.com/google/uuid"

type selectedUUIDState uuid.UUID

func (s *selectedUUIDState) Update(val selectedUUIDState) {
	*s = val
}

func (s selectedUUIDState) Get() selectedUUIDState {
	return s
}

func (s selectedUUIDState) GetAsUUID() uuid.UUID {
	return uuid.UUID(s)
}

var selectedUUID selectedUUIDState
