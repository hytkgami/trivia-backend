package domain

type Question struct {
	ID          string `db:"question_id"`
	CreatedBy   string `db:"created_by"`
	LobbyID     string `db:"lobby_id"`
	Title       string `db:"title"`
	OrderNumber int    `db:"order_number"`
	Score       int    `db:"score"`
}
