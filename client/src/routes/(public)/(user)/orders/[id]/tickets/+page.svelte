<script lang="ts">
	import QRCode from 'qrcode';
	import { onMount } from 'svelte';

	const userTickets = [
		{
			id: '4ea32066-53ee-4e09-8563-b0016cec23bd',
			eventId: 'ddaf8eb0-a68e-4316-8dc7-834d183faaf6',
			ticketId: 'd39ff313-db85-4a3f-9f33-fa8fb0c68019',
			ticketName: 'Gold',
			eventName: 'Event 3',
			qrCode: 'TICKET-d39ff313-db85-4a3f-9f33-fa8fb0c68019-2',
			isUsed: false
		},
		{
			id: '6f0be514-c261-4747-ade0-fc45518fcd45',
			eventId: 'ddaf8eb0-a68e-4316-8dc7-834d183faaf6',
			ticketId: '802a2b63-d941-4557-8dc7-0d0580d0580f',
			ticketName: 'Silver',
			eventName: 'Event 3',
			qrCode: 'TICKET-802a2b63-d941-4557-8dc7-0d0580d0580f-3',
			isUsed: false
		},
		{
			id: 'e600b894-be2c-42f9-80d7-77715f989ed2',
			eventId: 'ddaf8eb0-a68e-4316-8dc7-834d183faaf6',
			ticketId: '802a2b63-d941-4557-8dc7-0d0580d0580f',
			ticketName: 'Silver',
			eventName: 'Event 3',
			qrCode: 'TICKET-802a2b63-d941-4557-8dc7-0d0580d0580f-2',
			isUsed: false
		}
	];

	let qrImages: Record<string, string> = {};

	onMount(async () => {
		for (const ticket of userTickets) {
			const dataUrl = await QRCode.toDataURL(ticket.qrCode);
			qrImages[ticket.id] = dataUrl;
		}
	});

	function handlePrint(ticketId: string) {}
</script>

<svelte:head>
	<title>My Tickets | MyBrand</title>
	<meta name="description" content="See your ticket from your purchase detail" />
</svelte:head>

<section class="mx-auto max-w-3xl p-6">
	<h1 class="mb-6 text-2xl font-semibold">Your Tickets</h1>

	<div class="space-y-6">
		{#each userTickets as ticket}
			<div class="flex items-center justify-between rounded-xl border p-4 shadow-sm">
				<div>
					<h2 class="text-lg font-medium">{ticket.eventName}</h2>
					<p class="text-gray-600">Ticket: {ticket.ticketName}</p>
					<p class="mt-1 text-sm text-gray-500">QR: {ticket.qrCode}</p>
				</div>
				<div class="flex items-center gap-3">
					{#if qrImages[ticket.id]}
						<img src={qrImages[ticket.id]} alt="QR Code" class="h-20 w-20 object-contain" />
					{/if}
					<button
						on:click={() => handlePrint(ticket.id)}
						class="rounded-md border border-gray-300 px-4 py-2 text-sm hover:bg-gray-100"
					>
						Print
					</button>
				</div>
			</div>
		{/each}
	</div>
</section>
