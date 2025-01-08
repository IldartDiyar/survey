package svcSurvey

import (
	"fmt"
	"survey/pkg/models"
)

func (s *service) ProcessSurveyResponse(surveyResponses map[string]string) error {
	respondentID, err := s.repo.CreateRespondent()
	if err != nil {
		return fmt.Errorf("failed to create respondent: %v", err)
	}

	for question, answer := range surveyResponses {
		_, err := s.repo.CreateAnswer(respondentID, question, answer)
		if err != nil {
			return fmt.Errorf("failed to add answer for question '%s': %w", question, err)
		}
	}

	return nil
}

func (s *service) GetAnswers() ([]models.Answer, error) {
	answers, err := s.repo.GetAnswers()

	if err != nil {
		return nil, fmt.Errorf("failed to get all answers: %v", err)
	}

	return answers, nil
}
