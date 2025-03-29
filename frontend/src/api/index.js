import axios from 'axios';

// Create axios instance for API calls
const apiClient = axios.create({
  baseURL: '/api', // This will be proxied to the backend by vue.config.js
  timeout: 10000, // Request timeout
  headers: {
    'Content-Type': 'application/json'
  }
});

// Add request interceptor
apiClient.interceptors.request.use(
  config => {
    // You can add auth tokens here if needed
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// Add response interceptor
apiClient.interceptors.response.use(
  response => {
    return response.data;
  },
  error => {
    // Handle response errors (like 401, 404, 500, etc.)
    console.error('API request failed:', error);
    return Promise.reject(error);
  }
);

export default apiClient;