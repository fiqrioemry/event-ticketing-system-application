<script lang="ts">
	import OrderItem from '$lib/components/orders/OrderItem.svelte';
	import type { Order } from '$lib/types/api';

	// Props
	export let orders: Order[] = [];
	export let hasActiveFilters: boolean = false;
	export let onViewDetail: (orderId: string) => void;
	export let onViewTickets: (orderId: string) => void;
	export let onResetFilters: () => void;
</script>

<div class="mt-8">
	{#if orders.length === 0}
		<!-- Empty State -->
		<div class="rounded-lg border-2 border-dashed border-gray-300 p-12 text-center">
			<svg
				class="mx-auto h-12 w-12 text-gray-400"
				fill="none"
				stroke="currentColor"
				viewBox="0 0 24 24"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"
				/>
			</svg>
			<h3 class="mt-4 text-lg font-medium text-gray-900">No orders found</h3>
			<p class="mt-2 text-sm text-gray-500">
				{hasActiveFilters
					? 'Try adjusting your filters or search terms.'
					: 'Start by booking your first event!'}
			</p>
			{#if hasActiveFilters}
				<button
					on:click={onResetFilters}
					class="mt-4 inline-flex items-center rounded-md bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700"
				>
					Clear filters
				</button>
			{:else}
				<a
					href="/events"
					class="mt-4 inline-flex items-center rounded-md bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700"
				>
					Browse Events
				</a>
			{/if}
		</div>
	{:else}
		<!-- Orders Grid -->
		<div class="space-y-4">
			{#each orders as order}
				<OrderItem {order} {onViewDetail} {onViewTickets} />
			{/each}
		</div>
	{/if}
</div>
