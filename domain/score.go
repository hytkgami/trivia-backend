package domain

type Score struct {
	AnswerID string `db:"answer_id"`
	Mark     Mark   `db:"mark"`
	Value    int    `db:"value"`
}

type Mark string

const (
	MarkCorrect   Mark = "correct"
	MarkNeutral   Mark = "neutral"
	MarkIncorrect Mark = "incorrect"
)

func (m Mark) IsValid() bool {
	switch m {
	case MarkCorrect, MarkNeutral, MarkIncorrect:
		return true
	}
	return false
}

func (m Mark) String() string {
	return string(m)
}
