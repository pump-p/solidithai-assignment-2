import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080'; // Backend URL

const axiosInstance = axios.create({
  baseURL: API_BASE_URL
});

// Add token to Authorization header
axiosInstance.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export const login = (credentials) =>
  axiosInstance.post('/auth/login', credentials);
export const signup = (user) => axiosInstance.post('/auth/signup', user);
export const getUsers = () => axiosInstance.get('/users/');
export const getUserById = (id) => axiosInstance.get(`/users/${id}`);
export const createUser = (user) => axiosInstance.post('/users', user);
export const updateUser = (id, updates) =>
  axiosInstance.put(`/users/${id}`, updates);
export const deleteUser = (id) => axiosInstance.delete(`/users/${id}`);
