// src/routes/(public)/user/+layout.ts
import { redirect } from '@sveltejs/kit';

export async function load({ parent }: any) {
	const { user } = await parent();
	if (!user) {
		throw redirect(302, `/signin`);
	}

	return { user };
}
