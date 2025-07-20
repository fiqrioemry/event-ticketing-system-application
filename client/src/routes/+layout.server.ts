// src/routes/+layout.server.ts
const API_KEY = import.meta.env.VITE_API_KEY || 'your_api_key';
const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:5004/api/v1';

export async function load({ cookies, fetch, url }) {
	const refreshToken = cookies.get('refreshToken');
	console.log('GETTING TOKEN');
	console.log('CURRENT URL :', url.pathname);
	console.log('CURRENT TIME :', new Date().toISOString());
	console.log('LOGGING REFRESH TOKEN :', refreshToken);
	console.log('===========LAYOUT============');
	if (!refreshToken) {
		return { user: null };
	}

	try {
		console.log('!! REFRESHING TOKEN !!');
		const response = await fetch(`${BASE_URL}/auth/refresh-token`, {
			method: 'POST',
			credentials: 'include',
			headers: {
				'X-API-KEY': API_KEY,
				'Content-Type': 'application/json'
			}
		});

		if (!response.ok) {
			return { user: null };
		}

		const userData = await response.json();
		return { user: userData.data };
	} catch (error) {
		return { user: null };
	}
}

export let prerender = false;
export let ssr = true;
