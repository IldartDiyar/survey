CREATE TABLE respondents (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE answers (
    id SERIAL PRIMARY KEY,
    respondent_id INT REFERENCES respondents(id) ON DELETE CASCADE,
    question_text TEXT NOT NULL,
    answer_text TEXT NOT NULL,
    CONSTRAINT fk_respondent FOREIGN KEY (respondent_id) REFERENCES respondents(id) ON DELETE CASCADE
);

CREATE INDEX idx_respondent_id ON answers (respondent_id);