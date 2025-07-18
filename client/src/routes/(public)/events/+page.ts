// src/routes/(public)/events/+page.ts
import { pagination } from '$lib/utils/state.js';
import type { EventQueryParams } from '$lib/types/api';
import * as eventService from '$lib/services/event.service';

export async function load({ url }) {
	const params: EventQueryParams = {
		search: url.searchParams.get('search') || undefined,
		startDate: url.searchParams.get('startDate') || undefined,
		endDate: url.searchParams.get('endDate') || undefined,
		location: url.searchParams.get('location') || undefined,
		page: parseInt(url.searchParams.get('page') || '1'),
		limit: parseInt(url.searchParams.get('limit') || '10')
	};
	const response = await eventService.getAllEvents(params);

	if (!response.success) {
		throw new Error(`Failed to load events: ${response.message}`);
	}

	return {
		events: response.data || [],
		pagination: response.meta?.pagination || pagination,
		params
	};
}
