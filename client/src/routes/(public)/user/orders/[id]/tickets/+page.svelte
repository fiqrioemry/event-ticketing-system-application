<script lang="ts">
	import { getQRCodeURL } from '$lib/utils/formatter';
	import TicketDialog from '$lib/components/user-tickets/TicketDialog.svelte';

	export let data;
	$: tickets = data?.userTickets || [];

	// State untuk QR code yang sedang dipilih
	let selectedQR: any = null;
	let showQRModal: boolean = false;

	function showQRCode(qrCode: string) {
		selectedQR = qrCode;
		showQRModal = true;
	}

	function downloadQR(qrCode: string | null, ticketName: string, eventName: string) {
		const canvas = document.createElement('canvas');
		const ctx = canvas.getContext('2d');
		const img = new Image();

		img.crossOrigin = 'anonymous';
		img.onload = function () {
			canvas.width = img.width;
			canvas.height = img.height;
			if (ctx) {
				ctx.drawImage(img, 0, 0);

				const link = document.createElement('a');
				link.download = `ticket-${ticketName}-${eventName}.png`;
				link.href = canvas.toDataURL();
				link.click();
			}
		};

		img.src = getQRCodeURL(qrCode, 300);
	}
</script>

<svelte:head>
	<title>Tiketku - Your Tickets</title>
	<meta name="description" content="View and manage your digital tickets in Tiketku." />
</svelte:head>

<TicketDialog open={showQRModal} {selectedQR} {downloadQR} />

<!-- Event Info Header -->
{#if tickets[0]}
	<div
		class="mb-8 rounded-lg border border-gray-200 bg-gradient-to-r from-blue-50 to-indigo-50 p-6"
	>
		<h2 class="text-2xl font-bold text-gray-900">{tickets[0].eventName}</h2>
		<p class="mt-2 text-gray-600">Total tickets: {tickets.length}</p>
	</div>
{/if}

<!-- Tickets Grid -->
<div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
	{#each tickets as ticket, index}
		<div
			class="group relative overflow-hidden rounded-xl border border-gray-200 bg-white shadow-sm transition-all duration-300 hover:scale-[1.02] hover:shadow-lg"
		>
			<!-- Ticket Header -->
			<div class="bg-gradient-to-r from-blue-600 to-indigo-600 px-6 py-4">
				<div class="flex items-center justify-between">
					<h3 class="text-lg font-semibold text-white">{ticket.ticketName}</h3>
					<!-- Ticket Status -->
					<span
						class="inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium {ticket.isUsed
							? 'bg-red-100 text-red-800'
							: 'bg-green-100 text-green-800'}"
					>
						{ticket.isUsed ? 'Used' : 'Valid'}
					</span>
				</div>

				<!-- Ticket Number -->
				<p class="mt-1 text-sm text-blue-100">Ticket #{index + 1}</p>
			</div>

			<!-- Ticket Body -->
			<div class="p-6">
				<!-- QR Code Section -->
				<div class="mb-6 text-center">
					<div
						class="mx-auto mb-4 h-32 w-32 overflow-hidden rounded-lg border-2 border-gray-200 bg-white p-2"
					>
						<img
							src={getQRCodeURL(ticket?.qrCode, 120)}
							alt="QR Code for {ticket.ticketName}"
							class="h-full w-full object-contain"
							loading="lazy"
						/>
					</div>
					<p class="text-xs text-gray-500">QR Code: {ticket?.qrCode}</p>
				</div>

				<!-- Ticket Info -->
				<div class="space-y-3 border-t border-gray-100 pt-4">
					<div class="flex items-center text-sm text-gray-600">
						<svg
							class="mr-2 h-4 w-4 text-gray-400"
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
						<span class="font-medium">{ticket?.eventName}</span>
					</div>

					<div class="flex items-center text-sm text-gray-600">
						<svg
							class="mr-2 h-4 w-4 text-gray-400"
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
						<span>Type: {ticket.ticketName}</span>
					</div>

					<div class="flex items-center text-sm text-gray-600">
						<svg
							class="mr-2 h-4 w-4 text-gray-400"
							fill="none"
							stroke="currentColor"
							viewBox="0 0 24 24"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
							/>
						</svg>
						<span>Status: {ticket.isUsed ? 'Used' : 'Ready to use'}</span>
					</div>
				</div>

				<!-- Action Buttons -->
				<div class="mt-6 space-y-2">
					<button
						on:click={() => showQRCode(ticket.qrCode)}
						class="flex w-full items-center justify-center gap-2 rounded-lg bg-blue-600 px-4 py-2.5 text-sm font-medium text-white transition-colors hover:bg-blue-700"
					>
						<svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM10 7v3m0 0v3m0-3h3m-3 0H7"
							/>
						</svg>
						View QR Code
					</button>

					<button
						on:click={() => downloadQR(ticket.qrCode, ticket.ticketName, ticket.eventName)}
						class="flex w-full items-center justify-center gap-2 rounded-lg border border-gray-300 bg-white px-4 py-2.5 text-sm font-medium text-gray-700 transition-colors hover:bg-gray-50"
					>
						<svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
							/>
						</svg>
						Download QR
					</button>
				</div>
			</div>

			<!-- Decorative Elements -->
			<div class="absolute -top-6 -right-6 h-12 w-12 rounded-full bg-white opacity-10"></div>
			<div class="absolute -bottom-6 -left-6 h-12 w-12 rounded-full bg-white opacity-10"></div>
		</div>
	{/each}
</div>

<!-- Summary Info -->
<div class="mt-8 rounded-lg border border-gray-200 bg-gray-50 p-6">
	<h3 class="text-lg font-semibold text-gray-900">Important Information</h3>
	<div class="mt-4 space-y-2 text-sm text-gray-600">
		<div class="flex items-start gap-2">
			<svg
				class="mt-0.5 h-4 w-4 flex-shrink-0 text-blue-600"
				fill="none"
				stroke="currentColor"
				viewBox="0 0 24 24"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
				/>
			</svg>
			<span>Present your QR code at the venue entrance for scanning</span>
		</div>
		<div class="flex items-start gap-2">
			<svg
				class="mt-0.5 h-4 w-4 flex-shrink-0 text-blue-600"
				fill="none"
				stroke="currentColor"
				viewBox="0 0 24 24"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
				/>
			</svg>
			<span>Each ticket can only be used once</span>
		</div>
		<div class="flex items-start gap-2">
			<svg
				class="mt-0.5 h-4 w-4 flex-shrink-0 text-blue-600"
				fill="none"
				stroke="currentColor"
				viewBox="0 0 24 24"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
				/>
			</svg>
			<span>Keep your tickets safe and don't share screenshots with others</span>
		</div>
	</div>
</div>
