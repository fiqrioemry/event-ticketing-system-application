<script>
	import { onMount } from 'svelte';
	import { fly, fade, scale } from 'svelte/transition';
	import { Search, Lock, Heart } from '@lucide/svelte';

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
			{ threshold: 0.2 }
		);

		const section = document.getElementById('features');
		if (section) {
			observer.observe(section);
		}

		return () => observer.disconnect();
	});

	const features = [
		{
			icon: Search,
			title: 'Easy Discovery',
			description:
				'Find events that match your interests with our smart search and recommendation system.',
			delay: 200
		},
		{
			icon: Lock,
			title: 'Secure Booking',
			description:
				'Your transactions are protected with bank-level security and instant confirmation.',
			delay: 400
		},
		{
			icon: Heart,
			title: 'Amazing Experience',
			description: 'From booking to event day, we ensure you have the best possible experience.',
			delay: 600
		}
	];
</script>

<section id="features" class="overflow-hidden bg-white py-20">
	<div class="container mx-auto px-4">
		{#if mounted}
			<!-- Header with staggered animation -->
			<div class="mb-16 text-center">
				{#if sectionVisible}
					<h2
						class="mb-4 bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-4xl font-bold text-transparent md:text-5xl"
						in:fly={{ y: -30, duration: 800, delay: 0 }}
					>
						Why Choose TiketKu?
					</h2>
					<p
						class="mx-auto max-w-2xl text-xl text-gray-600"
						in:fade={{ duration: 800, delay: 200 }}
					>
						We make discovering and booking events simple, secure, and enjoyable
					</p>
				{/if}
			</div>

			<!-- Features Grid with enhanced animations -->
			<div class="grid gap-8 md:grid-cols-3">
				{#each features as feature, index}
					{#if sectionVisible}
						<div
							class="group cursor-pointer rounded-2xl p-8 text-center transition-all duration-500 hover:-translate-y-2 hover:bg-gradient-to-br hover:from-blue-50 hover:to-indigo-50 hover:shadow-xl"
							in:fly={{ y: 50, duration: 700, delay: feature.delay }}
						>
							<!-- Icon with enhanced animations -->
							<div
								class="mx-auto mb-6 flex h-16 w-16 items-center justify-center rounded-full bg-gradient-to-r from-blue-500 to-indigo-500 transition-all duration-500 group-hover:scale-125 group-hover:rotate-12 group-hover:shadow-lg"
								in:scale={{ duration: 600, delay: feature.delay + 200 }}
							>
								<svelte:component
									this={feature.icon}
									class="h-8 w-8 text-white transition-transform duration-300 group-hover:scale-110"
								/>
							</div>

							<!-- Title with entrance animation -->
							<h3
								class="mb-4 text-2xl font-bold text-gray-800 transition-colors duration-300 group-hover:text-blue-600"
								in:fade={{ duration: 600, delay: feature.delay + 400 }}
							>
								{feature.title}
							</h3>

							<!-- Description with staggered entrance -->
							<p
								class="leading-relaxed text-gray-600 transition-colors duration-300 group-hover:text-gray-700"
								in:fade={{ duration: 600, delay: feature.delay + 600 }}
							>
								{feature.description}
							</p>

							<!-- Decorative element that appears on hover -->
							<div
								class="mx-auto mt-6 h-1 w-0 rounded-full bg-gradient-to-r from-blue-500 to-indigo-500 transition-all duration-500 group-hover:w-12"
							></div>
						</div>
					{/if}
				{/each}
			</div>

			<!-- Additional floating elements for visual interest -->
			{#if sectionVisible}
				<div
					class="absolute top-20 left-10 h-12 w-12 animate-pulse rounded-full bg-blue-100 opacity-20"
					in:fade={{ duration: 800, delay: 1000 }}
				></div>
				<div
					class="absolute right-10 bottom-20 h-16 w-16 animate-pulse rounded-full bg-indigo-100 opacity-20"
					style="animation-delay: 2s;"
					in:fade={{ duration: 800, delay: 1200 }}
				></div>
				<div
					class="absolute top-1/2 left-5 h-8 w-8 animate-pulse rounded-full bg-purple-100 opacity-20"
					style="animation-delay: 1s;"
					in:fade={{ duration: 800, delay: 1400 }}
				></div>
			{/if}
		{/if}
	</div>
</section>

<style>
	/* Custom animation for floating elements */
	@keyframes gentle-float {
		0%,
		100% {
			transform: translateY(0px) rotate(0deg);
		}
		50% {
			transform: translateY(-8px) rotate(2deg);
		}
	}
</style>
