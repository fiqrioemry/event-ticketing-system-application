<script>
	import { onMount } from 'svelte';
	import { fly, fade, scale } from 'svelte/transition';
	import { goto } from '$app/navigation';

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
			{ threshold: 0.3 }
		);

		const section = document.getElementById('cta-section');
		if (section) {
			observer.observe(section);
		}

		return () => observer.disconnect();
	});
</script>

<section id="cta-section" class="relative overflow-hidden py-20">
	<!-- Background with parallax effect -->
	<div class="absolute inset-0 z-0">
		<img
			src="banner2.webp"
			alt="Event Background"
			class="h-full w-full object-cover transition-transform duration-1000 {sectionVisible
				? 'scale-105'
				: 'scale-100'}"
		/>
		<div class="absolute inset-0 bg-gradient-to-r from-blue-900/90 to-indigo-900/90"></div>

		<!-- Enhanced floating elements -->
		{#if mounted && sectionVisible}
			<div
				class="absolute top-20 left-10 h-16 w-16 animate-pulse rounded-full bg-white/10"
				in:scale={{ duration: 600, delay: 200 }}
			></div>
			<div
				class="absolute top-40 right-20 h-12 w-12 animate-pulse rounded-full bg-indigo-400/20"
				style="animation-delay: 1s;"
				in:scale={{ duration: 600, delay: 400 }}
			></div>
			<div
				class="absolute bottom-32 left-20 h-20 w-20 animate-pulse rounded-full bg-blue-400/15"
				style="animation-delay: 2s;"
				in:scale={{ duration: 600, delay: 600 }}
			></div>
			<div
				class="absolute right-10 bottom-20 h-8 w-8 animate-pulse rounded-full bg-purple-400/20"
				style="animation-delay: 0.5s;"
				in:scale={{ duration: 600, delay: 800 }}
			></div>
		{/if}
	</div>

	<div class="relative z-10 container mx-auto px-4 text-center text-white">
		{#if mounted && sectionVisible}
			<!-- Enhanced title with word-by-word animation -->
			<h2 class="mb-6 text-4xl font-bold md:text-5xl">
				<span in:fly={{ y: -30, duration: 600, delay: 100 }}> Ready to Start Your </span>
				<br class="hidden sm:block" />
				<span
					class="inline-block bg-gradient-to-r from-indigo-400 to-blue-400 bg-clip-text text-transparent"
					in:fly={{ y: -30, duration: 600, delay: 300 }}
				>
					Adventure?
				</span>
			</h2>

			<!-- Subtitle with enhanced animation -->
			<p
				class="mx-auto mb-8 max-w-2xl text-xl leading-relaxed text-gray-200"
				in:fade={{ duration: 800, delay: 500 }}
			>
				Join thousands of event-goers who trust TiketKu for their entertainment needs
			</p>

			<!-- Enhanced CTA button with multiple effects -->
			<div class="inline-block" in:scale={{ duration: 600, delay: 700 }}>
				<button
					on:click={() => goto('/events')}
					class="group relative inline-block overflow-hidden rounded-xl bg-gradient-to-r from-indigo-500 to-blue-600 px-12 py-4 text-xl font-semibold text-white transition-all duration-500 hover:-translate-y-1 hover:scale-110 hover:from-indigo-600 hover:to-blue-700 hover:shadow-2xl"
				>
					<!-- Button content -->
					<span
						class="relative z-20 transition-transform duration-300 group-hover:translate-y-[-2px]"
					>
						Browse Events Now
					</span>

					<!-- Animated background overlays -->
					<div
						class="absolute inset-0 bg-gradient-to-r from-indigo-600 to-blue-700 opacity-0 transition-opacity duration-300 group-hover:opacity-100"
					></div>
					<div
						class="absolute inset-0 bg-white opacity-0 transition-opacity duration-300 group-hover:opacity-10"
					></div>

					<!-- Shine effect -->
					<div
						class="absolute inset-0 -translate-x-full bg-gradient-to-r from-transparent via-white/20 to-transparent transition-transform duration-1000 group-hover:translate-x-full"
					></div>

					<!-- Glow effect -->
					<div
						class="absolute -inset-1 -z-10 bg-gradient-to-r from-indigo-500 to-blue-600 opacity-0 blur transition-opacity duration-300 group-hover:opacity-50"
					></div>
				</button>
			</div>
		{/if}
	</div>
</section>

<style>
	/* Enhanced pulse animation with different timing */
	@keyframes gentle-pulse {
		0%,
		100% {
			opacity: 0.6;
			transform: scale(1);
		}
		50% {
			opacity: 1;
			transform: scale(1.05);
		}
	}
	/* Shine effect keyframes */
	@keyframes shine {
		0% {
			transform: translateX(-100%) skewX(-15deg);
		}
		100% {
			transform: translateX(200%) skewX(-15deg);
		}
	}
</style>
