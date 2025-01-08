package service

import (
	"survey/internal/repo"
	"survey/internal/service/svcSurvey"
)

type Container struct {
	SurveyService svcSurvey.Service
}

func New(repo *repo.Container) *Container {
	svcS := svcSurvey.New(repo.SurveyRepo)

	return &Container{
		SurveyService: svcS,
	}
}
