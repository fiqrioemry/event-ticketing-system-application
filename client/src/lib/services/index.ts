// src/lib/services/index.ts
import axios from 'axios';
import { AuthService } from './auth';
import { browser } from '$app/environment';

export const authInstance = axios.create({
	baseURL: import.meta.env.VITE_API_SERVICES,
	withCredentials: true,
	headers: {
		'X-API-Key': import.meta.env.VITE_API_KEY
	}
});

let isRefreshing = false;
let failedQueue: Array<{
	resolve: (value?: any) => void;
	reject: (reason?: any) => void;
}> = [];

// Fungsi untuk memproses queue setelah refresh
const processQueue = (error: any, token: any) => {
	failedQueue.forEach(({ resolve, reject }) => {
		if (error) {
			reject(error);
		} else {
			resolve(token);
		}
	});

	failedQueue = [];
};

authInstance.interceptors.response.use(
	(response) => {
		return response;
	},
	async (error) => {
		const originalRequest = error.config;

		// Cek apakah ini 401 dan bukan request refresh/logout
		if (
			error.response?.status === 401 &&
			!originalRequest._retry &&
			!originalRequest.url.includes('/auth/refresh') &&
			!originalRequest.url.includes('/auth/logout')
		) {
			if (isRefreshing) {
				// Jika sedang refresh, masukkan ke queue
				return new Promise((resolve, reject) => {
					failedQueue.push({ resolve, reject });
				})
					.then(() => {
						return authInstance(originalRequest);
					})
					.catch((err) => {
						return Promise.reject(err);
					});
			}

			originalRequest._retry = true;
			isRefreshing = true;

			try {
				await AuthService.refreshSession();
				processQueue(null, true);
				return authInstance(originalRequest);
			} catch (refreshError) {
				processQueue(refreshError, null);

				if (browser) {
					window.location.href = '/signin';
				}

				return Promise.reject(refreshError);
			} finally {
				isRefreshing = false;
			}
		}

		return Promise.reject(error);
	}
);

authInstance.interceptors.request.use(
	(config) => {
		return config;
	},
	(error) => {
		return Promise.reject(error);
	}
);
