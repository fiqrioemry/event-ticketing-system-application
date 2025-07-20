// routes/events/+page.ts
import * as eventService from '$lib/services/event.service';

export async function load({ params }) {
	const eventId = params.id;

	const response = await eventService.getEventById(eventId);

	if (!response.success) {
		throw new Error(`Failed to load events: ${response.message}`);
	}

	return {
		event: response.data || null
	};
}

export const ssr = true;
export const prerender = false;
