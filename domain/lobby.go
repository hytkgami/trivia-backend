package domain

type Lobby struct {
	ID       string `db:"lobby_id"`
	OwnerUID string `db:"owner_uid"`
	Name     string `db:"name"`
	IsPublic bool   `db:"is_public"`
}
