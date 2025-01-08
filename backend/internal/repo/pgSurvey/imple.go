package pgSurvey

import (
	"fmt"
	"survey/pkg/models"
)

func (r *repository) CreateRespondent() (int, error) {
	var id int
	query := "INSERT INTO respondents DEFAULT VALUES RETURNING id"
	err := r.db.QueryRow(query).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create respondent: %w", err)
	}
	return id, nil
}

func (r *repository) CreateAnswer(respondentID int, questionText, answerText string) (int, error) {
	var id int
	query := "INSERT INTO answers (respondent_id, question_text, answer_text) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, respondentID, questionText, answerText).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create answer: %w", err)
	}
	return id, nil
}

func (r *repository) GetAnswers() ([]models.Answer, error) {
	query := "SELECT id, respondent_id, question_text, answer_text FROM answers"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch answers: %w", err)
	}
	defer rows.Close()

	var answers []models.Answer
	for rows.Next() {
		var answer models.Answer
		if err := rows.Scan(&answer.ID, &answer.RespondentID, &answer.QuestionText, &answer.AnswerText); err != nil {
			return nil, fmt.Errorf("failed to scan answer: %w", err)
		}
		answers = append(answers, answer)
	}
	return answers, nil
}
