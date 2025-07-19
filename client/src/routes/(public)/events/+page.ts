// src/routes/(public)/events/+page.ts
import { pagination } from '$lib/utils/state.js';
import type { EventQueryParams } from '$lib/types/api';
import * as eventService from '$lib/services/event.service';

export const prerender = false;

export async function load({ url }) {
	console.log('🍪 GETTING PAGE URL NOW:', url);
	const params: EventQueryParams = {
		search: url.searchParams.get('search') || undefined,
		startDate: url.searchParams.get('startDate') || undefined,
		endDate: url.searchParams.get('endDate') || undefined,
		location: url.searchParams.get('location') || undefined,
		page: parseInt(url.searchParams.get('page') || '1'),
		limit: parseInt(url.searchParams.get('limit') || '5')
	};
	try {
		const response = await eventService.getAllEvents(params);

		return {
			events: response.data || [],
			pagination: response.meta?.pagination || pagination,
			params
		};
	} catch (error) {
		console.error('Error loading events:', error);
		return {
			events: [],
			pagination,
			params
		};
	}
}
