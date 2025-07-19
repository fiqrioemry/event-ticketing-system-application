// src/routes/(public)/user/+layout.server.ts
import { redirect } from '@sveltejs/kit';

export async function load({ cookies, parent }) {
	const { user } = await parent();

	if (!user) {
		redirect(302, '/signin');
	}

	return {
		user
	};
}

export const ssr = true;
