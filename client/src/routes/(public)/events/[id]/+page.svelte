<!-- /routes/(public)/events/[id]/+page.svelte -->
<script lang="ts">
	import { goto } from '$app/navigation';
	import { authUser } from '$lib/stores/auth.store.js';
	import PurchaseDialog from '$lib/components/event-detail/PurchaseDialog.svelte';
	import TicketSelection from '$lib/components/event-detail/TicketSelection.svelte';
	import EventDetailHero from '$lib/components/event-detail/EventDetailHero.svelte';
	import EventDetailLoading from '$lib/components/event-detail/EventDetailLoading.svelte';

	export let data: any;

	$: event = data.event;
	let showPurchaseDialog = false;
	let selectedTickets: { [key: string]: number } = {};

	// Initialize selectedTickets when event loads
	$: if (event?.tickets) {
		const initialTickets: { [key: string]: number } = {};
		event.tickets.forEach((ticket: any) => {
			initialTickets[ticket.id] = 0;
		});
		// Only update if selectedTickets is empty to avoid overriding user selections
		if (Object.keys(selectedTickets).length === 0) {
			selectedTickets = initialTickets;
		}
	}

	// Calculate totals - these will be reactive to selectedTickets changes
	$: totalAmount =
		event?.tickets?.reduce((sum: number, ticket: any) => {
			return sum + ticket.price * (selectedTickets[ticket.id] || 0);
		}, 0) || 0;

	$: totalQuantity = Object.values(selectedTickets).reduce(
		(sum: number, qty: number) => sum + qty,
		0
	);

	// Handle ticket change from child component
	function handleTicketChange(updatedTickets: { [key: string]: number }) {
		selectedTickets = updatedTickets;
	}

	// Handle dialog close
	function handleCloseDialog() {
		showPurchaseDialog = false;
	}

	// Handle purchase button click
	function handlePurchase() {
		// Check if user is logged in
		if (!$authUser) {
			goto('/signin');
			return;
		}

		if (totalQuantity === 0) {
			alert('Please select at least one ticket');
			return;
		}

		// Show purchase dialog
		showPurchaseDialog = true;
	}
</script>

<svelte:head>
	<title>{event?.title || 'Event Detail Information'}</title>
	<meta
		name="description"
		content="Discover and book amazing events in Indonesia. Find concerts, festivals, workshops, and more!"
	/>
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<meta property="og:title" content="Tiketku - Book Amazing Events" />
	<meta
		property="og:description"
		content="Book concerts, festivals, workshops & sports events in Indonesia"
	/>
	<meta property="og:image" content="/og-tiketku.jpg" />
</svelte:head>

<!-- Purchase Dialog -->
<PurchaseDialog
	{event}
	{totalAmount}
	{showPurchaseDialog}
	{selectedTickets}
	onCloseDialog={handleCloseDialog}
/>

<main class="min-h-screen bg-slate-100">
	{#if event}
		<!-- Hero Section -->
		<EventDetailHero {event} />

		<div class="container mx-auto max-w-6xl px-4 py-8">
			<div class="grid grid-cols-1 gap-8 lg:grid-cols-3">
				<!-- Event Information -->
				<div class="lg:col-span-2">
					<div class="mb-6 rounded-lg bg-white p-6 shadow-md">
						<h2 class="mb-4 text-2xl font-bold text-gray-800">About This Event</h2>
						<p class="leading-relaxed text-gray-600">{event.description}</p>
					</div>
				</div>

				<!-- Ticket Selection -->
				<TicketSelection
					{event}
					{totalAmount}
					{totalQuantity}
					{handlePurchase}
					{selectedTickets}
					onTicketChange={handleTicketChange}
				/>
			</div>
		</div>
	{:else}
		<EventDetailLoading />
	{/if}
</main>
