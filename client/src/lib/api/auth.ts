// lib/api/auth.ts - Improved with proper typing
import { publicInstance, authInstance } from '.';
import type { ApiResponse } from '$lib/types/api';
import type { LoginRequest, RegisterRequest } from '$lib/types/auth';

export const AuthAPI = {
	login: async (data: LoginRequest): Promise<ApiResponse> => {
		return publicInstance.post('/auth/login', data);
	},

	register: async (data: RegisterRequest): Promise<ApiResponse> => {
		return publicInstance.post('/auth/signup', data);
	},

	logout: async (): Promise<ApiResponse> => {
		return authInstance.post('/auth/logout');
	},

	getMe: async (): Promise<ApiResponse> => {
		return authInstance.get('/auth/me');
	},

	refresh: async (): Promise<ApiResponse> => {
		return authInstance.post('/auth/refresh');
	}
};
