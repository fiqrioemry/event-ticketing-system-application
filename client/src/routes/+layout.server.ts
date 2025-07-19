const API_KEY = import.meta.env.VITE_API_KEY || 'your_api_key';
const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:5004/api/v1';

export async function load({ cookies, fetch, url }) {
	const refreshToken = cookies.get('refreshToken');

	console.log('🚀 ROOT LAYOUT SERVER EXECUTED');
	console.log('🌐 Current URL:', url.pathname);
	console.log('🍪 RefreshToken:', !!refreshToken);

	if (!refreshToken) {
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
			// Clear cookies tanpa redirect di root layout
			cookies.delete('refreshToken', { path: '/' });
			cookies.delete('accessToken', { path: '/' });
			return { user: null };
		}

		const userData = await response.json();
		return { user: userData.data || userData.user || null };
	} catch (error) {
		console.error('Auth error:', error);
		cookies.delete('refreshToken', { path: '/' });
		cookies.delete('accessToken', { path: '/' });
		return { user: null };
	}
}

export let ssr = true;
