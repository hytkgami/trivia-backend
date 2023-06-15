package domain

import (
	"encoding/json"
)

type Lobby struct {
	ID       string `db:"lobby_id"`
	OwnerUID string `db:"owner_uid"`
	Name     string `db:"name"`
	IsPublic bool   `db:"is_public"`
}

type LobbyStatus string

const (
	LobbyStatusWaiting  LobbyStatus = "waiting"
	LobbyStatusActive   LobbyStatus = "active"
	LobbyStatusFinished LobbyStatus = "finished"
)

func (s LobbyStatus) Valid() bool {
	switch s {
	case LobbyStatusWaiting, LobbyStatusActive, LobbyStatusFinished:
		return true
	}
	return false
}

func (s LobbyStatus) MarshalBinary() ([]byte, error) {
	return []byte(s), nil
}

func (s LobbyStatus) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}
