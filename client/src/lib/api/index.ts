import axios, { type AxiosError } from 'axios';
import { type ApiResponse } from '$lib/types/api';

export const authInstance = axios.create({
	baseURL: import.meta.env.VITE_API_SERVICES,
	withCredentials: true,
	headers: {
		'X-API-Key': import.meta.env.VITE_API_KEY
	}
});

export const publicInstance = axios.create({
	baseURL: import.meta.env.VITE_API_SERVICES,
	headers: {
		'X-API-Key': import.meta.env.VITE_API_KEY
	}
});

// Response interceptor untuk handle konsisten API response
[authInstance, publicInstance].forEach((instance) => {
	instance.interceptors.response.use(
		(response) => {
			return response.data;
		},
		(error: AxiosError<ApiResponse>) => {
			const apiError = {
				success: false,
				data: null,
				status: error.response?.status,
				errors: error.response?.data?.errors || null,
				message: error.response?.data?.message || error.message || 'Network error'
			};
			return Promise.reject(apiError);
		}
	);
});
