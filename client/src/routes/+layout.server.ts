// src/routes/(public)/+layout.server.ts
const API_KEY = import.meta.env.VITE_API_KEY || 'your_api_key';
const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:5004/api/v1';

export async function load({ cookies, fetch, request }) {
	const refreshToken = cookies.get('refreshToken');
	console.log('🚀 🚀 ROOT LAYOUT SERVER EXECUTED : /(PUBLIC)/LAYOUT.SERVER.TS');
	console.log('🌐 Request URL:', new Date().toString());
	console.log('🌐 Request URL:', request.url);
	console.log('🍪 RefreshToken from server:', refreshToken);

	if (!refreshToken) {
		console.log('❌ No refresh token found');
		return { user: null };
	}

	try {
		const response = await fetch(`${BASE_URL}/auth/refresh-token`, {
			method: 'POST',
			credentials: 'include',
			headers: {
				'X-API-KEY': API_KEY,
				'Content-Type': 'application/json'
			}
		});

		if (!response.ok) {
			cookies.delete('refreshToken', { path: '/' });
			cookies.delete('accessToken', { path: '/' });
			return { user: null };
		}

		const userData = await response.json();

		return {
			user: userData.data || userData.user || null
		};
	} catch (error) {
		cookies.delete('refreshToken', { path: '/' });
		cookies.delete('accessToken', { path: '/' });
		return { user: null };
	}
}

export const ssr = true;
