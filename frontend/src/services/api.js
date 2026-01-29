import axios from 'axios';

const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const getSettings = async (idInstance, apiToken) => {
  const response = await api.get('/settings', {
    params: {
      idInstance,
      apiToken,
    },
  });
  return response.data;
};

export const getStateInstance = async (idInstance, apiToken) => {
  const response = await api.get('/state', {
    params: {
      idInstance,
      apiToken,
    },
  });
  return response.data;
};

export const sendMessage = async (data) => {
  const response = await api.post('/message', data);
  return response.data;
};

export const sendFileByUrl = async (data) => {
  const response = await api.post('/file', data);
  return response.data;
};
