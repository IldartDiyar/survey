package repo

import (
	"survey/internal/infra"
	"survey/internal/repo/pgSurvey"
)

type Container struct {
	SurveyRepo pgSurvey.Repository
}

func New(infra *infra.Container) *Container {
	surveyRepo := pgSurvey.New(infra.Conn)

	return &Container{
		SurveyRepo: surveyRepo,
	}
}
