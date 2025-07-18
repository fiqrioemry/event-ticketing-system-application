<script>
	// @ts-nocheck

	import { goto } from '$app/navigation';
	import { ArrowLeft } from '@lucide/svelte';
	import Badge from '$lib/components/shared/Badge.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { formatPrice, formatDate } from '$lib/utils/formatter';

	export let data;

	$: orderDetail = data?.orderDetail || [];
	$: orderInfo = data?.orderInfo || {};

	$: totalAmount = orderDetail.reduce((sum, item) => sum + item.price * item.quantity, 0);
	$: totalItems = orderDetail.reduce((sum, item) => sum + item.quantity, 0);

	function goBackToOrders() {
		goto('/user/orders', { replaceState: true });
	}
</script>

<svelte:head>
	<title>Tiketku - Order Details</title>
	<meta name="description" content="View and manage your order details in Tiketku." />
</svelte:head>

<!-- Breadcrumb Navigation -->
<nav class="mb-6" aria-label="Breadcrumb">
	<div class="flex items-center space-x-2 text-sm text-gray-500">
		<button
			on:click={goBackToOrders}
			class="flex items-center transition-colors hover:text-gray-700"
		>
			<svg class="mr-1 h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
				<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
			</svg>
			Your Orders
		</button>
		<svg class="h-4 w-4" fill="currentColor" viewBox="0 0 20 20">
			<path
				fill-rule="evenodd"
				d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 111.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
				clip-rule="evenodd"
			/>
		</svg>
		<span class="font-medium text-gray-900">Order Details</span>
	</div>
</nav>

<div class="mt-8 space-y-8">
	<!-- Order Summary Card -->
	<div class="overflow-hidden rounded-xl border border-gray-200 bg-white shadow-sm">
		<div class="bg-gradient-to-r from-blue-600 to-indigo-600 px-6 py-4">
			<div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between"></div>
		</div>

		<div class="p-6">
			<div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
				<!-- Total Items -->
				<div class="flex items-center space-x-3">
					<div class="flex h-10 w-10 items-center justify-center rounded-full bg-green-100">
						<svg
							class="h-5 w-5 text-green-600"
							fill="none"
							stroke="currentColor"
							viewBox="0 0 24 24"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M15 5v2m0 4v2m0 4v2M5 5a2 2 0 00-2 2v3a2 2 0 110 4v3a2 2 0 002 2h14a2 2 0 002-2v-3a2 2 0 110-4V7a2 2 0 00-2-2H5z"
							/>
						</svg>
					</div>
					<div>
						<p class="text-sm font-medium text-gray-500">Total Tickets</p>
						<p class="text-lg font-semibold text-gray-900">
							{totalItems}
							{totalItems === 1 ? 'ticket' : 'tickets'}
						</p>
					</div>
				</div>

				<!-- Total Amount -->
				<div class="flex items-center space-x-3">
					<div class="flex h-10 w-10 items-center justify-center rounded-full bg-yellow-100">
						<svg
							class="h-5 w-5 text-yellow-600"
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
					</div>
					<div>
						<p class="text-sm font-medium text-gray-500">Total Amount</p>
						<p class="text-lg font-semibold text-gray-900">{formatPrice(totalAmount)}</p>
					</div>
				</div>
			</div>
		</div>
	</div>

	<!-- Ticket Details -->
	<div class="overflow-hidden rounded-xl border border-gray-200 bg-white shadow-sm">
		<div class="border-b border-gray-200 bg-gray-50 px-6 py-4">
			<h3 class="text-lg font-semibold text-gray-900">Ticket Details</h3>
			<p class="text-sm text-gray-600">Breakdown of tickets in this order</p>
		</div>

		<div class="divide-y divide-gray-200">
			{#each orderDetail as ticket, index}
				<div class="p-6 transition-colors hover:bg-gray-50">
					<div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
						<div class="flex items-center space-x-4">
							<!-- Ticket Icon -->
							<div
								class="flex h-12 w-12 items-center justify-center rounded-full bg-gradient-to-br from-blue-500 to-indigo-600"
							>
								<span class="text-lg font-bold text-white">#{index + 1}</span>
							</div>

							<!-- Ticket Info -->
							<div>
								<h4 class="text-lg font-semibold text-gray-900">{ticket.ticketName}</h4>
								<p class="text-sm text-gray-500">Purchased on {formatDate(ticket.createdAt)}</p>
							</div>
						</div>

						<!-- Ticket Pricing -->
						<div class="flex flex-col items-start sm:items-end">
							<div class="flex items-center space-x-2">
								<span class="text-sm font-medium text-gray-500">Quantity:</span>
								<span
									class="rounded-full bg-blue-100 px-2 py-1 text-sm font-semibold text-blue-800"
								>
									{ticket.quantity}
								</span>
							</div>
							<div class="mt-1 flex items-center space-x-2">
								<span class="text-sm text-gray-500">{formatPrice(ticket.price)} each</span>
								<span class="text-lg font-bold text-gray-900"
									>{formatPrice(ticket.price * ticket.quantity)}</span
								>
							</div>
						</div>
					</div>

					<!-- Ticket Features (if needed) -->
					<div class="mt-4 flex flex-wrap gap-2">
						{#if ticket?.ticketName.toLowerCase().includes('vip')}
							<span
								class="inline-flex items-center rounded-full bg-purple-100 px-2.5 py-0.5 text-xs font-medium text-purple-800"
							>
								VIP Access
							</span>
						{/if}
						{#if ticket?.ticketName.toLowerCase().includes('regular')}
							<span
								class="inline-flex items-center rounded-full bg-gray-100 px-2.5 py-0.5 text-xs font-medium text-gray-800"
							>
								Standard Access
							</span>
						{/if}
					</div>
				</div>
			{/each}
		</div>

		<!-- Order Total -->
		<div class="border-t border-gray-200 bg-gray-50 px-6 py-4">
			<div class="flex items-center justify-between">
				<span class="text-lg font-semibold text-gray-900">Total Order</span>
				<span class="text-2xl font-bold text-gray-900">{formatPrice(totalAmount)}</span>
			</div>
		</div>
	</div>

	<!-- Action Buttons -->
	<div class="flex flex-col gap-3 sm:flex-row sm:justify-between">
		<Button variant="event-outline" onclick={goBackToOrders}>
			<ArrowLeft class="h-4 w-4" />
			Back to Orders
		</Button>
	</div>
</div>
