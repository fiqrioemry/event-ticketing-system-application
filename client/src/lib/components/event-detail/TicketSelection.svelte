<!-- /lib/components/event-detail/TicketSelection.svelte -->
<script lang="ts">
	import type { EventDetail } from '$lib/types/api';
	import { formatPrice } from '$lib/utils/formatter';
	import Button from '$lib/components/ui/button/button.svelte';

	export let event: EventDetail;
	export let totalAmount: number;
	export let totalQuantity: number;
	export let handlePurchase: () => void;
	export let selectedTickets: { [key: string]: number };
	export let onTicketChange: (tickets: { [key: string]: number }) => void;

	// Handle quantity change
	function updateQuantity(ticketId: string, change: number) {
		const ticket = event.tickets.find((t: any) => t.id === ticketId);
		if (!ticket) return;

		const currentQty = selectedTickets[ticketId] || 0;
		const newQty = currentQty + change;

		// Check limits
		if (newQty < 0) return;
		if (newQty > (ticket.limit ?? Infinity)) return;
		if (newQty > ticket.quota - ticket.sold) return;

		// Update tickets and notify parent
		const updatedTickets = {
			...selectedTickets,
			[ticketId]: newQty
		};

		onTicketChange(updatedTickets);
	}
</script>

<div class="lg:col-span-1">
	<div class="sticky top-8 rounded-lg bg-white p-6 shadow-md">
		<h3 class="mb-4 text-xl font-bold text-gray-800">Select Tickets</h3>

		{#each event.tickets as ticket}
			<div class="mb-4 rounded-lg border p-4 last:mb-0">
				<div class="mb-2 flex items-start justify-between">
					<div>
						<h4 class="font-semibold text-gray-800">{ticket.name}</h4>
						<p class="text-lg font-bold text-blue-600">{formatPrice(ticket.price)}</p>
					</div>
					<div class="text-right">
						<p class="text-sm text-gray-500">Available: {ticket.quota - ticket.sold}</p>
						<p class="text-sm text-gray-500">Max per person: {ticket.limit}</p>
					</div>
				</div>

				{#if ticket.isRefundable}
					<span
						class="mb-3 inline-block rounded-full bg-green-100 px-2 py-1 text-xs text-green-800"
					>
						Refundable
					</span>
				{/if}

				<div class="flex items-center justify-between">
					<span class="text-sm text-gray-600">Quantity:</span>
					<div class="flex items-center gap-3">
						<Button
							onclick={() => updateQuantity(ticket.id, -1)}
							disabled={!selectedTickets[ticket.id] || selectedTickets[ticket.id] === 0}
							variant="outline"
							class="h-8 w-8 rounded-full">-</Button
						>
						<span class="w-8 text-center font-semibold">
							{selectedTickets[ticket.id] || 0}
						</span>
						<Button
							onclick={() => updateQuantity(ticket.id, 1)}
							disabled={selectedTickets[ticket.id] >= (ticket.limit ?? Infinity) ||
								selectedTickets[ticket.id] >= ticket.quota - ticket.sold}
							variant="primary"
							class="h-8 w-8 rounded-full">+</Button
						>
					</div>
				</div>
			</div>
		{/each}

		<!-- Total and Purchase Button -->
		<div class="mt-4 border-t pt-4">
			<div class="mb-4 flex items-center justify-between">
				<span class="text-lg font-semibold text-gray-800">Total:</span>
				<span class="text-xl font-bold text-blue-600">{formatPrice(totalAmount)}</span>
			</div>
			<button
				on:click={handlePurchase}
				disabled={totalQuantity === 0}
				class="w-full rounded-lg bg-blue-600 px-6 py-3 font-semibold text-white transition-colors hover:bg-blue-700 disabled:cursor-not-allowed disabled:opacity-50"
			>
				Purchase Tickets ({totalQuantity})
			</button>
		</div>
	</div>
</div>
