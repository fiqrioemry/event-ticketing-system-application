import * as auth from '$lib/services/auth.service';
import { redirect } from '@sveltejs/kit';

export async function load({ url }) {
	const token = url.searchParams.get('token') || 'xaxaxaxa';
	try {
		const response = await auth.validateResetToken(token);
		if (response.success) {
			return {
				valid: true,
				token
			};
		}
	} catch (error) {
		redirect(404, 'page not found');
	}
}
