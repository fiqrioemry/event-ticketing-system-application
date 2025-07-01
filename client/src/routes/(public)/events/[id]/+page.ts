// src/routes/events/[id]/+page.ts
import { getEventDetail } from '$lib/api/event';

export async function load({ params }) {
	const id = params.id;

	const data = await getEventDetail(id);

	if (!data) {
		return {
			status: 404,
			error: new Error('Event not found')
		};
	}

	return {
		event: data
	};
}
