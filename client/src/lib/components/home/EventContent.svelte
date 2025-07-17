<script lang="ts">
	import { goto } from '$app/navigation';
	import { formatDate, formatPrice } from '$lib/utils/formatter';

	export let error;
	export let isLoading;
	export let featuredEvents;

	function navigateToEvents() {
		goto('/events');
	}

	function navigateToEvent(eventId: string) {
		goto(`/events/${eventId}`);
	}
</script>

<section class="bg-gray-50 py-20">
	<div class="container mx-auto px-4">
		<div class="mb-16 text-center">
			<h2 class="mb-4 text-4xl font-bold text-gray-800">Featured Events</h2>
			<p class="text-xl text-gray-600">Don't miss these amazing upcoming events</p>
		</div>

		{#if isLoading}
			<div class="mb-12 grid gap-8 md:grid-cols-3">
				{#each Array(3) as _}
					<div class="animate-pulse overflow-hidden rounded-2xl bg-white shadow-lg">
						<div class="h-48 bg-gray-300"></div>
						<div class="p-6">
							<div class="mb-2 h-4 rounded bg-gray-300"></div>
							<div class="mb-4 h-3 rounded bg-gray-300"></div>
							<div class="flex items-center justify-between">
								<div class="h-4 w-20 rounded bg-gray-300"></div>
								<div class="h-3 w-16 rounded bg-gray-300"></div>
							</div>
						</div>
					</div>
				{/each}
			</div>
		{:else if error}
			<div class="py-12 text-center">
				<p class="mb-4 text-red-600">{error}</p>
				<button
					on:click={navigateToEvents}
					class="rounded-lg bg-blue-600 px-6 py-2 text-white transition-colors hover:bg-blue-700"
				>
					View All Events Instead
				</button>
			</div>
		{:else if featuredEvents.length > 0}
			<div class="mb-12 grid gap-8 md:grid-cols-3">
				{#each featuredEvents as event}
					<div
						class="group cursor-pointer overflow-hidden rounded-2xl bg-white shadow-lg transition-all duration-300 hover:shadow-xl"
						on:click={() => navigateToEvent(event.id)}
						role="button"
						tabindex="0"
						on:keydown={(e) => e.key === 'Enter' && navigateToEvent(event.id)}
					>
						<div class="relative h-48 overflow-hidden">
							<img
								src={event.image}
								alt={event.title}
								class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
							/>
							<div class="absolute inset-0 bg-black/20"></div>
							<div class="absolute bottom-4 left-4 text-white">
								<div class="text-sm font-medium">{formatDate(event.date)}</div>
								<div class="text-2xl font-bold">Event</div>
							</div>
							{#if event.status === 'active'}
								<div
									class="absolute top-4 right-4 rounded bg-green-500 px-2 py-1 text-xs font-medium text-white"
								>
									Available
								</div>
							{/if}
						</div>
						<div class="p-6">
							<h3 class="mb-2 line-clamp-1 text-xl font-bold text-gray-800">{event.title}</h3>
							<p class="mb-4 line-clamp-2 text-gray-600">{event.description}</p>
							<div class="flex items-center justify-between">
								<span class="text-lg font-bold text-blue-600">{formatPrice(event.startPrice)}</span>
								<span class="text-sm text-gray-500">{event.location}</span>
							</div>
						</div>
					</div>
				{/each}
			</div>
		{:else}
			<div class="py-12 text-center">
				<p class="mb-4 text-gray-600">No featured events available at the moment.</p>
				<button
					on:click={navigateToEvents}
					class="rounded-lg bg-blue-600 px-6 py-2 text-white transition-colors hover:bg-blue-700"
				>
					Browse All Events
				</button>
			</div>
		{/if}

		<div class="text-center">
			<button
				on:click={navigateToEvents}
				class="rounded-lg bg-gradient-to-r from-blue-600 to-indigo-600 px-8 py-3 font-semibold text-white transition-all duration-300 hover:from-blue-700 hover:to-indigo-700 hover:shadow-lg"
			>
				View All Events
			</button>
		</div>
	</div>
</section>
