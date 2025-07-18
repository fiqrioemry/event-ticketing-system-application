// src/lib/api/client.ts
import axios from 'axios';
import { authStore } from '$lib/stores/auth.store';
import { browser } from '$app/environment'; // Import ini

// Public instance - untuk endpoint tanpa autentikasi
export const publicInstance = axios.create({
	baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:5004/api/v1',
	withCredentials: true,
	headers: {
		'X-API-Key': import.meta.env.VITE_API_KEY
	}
});

// Auth instance - untuk endpoint yang butuh autentikasi
export const authInstance = axios.create({
	baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:5004/api/v1',
	withCredentials: true,
	headers: {
		'X-API-Key': import.meta.env.VITE_API_KEY
	}
});

// Interceptor for auto refresh token
authInstance.interceptors.response.use(
	(response) => response,
	async (error) => {
		const originalRequest = error.config;

		if (
			error.response?.status === 401 &&
			!originalRequest._retry &&
			!originalRequest.url.includes('/auth/refresh-token')
		) {
			originalRequest._retry = true;

			try {
				const refreshResponse = await publicInstance.post('/auth/refresh-token');

				if (refreshResponse.data.success) {
					authStore.setUser(refreshResponse.data.data);
					return authInstance(originalRequest);
				}
			} catch (refreshError) {
				authStore.reset();

				if (browser) {
					window.location.href = '/signin';
				}

				return Promise.reject(refreshError);
			}
		}

		return Promise.reject(error);
	}
);
