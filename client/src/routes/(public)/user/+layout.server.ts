// src/routes/(public)/user/+layout.server.ts
import { redirect } from '@sveltejs/kit';

export async function load({ cookies }) {
	const refreshToken = cookies.get('refreshToken');

	if (!refreshToken) {
		redirect(302, '/signin');
	}
}

export const prerender = false;
export const ssr = true;
