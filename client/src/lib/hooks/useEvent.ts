import { getEvents } from '$lib/api/event';
import { eventResponse, isLoading } from '$lib/stores/event.store';
import type { EventQueryParams } from '$lib/types';

export async function useGetEvents(params: EventQueryParams) {
	try {
		isLoading.set(true);
		const response = await getEvents(params);
		eventResponse.set(response);
	} catch (error) {
		console.error('Error fetching events:', error);
	} finally {
		isLoading.set(false);
	}
}
