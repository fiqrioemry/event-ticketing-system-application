// src/lib/api/client.ts (Simplified)
import { browser } from '$app/environment';

const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:5004/api/v1';
const API_KEY = import.meta.env.VITE_API_KEY;

export async function apiRequest(endpoint: string, options: RequestInit = {}) {
	const url = `${BASE_URL}${endpoint}`;

	const defaultOptions: RequestInit = {
		headers: {
			'Content-Type': 'application/json',
			'X-API-KEY': API_KEY,
			...options.headers
		},
		credentials: 'include'
	};

	const mergedOptions = { ...defaultOptions, ...options };

	try {
		const response = await fetch(url, mergedOptions);

		if (browser && response.status === 401 && !endpoint.includes('refresh-token')) {
			const refreshResponse = await fetch(`${BASE_URL}/auth/refresh-token`, {
				method: 'POST',
				credentials: 'include',
				headers: { 'X-API-KEY': API_KEY }
			});

			if (refreshResponse.ok) {
				return fetch(url, mergedOptions);
			} else {
				window.location.href = '/signin';
				throw new Error('Session expired');
			}
		}

		return response;
	} catch (error) {
		throw error;
	}
}

export const api = {
	get: (endpoint: string, options: RequestInit = {}) => {
		return apiRequest(endpoint, {
			method: 'GET',
			...options
		});
	},

	post: (endpoint: string, data: any) =>
		apiRequest(endpoint, {
			method: 'POST',
			body: JSON.stringify(data)
		}),

	put: (endpoint: string, data: any) =>
		apiRequest(endpoint, {
			method: 'PUT',
			body: JSON.stringify(data)
		}),

	delete: (endpoint: string) => apiRequest(endpoint, { method: 'DELETE' })
};
