import { redirect } from '@sveltejs/kit';

export async function load({ cookies }: any) {
	// i choose to directly check for the refresh token
	// because if the user is not authenticated, they should not be able to access this page
	// and no need to do things when checking for the access token

	const refreshToken = cookies.get('refreshToken');

	// if no refresh token, set user to null
	if (refreshToken) {
		redirect(302, '/user/profile');
	}
}
// the ssr set to be false because no content is rendered on this page
// for fast redirects and to avoid unnecessary server-side rendering
export let prerender = true;
export let ssr = false;
