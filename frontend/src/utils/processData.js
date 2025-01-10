// src/utils/processData.js

/**
 * Processes raw survey data into a structured format for visualization.
 * @param {Array} data - Raw survey data.
 * @returns {Object} - Processed data categorized by questions.
 */
export const processSurveyData = (data) => {
    const questions = {};
  
    data.forEach((response) => {
      const { QuestionText, AnswerText } = response;
  
      if (!questions[QuestionText]) {
        questions[QuestionText] = [];
      }
  
      questions[QuestionText].push(AnswerText);
    });
  
    return questions;
  };
  