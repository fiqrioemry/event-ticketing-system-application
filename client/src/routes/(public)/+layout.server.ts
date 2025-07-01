import { redirect } from '@sveltejs/kit';

export async function load({ locals }) {
	const { userSession } = await locals.safeGetSession!();

	if (userSession && userSession.role === 'admin') {
		// prevent admin accessing public layout after login
		// redirect to dashboard if user is admin
		throw redirect(303, '/dashboard');
	}
	return {
		userSession
	};
}
