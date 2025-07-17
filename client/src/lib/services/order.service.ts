import qs from 'qs';
import type {
	OrderDetail,
	Order,
	UserTicket,
	RefundRequest,
	CheckoutSessionResponse,
	OrderQueryParams,
	CreateOrderRequest
} from '$lib/types/api';
import type { ApiResponse } from '$lib/types/api';
import { authInstance } from '$lib/services/client';

// GET /api/orders
export const getMyOrders = async (params: OrderQueryParams): Promise<ApiResponse<Order[]>> => {
	const queryString = qs.stringify(params, { skipNulls: true });
	const res = await authInstance.get(`/orders?${queryString}`);
	return res.data;
};

// GET /api/orders/:id
export const getOrderDetail = async (id: string): Promise<ApiResponse<OrderDetail>> => {
	const res = await authInstance.get(`/orders/${id}`);
	return res.data;
};

// POST /api/orders
export const createOrder = async (
	data: CreateOrderRequest
): Promise<ApiResponse<CheckoutSessionResponse>> => {
	const res = await authInstance.post('/orders', data);
	console.log('createOrder response:', res);
	return res.data;
};

// GET /api/orders/:id/user-tickets
export const getUserTicketsByOrderId = async (id: string): Promise<ApiResponse<UserTicket[]>> => {
	const res = await authInstance.get(`/orders/${id}/user-tickets`);
	console.log(res);
	return res.data;
};

// POST /api/orders/:id/refund
export const refundOrder = async (id: string): Promise<ApiResponse<RefundRequest>> => {
	const res = await authInstance.post(`/orders/${id}/refund`);
	return res.data;
};
