package model

type Lobby struct {
	ID       string `json:"id"`
	OwnerUID string `json:"-"`
	Name     string `json:"name"`
	Public   bool   `json:"public"`
}

func (Lobby) IsNode() {}

func (l Lobby) GetID() string {
	return l.ID
}
