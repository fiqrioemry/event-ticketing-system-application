// src/routes/+layout.server.ts
const API_KEY = import.meta.env.VITE_API_KEY || 'your_api_key';
const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:5004/api/v1';

export async function load({ cookies, fetch, request }) {
	const refreshToken = cookies.get('refreshToken');
	console.log('Refresh Token:', refreshToken);

	// ✅ Debug: Log all cookies
	console.log('🍪 All server cookies:', cookies.getAll());
	console.log('🍪 RefreshToken from server:', refreshToken);
	console.log('🌐 Request URL:', request.url);
	console.log('🌐 Request headers:', request.headers.get('cookie'));
	console.log('🌐 Request headers:', BASE_URL);

	// No refresh token = not authenticated
	if (!refreshToken) {
		return {
			user: null
		};
	}

	// Try to refresh token and get user data
	try {
		const response = await fetch(`https://tiketku-api.ahmadfiqrioemry.com/auth/refresh-token`, {
			method: 'POST',
			credentials: 'include',
			headers: {
				'X-API-KEY': API_KEY,
				'Content-Type': 'application/json',
				Cookie: `refreshToken=${refreshToken}`
			}
		});

		if (!response.ok) {
			// ✅ Clear invalid token
			cookies.delete('refreshToken', {
				path: '/',
				domain: 'https://tiketku.ahmadfiqrioemry.com'
			});
			cookies.delete('accessToken', {
				path: '/',
				domain: 'https://tiketku.ahmadfiqrioemry.com'
			});

			return {
				user: null
			};
		}

		const userData = await response.json();
		return {
			user: userData.data || userData.user || null
		};
	} catch (error) {
		console.error('Token refresh failed:', error);

		cookies.delete('refreshToken', { path: '/' });
		cookies.delete('accessToken', { path: '/' });

		return {
			user: null
		};
	}
}

export const ssr = true;
