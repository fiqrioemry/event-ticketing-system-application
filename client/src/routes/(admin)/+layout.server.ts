// src/routes/(admin)/dashboard/+layout.server.ts
import { error } from '@sveltejs/kit';

export async function load({ locals }) {
	const { userSession } = (await locals.safeGetSession?.()) ?? { userSession: null };

	if (!userSession || userSession.role !== 'admin') {
		throw error(404, 'Not Found');
	}

	return { userSession };
}
