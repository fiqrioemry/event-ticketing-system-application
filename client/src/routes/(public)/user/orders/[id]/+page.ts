import * as order from '$lib/services/order.service';

export async function load({ params }) {
	const orderId = params.id;
	try {
		const response = await order.getOrderDetail(orderId);
		return {
			orderDetail: response.data || []
		};
	} catch (error) {
		throw new Error('Failed to fetch order detail');
	}
}
