// src/routes/(public)/user/orders/[id]/tickets/+page.ts
import * as order from '$lib/services/order.service';

export async function load({ params }: { params: { id: string } }) {
	const orderId = params.id;
	try {
		const response = await order.getUserTicketsByOrderId(orderId);

		return {
			userTickets: response.data || []
		};
	} catch (error) {
		throw new Error('Failed to fetch user tickets');
	}
}

export const ssr = true;
