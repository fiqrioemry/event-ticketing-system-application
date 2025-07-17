import type {
	Order,
	OrderDetail,
	UserTicket,
	Pagination,
	ErrorResponse,
	OrderQueryParams,
	CreateOrderRequest
} from '$lib/types/api';
import { toast } from 'svelte-sonner';
import { loadStripe } from '@stripe/stripe-js';
import { writable, derived } from 'svelte/store';
import { createStoreActions } from '$lib/utils/store';
import * as orderService from '$lib/services/order.service';
import { goto } from '$app/navigation';

const stripePromise = loadStripe(import.meta.env.VITE_STRIPE_PUBLISHABLE_KEY);

interface OrderState {
	error: ErrorResponse | null;
	isLoading: boolean;
	isCreating: boolean;
	isRefunding: boolean;
	orders: Order[];
	orderDetail: OrderDetail[] | null;
	userTickets: UserTicket[];
	pagination?: Pagination;
}

const initialState: OrderState = {
	error: null,
	isLoading: false,
	isCreating: false,
	isRefunding: false,
	orders: [],
	orderDetail: null,
	userTickets: []
};

function createOrderStore() {
	const { subscribe, set, update } = writable<OrderState>(initialState);
	const actions = createStoreActions<Order>('orders', 'orderDetail');

	return {
		subscribe,

		async getMyOrders(params: OrderQueryParams) {
			actions.clearError(update);
			actions.setLoading(update, true);

			try {
				const response: any = await orderService.getMyOrders(params);
				actions.setItems(update, response.data, response.meta?.pagination);
			} catch (error: any) {
				const message = error.response?.data?.message || 'Failed to fetch orders';
				actions.setError(update, { message });
				toast.error(message);
			}
		},

		async getOrderDetail(id: string) {
			actions.clearError(update);
			actions.setLoading(update, true);

			try {
				const response: any = await orderService.getOrderDetail(id);
				update((state) => ({
					...state,
					orderDetail: response.data,
					error: null,
					isLoading: false
				}));
			} catch (error: any) {
				const message = error.response?.data?.message || 'Failed to fetch order detail';
				actions.setError(update, { message });
				toast.error(message);
			}
		},

		async getUserTicketsByOrderId(orderId: string) {
			actions.clearError(update);
			actions.setLoading(update, true);

			try {
				const response: any = await orderService.getUserTicketsByOrderId(orderId);
				update((state) => ({
					...state,
					userTickets: response.data,
					error: null,
					isLoading: false
				}));
			} catch (error: any) {
				const message = error.response?.data?.message || 'Failed to fetch user tickets';
				actions.setError(update, { message });
				toast.error(message);
			}
		},

		async createOrder(data: CreateOrderRequest): Promise<void> {
			actions.clearError(update);
			actions.setCreating(update, true);

			try {
				const response: any = await orderService.createOrder(data);
				const session = response.data;
				window.location.href = session.URL;
			} catch (error: any) {
				const message = error.response?.data?.message || 'Failed to initiate checkout';
				actions.setError(update, { message });
				toast.error(message);
			} finally {
				actions.setCreating(update, false);
			}
		},

		async refundOrder(orderId: string) {
			actions.clearError(update);
			update((state) => ({ ...state, isRefunding: true }));

			try {
				const response: any = await orderService.refundOrder(orderId);
				toast.success(response.message || 'Refund processed successfully');
				return response.data;
			} catch (error: any) {
				const message = error.response?.data?.message || 'Failed to process refund';
				actions.setError(update, { message });
				toast.error(message);
				throw error;
			} finally {
				update((state) => ({ ...state, isRefunding: false }));
			}
		},

		clearError: () => {
			actions.clearError(update);
		},

		reset: () => {
			actions.reset(update, initialState);
		}
	};
}

export const orderStore = createOrderStore();

// Derived states
export const orderList = derived(orderStore, ($s) => $s.orders);
export const orderDetail = derived(orderStore, ($s) => $s.orderDetail);
export const orderTickets = derived(orderStore, ($s) => $s.userTickets);
export const orderError = derived(orderStore, ($s) => $s.error);
export const isLoading = derived(orderStore, ($s) => $s.isLoading);
export const isCreating = derived(orderStore, ($s) => $s.isCreating);
export const isRefunding = derived(orderStore, ($s) => $s.isRefunding);
export const pagination = derived(orderStore, ($s) => $s.pagination);
