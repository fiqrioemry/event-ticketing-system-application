<script>
	import { onMount } from 'svelte';
	import { fly, fade, scale } from 'svelte/transition';
	import { MapPin, Calendar, Clock } from '@lucide/svelte';
	import { formatDate } from '$lib/utils/formatter';

	export let event;

	let mounted = false;
	let imageLoaded = false;

	onMount(() => {
		mounted = true;
	});

	function handleImageLoad() {
		imageLoaded = true;
	}
</script>

<div class="mx-auto max-w-7xl px-4 pt-8 lg:px-8">
	<div
		class="relative h-[32rem] overflow-hidden rounded-2xl bg-gradient-to-r from-blue-600 to-indigo-600 shadow-2xl lg:h-[36rem]"
	>
		<!-- Background Image with animation -->
		{#if mounted}
			<img
				src={event.image}
				alt={event.title}
				class="absolute inset-0 h-full w-full object-cover transition-all duration-1000 {imageLoaded
					? 'scale-100 opacity-100'
					: 'scale-105 opacity-0'}"
				on:load={handleImageLoad}
				in:scale={{ duration: 1000, delay: 200 }}
			/>
		{/if}

		<!-- Improved gradient overlay for better text readability -->
		<div class="absolute inset-0 bg-gradient-to-t from-black/70 via-black/30 to-transparent"></div>
		<div class="absolute inset-0 bg-gradient-to-r from-blue-900/40 to-indigo-900/40"></div>

		<!-- Animated decorative elements -->
		{#if mounted}
			<div
				class="absolute top-8 right-8 h-32 w-32 animate-pulse rounded-full bg-white/10 blur-3xl"
				in:scale={{ duration: 800, delay: 400 }}
			></div>
			<div
				class="absolute right-16 bottom-16 h-24 w-24 animate-pulse rounded-full bg-blue-400/20 blur-2xl"
				style="animation-delay: 1s;"
				in:scale={{ duration: 800, delay: 600 }}
			></div>
			<!-- Additional floating elements -->
			<div
				class="absolute top-20 left-8 h-16 w-16 animate-pulse rounded-full bg-indigo-400/15 blur-2xl"
				style="animation-delay: 0.5s;"
				in:fade={{ duration: 800, delay: 800 }}
			></div>
		{/if}

		<!-- Content Container with enhanced animations -->
		<div class="absolute inset-0 flex flex-col justify-end">
			<div class="p-8 lg:p-12">
				{#if mounted}
					<!-- Title with improved typography and animation -->
					<h1
						class="mb-4 text-3xl leading-tight font-bold text-white drop-shadow-lg sm:text-4xl lg:text-5xl xl:text-6xl"
						in:fly={{ y: 50, duration: 800, delay: 600 }}
					>
						{event.title}
					</h1>

					<!-- Event details with improved layout and staggered animations -->
					<div
						class="flex flex-col gap-4 text-white/90 sm:flex-row sm:items-center sm:gap-6"
						in:fly={{ y: 30, duration: 700, delay: 800 }}
					>
						<!-- Location with enhanced animation -->
						<div
							class="group flex items-center gap-2 transition-transform duration-300 hover:scale-105"
							in:fade={{ duration: 500, delay: 1000 }}
						>
							<div
								class="flex h-8 w-8 items-center justify-center rounded-full bg-white/20 backdrop-blur-sm transition-all duration-300 group-hover:scale-110 group-hover:bg-white/30"
							>
								<MapPin class="h-4 w-4" />
							</div>
							<span
								class="text-sm font-medium transition-all duration-300 group-hover:text-white sm:text-base"
							>
								{event.location}
							</span>
						</div>

						<!-- Date with enhanced animation -->
						<div
							class="group flex items-center gap-2 transition-transform duration-300 hover:scale-105"
							in:fade={{ duration: 500, delay: 1200 }}
						>
							<div
								class="flex h-8 w-8 items-center justify-center rounded-full bg-white/20 backdrop-blur-sm transition-all duration-300 group-hover:scale-110 group-hover:bg-white/30"
							>
								<Calendar class="h-4 w-4" />
							</div>
							<span
								class="text-sm font-medium transition-all duration-300 group-hover:text-white sm:text-base"
							>
								{formatDate(event.date)}
							</span>
						</div>

						<!-- Time with enhanced animation -->
						<div
							class="group flex items-center gap-2 transition-transform duration-300 hover:scale-105"
							in:fade={{ duration: 500, delay: 1400 }}
						>
							<div
								class="flex h-8 w-8 items-center justify-center rounded-full bg-white/20 backdrop-blur-sm transition-all duration-300 group-hover:scale-110 group-hover:bg-white/30"
							>
								<Clock class="h-4 w-4" />
							</div>
							<span
								class="text-sm font-medium transition-all duration-300 group-hover:text-white sm:text-base"
							>
								{event.startTime}:00 - {event.endTime}:00
							</span>
						</div>
					</div>

					<!-- Additional status indicator if available -->
					{#if event.status === 'active'}
						<div
							class="mt-6 inline-flex items-center gap-2 rounded-full border border-green-400/30 bg-green-500/20 px-4 py-2 text-sm font-medium text-green-400 backdrop-blur-sm"
							in:scale={{ duration: 400, delay: 1600 }}
						>
							<div class="h-2 w-2 animate-pulse rounded-full bg-green-400"></div>
							Available for Booking
						</div>
					{:else if event.status === 'sold-out'}
						<div
							class="mt-6 inline-flex items-center gap-2 rounded-full border border-red-400/30 bg-red-500/20 px-4 py-2 text-sm font-medium text-red-400 backdrop-blur-sm"
							in:scale={{ duration: 400, delay: 1600 }}
						>
							<div class="h-2 w-2 rounded-full bg-red-400"></div>
							Sold Out
						</div>
					{/if}
				{/if}
			</div>
		</div>

		<!-- Shimmer effect overlay while image loads -->
		{#if !imageLoaded}
			<div
				class="animate-shimmer absolute inset-0 bg-gradient-to-r from-gray-400 via-gray-300 to-gray-400"
			></div>
		{/if}
	</div>
</div>

<style>
	/* Shimmer animation for loading state */
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

	/* Enhanced pulse for decorative elements */
	@keyframes gentle-pulse {
		0%,
		100% {
			opacity: 0.6;
			transform: scale(1);
		}
		50% {
			opacity: 1;
			transform: scale(1.1);
		}
	}
</style>
