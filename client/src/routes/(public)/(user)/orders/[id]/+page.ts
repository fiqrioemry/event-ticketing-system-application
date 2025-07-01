// src/routes/events/[id]/+page.ts
import { getOrderDetail } from '$lib/api/order.js';

export async function load({ params }) {
	const id = params.id;

	const data = await getOrderDetail(id);

	if (!data) {
		return {
			status: 404,
			error: new Error('Order not found')
		};
	}

	return {
		order: data
	};
}
