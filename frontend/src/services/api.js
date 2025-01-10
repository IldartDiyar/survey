import axios from 'axios';

const API_BASE_URL = 'http://165.227.143.119:8089/survey';

export const fetchSurveyResponses = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching survey responses:', error);
    return [];
  }
};
