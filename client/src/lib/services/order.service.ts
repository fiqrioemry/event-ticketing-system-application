import qs from 'qs';
import { api } from '$lib/api/client';
import type { ApiResponse } from '$lib/types/api';
import { authInstance } from '$lib/services/client';
import type { OrderQueryParams, CreateOrderRequest } from '$lib/types/api';

// GET /api/orders
export const getMyOrders = async (params: OrderQueryParams) => {
	const queryString = qs.stringify(params, { skipNulls: true });
	const res = await authInstance.get(`/orders?${queryString}`);
	return res.data;
};

// GET /api/orders/:id
export const getOrderDetail = async (id: string) => {
	const res = await authInstance.get(`/orders/${id}`);
	return res.data;
};

// POST /api/orders
export const createOrder = async (data: CreateOrderRequest) => {
	const res = await authInstance.post('/orders', data);
	return res.data;
};

// GET /api/orders/:id/user-tickets
export const getUserTicketsByOrderId = async (id: string) => {
	const res = await authInstance.get(`/orders/${id}/user-tickets`);
	return res.data;
};
// POST /api/orders/:id/refund
export const refundOrder = async (id: string) => {
	const res = await authInstance.post(`/orders/${id}/refund`);
	return res.data;
};
