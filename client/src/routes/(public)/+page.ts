import { getEvents } from '$lib/api/event';

export async function load() {
	const res = await getEvents({ limit: 5 });
	return { events: res.events, pagination: res.pagination };
}
