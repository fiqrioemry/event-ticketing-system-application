<script lang="ts">
	import { CreditCard, Eye, Ticket } from '@lucide/svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { formatPrice, formatDate } from '$lib/utils/formatter';
	import type { Order } from '$lib/types/api';

	// Props
	export let order: Order;
	export let onViewDetail: (orderId: string) => void;
	export let onViewTickets: (orderId: string) => void;

	type OrderStatus = 'paid' | 'pending' | 'cancelled' | 'refunded';
	function getStatusBadge(status: string) {
		const statusMap: Record<OrderStatus, { class: string; label: string }> = {
			paid: { class: 'bg-green-100 text-green-800 border-green-200', label: 'Paid' },
			pending: { class: 'bg-yellow-100 text-yellow-800 border-yellow-200', label: 'Pending' },
			cancelled: { class: 'bg-red-100 text-red-800 border-red-200', label: 'Cancelled' },
			refunded: { class: 'bg-gray-100 text-gray-800 border-gray-200', label: 'Refunded' }
		};
		return (
			statusMap[status as OrderStatus] || {
				class: 'bg-gray-100 text-gray-800 border-gray-200',
				label: status
			}
		);
	}
</script>

<div
	class="overflow-hidden rounded-lg border border-gray-200 bg-white shadow-sm transition-shadow hover:shadow-md"
>
	<div class="p-6">
		<div class="flex flex-col gap-4 lg:flex-row lg:items-center lg:gap-6">
			<!-- Event Image -->
			<div class="flex-shrink-0">
				<div class="h-20 w-20 overflow-hidden rounded-lg lg:h-24 lg:w-24">
					<img src={order.eventImage} alt={order.eventName} class="h-full w-full object-cover" />
				</div>
			</div>

			<!-- Order Info -->
			<div class="flex-1 space-y-3">
				<div class="flex flex-col gap-2 lg:flex-row lg:items-start lg:justify-between">
					<div>
						<h3 class="text-lg font-semibold text-gray-900 lg:text-xl">
							{order.eventName}
						</h3>
						<p class="text-sm text-gray-500">Order ID: {order.id}</p>
					</div>

					<!-- Status Badge -->
					<span
						class="inline-flex items-center rounded-full border px-3 py-1 text-xs font-medium {getStatusBadge(
							order.status
						).class}"
					>
						{getStatusBadge(order.status).label}
					</span>
				</div>

				<!-- Order Details -->
				<div class="grid grid-cols-1 gap-3 text-sm text-gray-600 sm:grid-cols-2 lg:grid-cols-3">
					<div class="flex items-center gap-2">
						<svg
							class="h-4 w-4 text-gray-400"
							fill="none"
							stroke="currentColor"
							viewBox="0 0 24 24"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
							/>
						</svg>
						<span>Ordered: {formatDate(order.createdAt)}</span>
					</div>

					<div class="flex items-center gap-2">
						<svg
							class="h-4 w-4 text-gray-400"
							fill="none"
							stroke="currentColor"
							viewBox="0 0 24 24"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"
							/>
						</svg>
						<span class="font-medium text-gray-900">{formatPrice(order.totalPrice)}</span>
					</div>

					<div class="flex items-center gap-2">
						<svg
							class="h-4 w-4 text-gray-400"
							fill="none"
							stroke="currentColor"
							viewBox="0 0 24 24"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
							/>
						</svg>
						<span>{order.fullname}</span>
					</div>
				</div>
			</div>

			<!-- Action Buttons -->
			<div class="flex flex-col gap-2 lg:flex-shrink-0">
				<Button variant="outline" onclick={() => onViewDetail(order.id)}>
					<Eye /> View Details
				</Button>

				{#if order.status === 'paid'}
					<Button variant="primary" onclick={() => onViewTickets(order.id)}>
						<Ticket /> View Tickets
					</Button>
				{:else if order.status === 'pending' && order?.paymentUrl}
					<Button variant="event-secondary" href={order?.paymentUrl} target="_blank">
						<CreditCard />
						Complete Payment
					</Button>
				{/if}
			</div>
		</div>
	</div>
</div>
