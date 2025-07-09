import { type Handle } from '@sveltejs/kit';
import { AuthService } from '$lib/services/auth';

export const handle: Handle = async ({ event, resolve }) => {
	const accessToken = event.cookies.get('accessToken');

	if (accessToken) {
		try {
			// Validasi token dengan backend
			const result = await AuthService.getUserSession();
			if (result && result.user) {
				event.locals.user = result.user;
			} else {
				// Jika user tidak valid, hapus cookie
				event.cookies.delete('accessToken', { path: '/' });
				event.cookies.delete('refreshToken', { path: '/' });
			}
		} catch (error) {
			// Jika error (misal token expired), hapus cookies
			event.cookies.delete('accessToken', { path: '/' });
			event.cookies.delete('refreshToken', { path: '/' });
		}
	}

	// Pastikan user selalu ada (null jika tidak authenticated)
	event.locals.user = event.locals.user || null;

	return resolve(event);
};
