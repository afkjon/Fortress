import axios from 'axios';

const API_URI : string = "http://localhost:8080"

const api = axios.create({
  baseURL: API_URI,
  withCredentials: true, // Important for sending cookies
  headers: {
    "Content-Type": "application/json",
  }
});

export default api;
