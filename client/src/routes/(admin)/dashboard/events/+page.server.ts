import { getEvents } from '$lib/api/event';

export async function load() {
	const response = await getEvents({
		search: '',
		limit: 10
	});

	return {
		events: response.events,
		pagination: response.pagination
	};
}
