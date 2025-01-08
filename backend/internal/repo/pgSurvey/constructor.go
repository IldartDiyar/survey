package pgSurvey

import (
	"database/sql"
	"survey/pkg/models"
)

type Repository interface {
	CreateRespondent() (int, error)
	CreateAnswer(respondentID int, questionText, answerText string) (int, error)
	GetAnswers() ([]models.Answer, error)
}

type repository struct {
	db *sql.DB
}

func New(conn *sql.DB) Repository {
	return &repository{
		db: conn,
	}
}
