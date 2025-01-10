// src/App.js
import React, { useEffect, useState } from 'react';
import { fetchSurveyResponses } from './services/api';
import { processSurveyData } from './utils/processData';
import BarChartComponent from './components/BarChartComponent';
import PieChartComponent from './components/PieChartComponent';

function App() {
  const [surveyData, setSurveyData] = useState({});
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getSurveyData = async () => {
      const data = await fetchSurveyResponses();
      const processedData = processSurveyData(data);
      setSurveyData(processedData);
      setLoading(false);
    };

    getSurveyData();
  }, []);

  if (loading) {
    return <div>Loading survey data...</div>;
  }

  const barChartQuestions = [
    'How estimate our service? (from 0-10)',
    'How satisfied are you with the support you received?',
    'How well does our product/service meet your needs?',
    'How would you rate the quality of our product/service?',
    'How satisfied are you with our product/service overall?',
  ];

  const pieChartQuestions = [
    'Did you feel the support team understood your issue?',
  ];

  const openEndedQuestions = [
    'What three words would you use to describe your experience with our product/service?',
    'Are there any features or aspects of the product/service you find especially valuable?',
  ];

  return (
    <div>
      <h1>Survey Data Visualization</h1>

      {barChartQuestions.map((question) => {
        const answers = surveyData[question] || [];
        // For rating questions, we can count frequency of each rating
        const frequency = {};

        answers.forEach((answer) => {
          const numeric = parseInt(answer, 10);
          if (!isNaN(numeric)) {
            frequency[numeric] = (frequency[numeric] || 0) + 1;
          } else {
            frequency[answer] = (frequency[answer] || 0) + 1;
          }
        });

        const data = Object.keys(frequency).map((key) => ({
          label: key,
          value: frequency[key],
        }));

        return (
          <BarChartComponent
            key={question}
            data={data}
            dataKey={{ label: 'label', value: 'value' }}
            title={question}
          />
        );
      })}

      {pieChartQuestions.map((question) => {
        const answers = surveyData[question] || [];
        const frequency = {};

        answers.forEach((answer) => {
          frequency[answer] = (frequency[answer] || 0) + 1;
        });

        const data = Object.keys(frequency).map((key) => ({
          label: key,
          value: frequency[key],
        }));

        return (
          <PieChartComponent
            key={question}
            data={data}
            dataKey={{ label: 'label', value: 'value' }}
            title={question}
          />
        );
      })}

      {openEndedQuestions.map((question) => {
        const answers = surveyData[question] || [];
        const wordFrequency = {};

        answers.forEach((answer) => {
          const words = answer.split(/[, ]+/);
          words.forEach((word) => {
            const cleanedWord = word.toLowerCase().trim();
            if (cleanedWord) {
              wordFrequency[cleanedWord] = (wordFrequency[cleanedWord] || 0) + 1;
            }
          });
        });

        const data = Object.keys(wordFrequency).map((word) => ({
          label: word,
          value: wordFrequency[word],
        }));

        return (
          <BarChartComponent
            key={question}
            data={data}
            dataKey={{ label: 'label', value: 'value' }}
            title={question}
          />
        );
      })}
    </div>
  );
}

export default App;
