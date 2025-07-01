// /src/routes/(public)/(auth)/signout/+server.ts
import { logout } from '$lib/api/auth.js';
import { redirect, type RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async ({ locals }) => {
	const { userSession } = await locals.safeGetSession!();
	if (!userSession) throw redirect(303, '/');

	await logout();
	throw redirect(307, '/');
};
