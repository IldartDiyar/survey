CREATE TABLE Surveys (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20),
    created_at TIMESTAMP
);

CREATE TABLE Questions (
    id SERIAL PRIMARY KEY,
    question TEXT
);

CREATE TABLE Respondents (
    id SERIAL PRIMARY KEY,
    time TIMESTAMP
);

CREATE TABLE Answers (
    survey_id INT NOT NULL,
    question_id INT NOT NULL,
    respondent_id INT NOT NULL,
    answer_text TEXT,
    PRIMARY KEY (survey_id, question_id, respondent_id),
    CONSTRAINT fk_survey FOREIGN KEY (survey_id) REFERENCES Surveys (id),
    CONSTRAINT fk_question FOREIGN KEY (question_id) REFERENCES Questions (id),
    CONSTRAINT fk_respondent FOREIGN KEY (respondent_id) REFERENCES Respondents (id)
);
