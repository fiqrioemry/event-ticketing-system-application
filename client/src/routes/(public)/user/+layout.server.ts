import { redirect } from '@sveltejs/kit';

export async function load({ parent, url }) {
	const { user } = await parent();

	if (!user) {
		const from = url.pathname + url.search;
		throw redirect(302, `/signin?redirectTo=${encodeURIComponent(from)}`);
	}

	return { user };
}
