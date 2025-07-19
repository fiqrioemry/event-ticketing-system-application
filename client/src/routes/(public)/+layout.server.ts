// src/routes/(public)/+layout.server.ts
const API_KEY = import.meta.env.VITE_API_KEY || 'your_api_key';
const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:5004/api/v1';

export async function load({ cookies, fetch, request }) {
	console.log('🚀 PUBLIC LAYOUT SERVER EXECUTED'); // Should appear

	const refreshToken = cookies.get('refreshToken');

	console.log('🌐 Request URL:', request.url);
	console.log('🍪 All server cookies:', cookies.getAll());
	console.log('🍪 RefreshToken from server:', refreshToken);

	if (!refreshToken) {
		console.log('❌ No refresh token found');
		return { user: null };
	}

	try {
		console.log('🔄 Attempting token refresh...');

		const response = await fetch(`${BASE_URL}/auth/refresh-token`, {
			method: 'POST',
			credentials: 'include',
			headers: {
				'X-API-KEY': API_KEY,
				'Content-Type': 'application/json',
				Cookie: `refreshToken=${refreshToken}`
			}
		});

		console.log('🔄 Refresh response status:', response.status);

		if (!response.ok) {
			console.log('❌ Refresh failed, clearing cookies');
			cookies.delete('refreshToken', { path: '/' });
			cookies.delete('accessToken', { path: '/' });
			return { user: null };
		}

		const userData = await response.json();
		console.log('✅ User authenticated via server');

		return {
			user: userData.data || userData.user || null
		};
	} catch (error) {
		console.error('💥 Token refresh failed:', error);
		cookies.delete('refreshToken', { path: '/' });
		cookies.delete('accessToken', { path: '/' });
		return { user: null };
	}
}

// ✅ FIXED - Remove prerender, keep SSR only
export const ssr = true;
