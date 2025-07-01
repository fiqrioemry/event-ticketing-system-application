import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals }) => {
	const { userSession } = await locals.safeGetSession!();

	if (userSession) throw redirect(303, '/');
};
