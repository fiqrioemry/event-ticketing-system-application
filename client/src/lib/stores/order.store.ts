import { toast } from 'svelte-sonner';
import { loadStripe } from '@stripe/stripe-js';
import { writable, derived } from 'svelte/store';
import { createStoreActions } from '$lib/utils/store';
import * as orderService from '$lib/services/order.service';
import type { Order, ErrorResponse, CreateOrderRequest } from '$lib/types/api';

const stripePromise = loadStripe(import.meta.env.VITE_STRIPE_PUBLISHABLE_KEY);

interface OrderState {
	error: ErrorResponse | null;
	isCreating: boolean;
	isRefunding: boolean;
}

const initialState: OrderState = {
	error: null,
	isCreating: false,
	isRefunding: false
};

function createOrderStore() {
	const { subscribe, set, update } = writable<OrderState>(initialState);
	const actions = createStoreActions<Order>('orders', 'orderDetail');

	return {
		subscribe,

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
export const orderError = derived(orderStore, ($s) => $s.error);
export const isCreating = derived(orderStore, ($s) => $s.isCreating);
export const isRefunding = derived(orderStore, ($s) => $s.isRefunding);
