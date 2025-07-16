import axios from 'axios';

const apiInstance = axios.create({
  baseURL: 'http://localhost:3030/api/v1',
  timeout: 5000,
  headers: {
    'x-api-key': 'abcdefgh',
  },
});

apiInstance.interceptors.response.use(
  (response) => {
    if (response.data && response.data?.metadata) {
      return response.data;
    }

    return response.data;
  },
  (error) => error
);

export default apiInstance;
