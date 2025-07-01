import { userSession } from '$lib/stores/auth.store';

export function load({ data }) {
	if (data.userSession) {
		userSession.set(data.userSession);
	} else {
		userSession.set(null);
	}
	return {};
}
