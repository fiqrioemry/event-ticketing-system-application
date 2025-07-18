const API_KEY = import.meta.env.VITE_API_KEY || 'your_api_key';
const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:5004/api/v1';

export async function load({ cookies, fetch }: any) {
	const refreshToken = cookies.get('refreshToken');

	// if no refresh token, set user to null
	if (!refreshToken) {
		return {
			user: null
		};
	}
	// else, fetch user data using refresh token
	try {
		const response = await fetch(`${BASE_URL}/auth/refresh-token`, {
			method: 'POST',
			credentials: 'include',
			headers: {
				'X-API-KEY': API_KEY,
				'Content-Type': 'application/json'
			}
		});

		// if the response is not ok, delete the refresh token cookie
		if (!response.ok) {
			cookies.delete('refreshToken', { path: '/' });
		}

		// if the response is ok, parse the user data
		const userData = await response.json();
		return { user: userData.data };
	} catch (error) {
		throw new Error('Failed to refresh token: ');
	}
}
