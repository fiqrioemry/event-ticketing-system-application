import { redirect } from '@sveltejs/kit';

export async function load({ locals }) {
	const { userSession } = await locals.safeGetSession!();

	if (!userSession) throw redirect(303, '/');

	return {
		userSession
	};
}
