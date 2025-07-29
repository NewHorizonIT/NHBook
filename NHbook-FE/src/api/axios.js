import useAuthStore from '@/store/usAuthStore';
import axios from 'axios';

const apiInstance = axios.create({
  baseURL: 'http://localhost:3030/api/v1',
  timeout: 5000,
  headers: {
    'x-api-key': 'abcdefgh',
  },
  withCredentials: true, // Quan trọng khi dùng HttpOnly Cookie
});

// 1. Response data unwrap
apiInstance.interceptors.response.use(
  (response) => {
    if (response.data?.metadata) {
      return response.data;
    }
    return response;
  },
  (error) => Promise.reject(error)
);

// 2. Gắn accessToken vào headers
apiInstance.interceptors.request.use((config) => {
  const token = useAuthStore.getState().accessToken;
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// 3. Tự động refresh token nếu token hết hạn
apiInstance.interceptors.response.use(
  (res) => res,
  async (error) => {
    const originalRequest = error.config;

    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      try {
        // Gọi API refresh token (cookie HttpOnly sẽ tự được gửi vì withCredentials: true)
        const res = await apiInstance.post('/auth/refresh-token');

        const newToken = res.data.access_token;
        useAuthStore.getState().setIsAuth(true);
        useAuthStore.getState().setAccessToken(newToken);

        // Gắn token mới vào request cũ và gửi lại
        originalRequest.headers.Authorization = `Bearer ${newToken}`;
        return apiInstance(originalRequest);
      } catch (e) {
        useAuthStore.getState().logout();
        window.location.href = '/auth/login';
        return Promise.reject(e);
      }
    }

    return Promise.reject(error);
  }
);

export default apiInstance;
