import { getEvents } from '$lib/api/event';
import { eventFilters } from '$lib/stores/event.store';
import { eventResponse, isLoading } from '$lib/stores/event.store';

export async function load({ url }) {
	const search = url.searchParams.get('search') || '';
	const page = Number(url.searchParams.get('page') || '1');
	const status = url.searchParams.get('status') || 'all';
	const sort = url.searchParams.get('sort') || 'date_asc';
	const location = url.searchParams.get('location') || 'all';
	const limit = Number(url.searchParams.get('limit') || '10');

	eventFilters.set({
		search,
		page,
		location,
		limit,
		status,
		sort
	});

	const response = await getEvents({ search, page, status, sort, limit, location });

	isLoading.set(false);

	eventResponse.set({
		events: response.events || [],
		pagination: response.pagination || {}
	});
}
