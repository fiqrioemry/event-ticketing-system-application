<script>
	import { onMount } from 'svelte';
	import { fly, fade, scale } from 'svelte/transition';
	import { goto } from '$app/navigation';

	let mounted = false;

	onMount(() => {
		mounted = true;
	});
</script>

<section class="relative flex min-h-screen items-center justify-center overflow-hidden">
	<!-- Background with overlay -->
	<div class="absolute inset-0 z-0">
		<img src="banner1.webp" alt="Event Background" class="h-full w-full object-cover" />
		<div
			class="absolute inset-0 bg-gradient-to-r from-blue-900/80 via-blue-800/70 to-indigo-800/80"
		></div>
	</div>

	<!-- Enhanced Floating Elements with staggered animations -->
	{#if mounted}
		<div
			class="absolute top-20 left-10 h-20 w-20 animate-pulse rounded-full bg-white/10 transition-colors duration-300 hover:bg-white/20"
			in:scale={{ duration: 500, delay: 50 }}
		></div>
		<div
			class="absolute top-40 right-20 h-16 w-16 animate-pulse rounded-full bg-indigo-500/20 transition-colors duration-300 hover:bg-indigo-500/30"
			style="animation-delay: 2s;"
			in:scale={{ duration: 500, delay: 200 }}
		></div>
		<div
			class="absolute bottom-40 left-20 h-12 w-12 animate-pulse rounded-full bg-blue-500/20 transition-colors duration-300 hover:bg-blue-500/30"
			style="animation-delay: 4s;"
			in:scale={{ duration: 500, delay: 400 }}
		></div>

		<!-- Additional floating elements for more dynamic feel -->
		<div
			class="absolute top-60 left-1/4 h-8 w-8 animate-pulse rounded-full bg-purple-500/15"
			style="animation-delay: 1s;"
			in:fade={{ duration: 700, delay: 600 }}
		></div>
		<div
			class="absolute right-1/3 bottom-60 h-14 w-14 animate-pulse rounded-full bg-cyan-500/15"
			style="animation-delay: 3s;"
			in:fade={{ duration: 700, delay: 800 }}
		></div>
	{/if}

	<!-- Hero Content -->
	<div class="relative z-10 container mx-auto px-4 text-center text-white">
		{#if mounted}
			<div class="mx-auto max-w-4xl">
				<!-- Main Title with enhanced animation -->
				<h1
					class="mb-6 text-5xl font-bold md:text-7xl"
					in:fly={{ y: -50, duration: 900, delay: 100 }}
				>
					<span in:fly={{ y: -30, duration: 700, delay: 300 }}> Discover </span>
					<span
						class="bg-gradient-to-r from-indigo-400 to-blue-400 bg-clip-text text-transparent"
						in:fly={{ y: -30, duration: 700, delay: 500 }}
					>
						Amazing
					</span>
					<span in:fly={{ y: -30, duration: 700, delay: 700 }}> Events </span>
				</h1>

				<!-- Subtitle with fade in -->
				<p class="mb-8 text-xl text-gray-200 md:text-2xl" in:fade={{ duration: 900, delay: 900 }}>
					Your gateway to unforgettable experiences. Find, book, and enjoy the best events in your
					city.
				</p>

				<!-- CTA Buttons with staggered entrance -->
				<div
					class="mb-16 flex flex-col justify-center gap-4 sm:flex-row"
					in:fly={{ y: 30, duration: 700, delay: 1100 }}
				>
					<button
						on:click={() => goto('/events')}
						class="group relative overflow-hidden rounded-xl bg-gradient-to-r from-indigo-500 to-blue-600 px-8 py-4 text-lg font-semibold text-white transition-all duration-300 hover:scale-105 hover:from-indigo-600 hover:to-blue-700 hover:shadow-2xl"
						in:scale={{ duration: 400, delay: 1300 }}
					>
						<span class="relative z-10">Explore Events</span>
						<div
							class="absolute inset-0 bg-gradient-to-r from-indigo-600 to-blue-700 opacity-0 transition-opacity duration-300 group-hover:opacity-100"
						></div>
					</button>
					<a
						href="#features"
						class="group rounded-xl border-2 border-white/30 px-8 py-4 text-lg font-semibold text-white transition-all duration-300 hover:scale-105 hover:border-white/60 hover:bg-white/10"
						in:scale={{ duration: 400, delay: 1500 }}
					>
						<span class="transition-all duration-300 group-hover:text-indigo-200">Learn More</span>
					</a>
				</div>

				<!-- Stats with enhanced staggered animation -->
				<div class="grid grid-cols-3 gap-8" in:fly={{ y: 40, duration: 700, delay: 1700 }}>
					<div class="group cursor-pointer text-center" in:scale={{ duration: 300, delay: 1800 }}>
						<div
							class="text-3xl font-bold text-indigo-400 transition-all duration-300 group-hover:scale-110 group-hover:text-indigo-300 md:text-4xl"
						>
							1000+
						</div>
						<div class="text-gray-300 transition-colors duration-300 group-hover:text-gray-200">
							Events
						</div>
					</div>
					<div class="group cursor-pointer text-center" in:scale={{ duration: 300, delay: 2000 }}>
						<div
							class="text-3xl font-bold text-blue-400 transition-all duration-300 group-hover:scale-110 group-hover:text-blue-300 md:text-4xl"
						>
							50K+
						</div>
						<div class="text-gray-300 transition-colors duration-300 group-hover:text-gray-200">
							Happy Users
						</div>
					</div>
					<div class="group cursor-pointer text-center" in:scale={{ duration: 300, delay: 2200 }}>
						<div
							class="text-3xl font-bold text-indigo-400 transition-all duration-300 group-hover:scale-110 group-hover:text-indigo-300 md:text-4xl"
						>
							15+
						</div>
						<div class="text-gray-300 transition-colors duration-300 group-hover:text-gray-200">
							Cities
						</div>
					</div>
				</div>
			</div>
		{/if}
	</div>
</section>

<style>
	/* Custom pulse animation for floating elements */
	@keyframes float {
		0%,
		100% {
			transform: translateY(0px);
		}
		50% {
			transform: translateY(-10px);
		}
	}
</style>
