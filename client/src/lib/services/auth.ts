// src/lib/services/auth.ts
import type { LoginRequest, RegisterRequest } from '$lib/types/auth';
import { authInstance } from '.';

export class AuthService {
	static async login(data: LoginRequest) {
		console.log('Login data:', data);
		const response = await authInstance.post('/auth/login', data);
		return response.data;
	}

	static async register(data: RegisterRequest) {
		const response = await authInstance.post('/auth/register', data);
		return response.data;
	}

	static async logout() {
		const response = await authInstance.post('/auth/logout');
		return response.data;
	}

	static async getUserSession() {
		const response = await authInstance.get('/auth/me');
		return response.data;
	}

	static async refreshSession() {
		try {
			const response = await authInstance.post('/auth/refresh');
			return response.data;
		} catch (error) {
			try {
				await authInstance.post('/auth/logout');
			} catch (logoutError) {
				console.error('Logout error:', logoutError);
			}
			throw error;
		}
	}
}
