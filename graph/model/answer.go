package model

type Answer struct {
	ID         string `json:"answer_id"`
	QuestionID string `json:"-"`
	UID        string `json:"-"`
	Content    string `json:"content"`
}

func (Answer) IsNode() {}

func (a Answer) GetID() string {
	return a.ID
}
