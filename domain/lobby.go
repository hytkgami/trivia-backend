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

const (
	LobbyStatusWaitingID  = "01H2ZS0YPS2NHG3G6PBW6DZTX1"
	LobbyStatusActiveID   = "01H2ZS10YPFV6PZWAGX5HNV4KT"
	LobbyStatusFinishedID = "01H2ZS13AVXH0CVS0S59PC2DHX"
)

func (s LobbyStatus) ID() string {
	switch s {
	case LobbyStatusWaiting:
		return LobbyStatusWaitingID
	case LobbyStatusActive:
		return LobbyStatusActiveID
	case LobbyStatusFinished:
		return LobbyStatusFinishedID
	default:
		return ""
	}
}

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
