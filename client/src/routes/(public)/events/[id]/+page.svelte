<script lang="ts">
	import { formatDate } from '$lib/utils';
	import { CalendarDays, MapPin, Minus, Plus } from '@lucide/svelte';

	export let data;
	const event = data.event;

	let selectedTickets: Record<string, number> = {};

	function increment(id: string) {
		selectedTickets[id] = (selectedTickets[id] || 0) + 1;
	}

	function decrement(id: string) {
		if (selectedTickets[id] > 0) {
			selectedTickets[id]--;
		}
	}
</script>

<section class="mx-auto max-w-5xl px-4 py-10">
	<!-- HEADER -->
	<div class="grid grid-cols-1 gap-8 md:grid-cols-2">
		<img src={event.image} alt={event.title} class="w-full rounded-xl object-cover shadow" />

		<div class="flex flex-col justify-between">
			<div>
				<h1 class="text-3xl font-bold">{event.title}</h1>
				<p class="text-muted-foreground mt-2">{event.description}</p>

				<div class="text-muted-foreground mt-4 space-y-2 text-sm">
					<div class="flex items-center gap-2">
						<CalendarDays class="h-4 w-4" />
						<span>{formatDate(event.date)} | {event.startTime}:00 – {event.endTime}:00</span>
					</div>
					<div class="flex items-center gap-2">
						<MapPin class="h-4 w-4" />
						<span>{event.location}</span>
					</div>
					<div>Status: <span class="font-semibold capitalize">{event.Status}</span></div>
				</div>
			</div>
		</div>
	</div>

	<!-- TICKET SELECTION -->
	<div class="mt-10 border-t pt-6">
		<h2 class="mb-4 text-xl font-semibold">Select Tickets</h2>

		{#each event.tickets as ticket}
			<div
				class="mb-4 flex flex-col gap-2 rounded-lg border px-4 py-3 sm:flex-row sm:items-center sm:justify-between dark:border-zinc-700"
			>
				<div>
					<p class="text-lg font-medium">{ticket.name}</p>
					<p class="text-muted-foreground text-sm">
						IDR {ticket.price.toLocaleString()} | Quota: {ticket.quota}
					</p>
				</div>
				<div class="flex items-center gap-2">
					<button
						on:click={() => decrement(ticket.id)}
						class="rounded border p-1"
						aria-label="Kurangi"
					>
						<Minus class="h-4 w-4" />
					</button>
					<span class="w-6 text-center">{selectedTickets[ticket.id] || 0}</span>
					<button
						on:click={() => increment(ticket.id)}
						class="rounded border p-1"
						aria-label="Tambah"
					>
						<Plus class="h-4 w-4" />
					</button>
				</div>
			</div>
		{/each}
	</div>
</section>
