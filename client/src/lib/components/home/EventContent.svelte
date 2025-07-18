<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import type { EventResponse } from '$lib/types/api';
	import { fly, fade, scale } from 'svelte/transition';
	import * as eventService from '$lib/services/event.service';
	import { formatDate, formatPrice } from '$lib/utils/formatter';

	let isLoading: boolean = true;
	let error: string | null = null;
	let featuredEvents: EventResponse[] = [];

	onMount(async () => {
		try {
			const response = await eventService.getAllEvents({ limit: 3 });
			featuredEvents = response.data;
		} catch (error: any) {
			console.log(error.message || 'error fetching featured events');
		} finally {
			isLoading = false;
		}
	});
	let mounted = false;
	let sectionVisible = false;

	onMount(() => {
		mounted = true;

		// Intersection Observer untuk trigger animation saat scroll
		const observer = new IntersectionObserver(
			(entries) => {
				entries.forEach((entry) => {
					if (entry.isIntersecting) {
						sectionVisible = true;
					}
				});
			},
			{ threshold: 0.1 }
		);

		const section = document.getElementById('featured-events');
		if (section) {
			observer.observe(section);
		}

		return () => observer.disconnect();
	});

	function navigateToEvents() {
		goto('/events');
	}

	function navigateToEvent(eventId: string) {
		goto(`/events/${eventId}`);
	}
</script>

<section id="featured-events" class="relative overflow-hidden bg-gray-50 py-20">
	<!-- Background decorative elements -->
	<div class="absolute inset-0 opacity-5">
		<div
			class="absolute top-10 left-10 h-32 w-32 animate-pulse rounded-full bg-blue-500 blur-3xl"
		></div>
		<div
			class="absolute right-10 bottom-10 h-40 w-40 animate-pulse rounded-full bg-indigo-500 blur-3xl"
			style="animation-delay: 2s;"
		></div>
	</div>

	<div class="relative z-10 container mx-auto px-4">
		{#if mounted}
			<!-- Header with enhanced animation -->
			<div class="mb-16 text-center">
				{#if sectionVisible}
					<h2
						class="mb-4 text-4xl font-bold text-gray-800"
						in:fly={{ y: -30, duration: 800, delay: 0 }}
					>
						Featured Events
					</h2>
					<p class="text-xl text-gray-600" in:fade={{ duration: 800, delay: 200 }}>
						Don't miss these amazing upcoming events
					</p>
				{/if}
			</div>

			{#if isLoading}
				<!-- Enhanced loading skeleton with animation -->
				<div class="mb-12 grid gap-8 md:grid-cols-3">
					{#each Array(3) as _, index}
						{#if sectionVisible}
							<div
								class="animate-pulse overflow-hidden rounded-2xl bg-white shadow-lg"
								in:scale={{ duration: 400, delay: index * 150 }}
							>
								<div
									class="animate-shimmer h-48 bg-gradient-to-r from-gray-200 via-gray-300 to-gray-200"
								></div>
								<div class="p-6">
									<div
										class="animate-shimmer mb-2 h-4 rounded bg-gradient-to-r from-gray-200 via-gray-300 to-gray-200"
									></div>
									<div
										class="animate-shimmer mb-4 h-3 rounded bg-gradient-to-r from-gray-200 via-gray-300 to-gray-200"
									></div>
									<div class="flex items-center justify-between">
										<div
											class="animate-shimmer h-4 w-20 rounded bg-gradient-to-r from-gray-200 via-gray-300 to-gray-200"
										></div>
										<div
											class="animate-shimmer h-3 w-16 rounded bg-gradient-to-r from-gray-200 via-gray-300 to-gray-200"
										></div>
									</div>
								</div>
							</div>
						{/if}
					{/each}
				</div>
			{:else if error}
				<!-- Error state with animation -->
				{#if sectionVisible}
					<div class="py-12 text-center" in:fade={{ duration: 600, delay: 300 }}>
						<div class="mb-4 text-lg text-red-600" in:fly={{ y: 20, duration: 500, delay: 500 }}>
							{error}
						</div>
						<button
							on:click={navigateToEvents}
							class="group relative overflow-hidden rounded-lg bg-blue-600 px-6 py-2 text-white transition-all duration-300 hover:scale-105 hover:bg-blue-700 hover:shadow-lg"
							in:scale={{ duration: 400, delay: 700 }}
						>
							<span class="relative z-10">View All Events Instead</span>
							<div
								class="absolute inset-0 bg-blue-700 opacity-0 transition-opacity duration-300 group-hover:opacity-100"
							></div>
						</button>
					</div>
				{/if}
			{:else if featuredEvents.length > 0}
				<!-- Featured events grid with staggered animation -->
				<div class="mb-12 grid gap-8 md:grid-cols-3">
					{#each featuredEvents as event, index}
						{#if sectionVisible}
							<div
								class="group cursor-pointer overflow-hidden rounded-2xl bg-white shadow-lg transition-all duration-500 hover:-translate-y-3 hover:scale-105 hover:shadow-2xl"
								on:click={() => navigateToEvent(event.id)}
								role="button"
								tabindex="0"
								on:keydown={(e) => e.key === 'Enter' && navigateToEvent(event.id)}
								in:fly={{ y: 50, duration: 600, delay: index * 200 }}
							>
								<div class="relative h-48 overflow-hidden">
									<img
										src={event.image}
										alt={event.title}
										class="h-full w-full object-cover transition-all duration-500 group-hover:scale-110 group-hover:brightness-110"
									/>
									<div
										class="absolute inset-0 bg-gradient-to-t from-black/50 to-transparent transition-opacity duration-300 group-hover:from-black/40"
									></div>
									<div
										class="absolute bottom-4 left-4 text-white transition-transform duration-300 group-hover:translate-y-[-4px]"
										in:fade={{ duration: 500, delay: index * 200 + 300 }}
									>
										<div class="text-sm font-medium opacity-90">{formatDate(event.date)}</div>
										<div class="text-2xl font-bold">Event</div>
									</div>
									{#if event.status === 'active'}
										<div
											class="absolute top-4 right-4 rounded-full bg-green-500 px-3 py-1 text-xs font-medium text-white shadow-lg transition-all duration-300 group-hover:scale-110 group-hover:bg-green-400"
											in:scale={{ duration: 400, delay: index * 200 + 400 }}
										>
											Available
										</div>
									{/if}
								</div>
								<div
									class="p-6 transition-all duration-300 group-hover:bg-gray-50"
									in:fade={{ duration: 500, delay: index * 200 + 500 }}
								>
									<h3
										class="mb-2 line-clamp-1 text-xl font-bold text-gray-800 transition-colors duration-300 group-hover:text-blue-600"
									>
										{event.title}
									</h3>
									<p
										class="mb-4 line-clamp-2 text-gray-600 transition-colors duration-300 group-hover:text-gray-700"
									>
										{event.description}
									</p>
									<div class="flex items-center justify-between">
										<span
											class="text-lg font-bold text-blue-600 transition-all duration-300 group-hover:scale-105 group-hover:text-blue-700"
										>
											{formatPrice(event.startPrice)}
										</span>
										<span
											class="text-sm text-gray-500 transition-colors duration-300 group-hover:text-gray-600"
										>
											{event.location}
										</span>
									</div>
								</div>
								<!-- Hover indicator line -->
								<div
									class="h-1 w-0 bg-gradient-to-r from-blue-500 to-indigo-500 transition-all duration-300 group-hover:w-full"
								></div>
							</div>
						{/if}
					{/each}
				</div>
			{:else}
				<!-- Empty state with animation -->
				{#if sectionVisible}
					<div class="py-12 text-center" in:fade={{ duration: 600, delay: 300 }}>
						<div class="mb-6 text-6xl opacity-20" in:scale={{ duration: 600, delay: 500 }}>🎫</div>
						<p class="mb-4 text-lg text-gray-600" in:fly={{ y: 20, duration: 500, delay: 700 }}>
							No featured events available at the moment.
						</p>
						<button
							on:click={navigateToEvents}
							class="group relative overflow-hidden rounded-lg bg-blue-600 px-6 py-2 text-white transition-all duration-300 hover:scale-105 hover:bg-blue-700 hover:shadow-lg"
							in:scale={{ duration: 400, delay: 900 }}
						>
							<span class="relative z-10">Browse All Events</span>
							<div
								class="absolute inset-0 bg-blue-700 opacity-0 transition-opacity duration-300 group-hover:opacity-100"
							></div>
						</button>
					</div>
				{/if}
			{/if}

			<!-- CTA Button with enhanced animation -->
			{#if sectionVisible}
				<div
					class="text-center"
					in:fly={{ y: 30, duration: 600, delay: featuredEvents.length * 200 + 600 }}
				>
					<button
						on:click={navigateToEvents}
						class="group relative overflow-hidden rounded-lg bg-gradient-to-r from-blue-600 to-indigo-600 px-8 py-3 font-semibold text-white transition-all duration-300 hover:scale-105 hover:from-blue-700 hover:to-indigo-700 hover:shadow-xl"
					>
						<span
							class="relative z-10 transition-transform duration-300 group-hover:translate-y-[-1px]"
							>View All Events</span
						>
						<div
							class="absolute inset-0 bg-gradient-to-r from-blue-700 to-indigo-700 opacity-0 transition-opacity duration-300 group-hover:opacity-100"
						></div>
						<div
							class="absolute inset-0 bg-white opacity-0 transition-opacity duration-300 group-hover:opacity-10"
						></div>
					</button>
				</div>
			{/if}
		{/if}
	</div>
</section>

<style>
	/* Enhanced shimmer animation for loading skeletons */
	@keyframes shimmer {
		0% {
			background-position: -468px 0;
		}
		100% {
			background-position: 468px 0;
		}
	}

	.animate-shimmer {
		background-size: 1000px 100%;
		animation: shimmer 2s infinite linear;
	}
</style>
