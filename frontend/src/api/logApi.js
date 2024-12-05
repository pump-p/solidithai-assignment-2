import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080';

export const queryLogs = (query) => {
  const token = localStorage.getItem('token');
  return axios.get(`${API_BASE_URL}/logs?q=${query}`, {
    headers: {
      Authorization: `Bearer ${token}`
    }
  });
};
