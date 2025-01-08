package models

type Answer struct {
	ID           int
	RespondentID int
	QuestionText string
	AnswerText   string
}
