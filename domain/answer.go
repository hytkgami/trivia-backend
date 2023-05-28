package domain

type Answer struct {
	ID         string `db:"answer_id"`
	QuestionID string `db:"question_id"`
	UID        string `db:"uid"`
	Content    string `db:"content"`
}
