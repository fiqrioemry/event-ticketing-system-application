// src/routes/(public)/user/+layout.server.ts
import * as order from '$lib/services/order.service';

export async function load({ params }) {
	const orderId = params.id;
	try {
		const response = await order.getOrderDetail(orderId);
		return { orderDetail: response.data || [] };
	} catch {
		throw new Error('Failed to fetch order detail');
	}
}
export const ssr = true;
export const prerender = false;
