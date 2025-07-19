// src/routes/(auth)/+layout.ts
import { redirect } from '@sveltejs/kit';

export async function load({ parent }: any) {
	const { user } = await parent();

	if (user) {
		throw redirect(302, `/user/profile`);
	}

	return { user };
}

// the ssr set to be false because no content is rendered on this page
// for fast redirects and to avoid unnecessary server-side rendering
