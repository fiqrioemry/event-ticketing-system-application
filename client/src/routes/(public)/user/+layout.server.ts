import { redirect } from '@sveltejs/kit';

export async function load({ cookies }) {
	const refreshToken = cookies.get('refreshToken');

	if (!refreshToken) {
		redirect(302, '/signin');
	}
}

export let prerender = true;
export let ssr = false;
