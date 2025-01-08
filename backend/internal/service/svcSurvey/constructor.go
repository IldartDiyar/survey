package svcSurvey

import (
	"survey/internal/repo/pgSurvey"
	"survey/pkg/models"
)

type service struct {
	repo pgSurvey.Repository
}

type Service interface {
	ProcessSurveyResponse(responseJSON map[string]string) error
	GetAnswers() ([]models.Answer, error)
}

func New(r pgSurvey.Repository) Service {
	return &service{
		repo: r,
	}
}
