import type { OrderQueryParams } from '$lib/types/api';
import * as order from '$lib/services/order.service';

export async function load({ url }) {
	const params: OrderQueryParams = {
		search: url.searchParams.get('search') || undefined,
		status: url.searchParams.get('status') || undefined,
		page: parseInt(url.searchParams.get('page') || '1'),
		limit: parseInt(url.searchParams.get('limit') || '10')
	};

	try {
		const response = await order.getMyOrders(params);

		return {
			orders: response.data || [],
			pagination: response.meta?.pagination || {},
			params
		};
	} catch (error) {
		return {
			orders: [],
			pagination: {},
			params
		};
	}
}

export const prerender = false;
export const ssr = false;
