<script>
	import { goto } from '$app/navigation';
</script>

<section class="relative flex min-h-screen items-center justify-center overflow-hidden">
	<!-- Background with overlay -->
	<div class="absolute inset-0 z-0">
		<img src="banner1.webp" alt="Event Background" class="h-full w-full object-cover" />
		<div
			class="absolute inset-0 bg-gradient-to-r from-blue-900/80 via-blue-800/70 to-indigo-800/80"
		></div>
	</div>

	<!-- Floating Elements -->
	<div class="absolute top-20 left-10 h-20 w-20 animate-pulse rounded-full bg-white/10"></div>
	<div
		class="absolute top-40 right-20 h-16 w-16 animate-pulse rounded-full bg-indigo-500/20"
		style="animation-delay: 2s;"
	></div>
	<div
		class="absolute bottom-40 left-20 h-12 w-12 animate-pulse rounded-full bg-blue-500/20"
		style="animation-delay: 4s;"
	></div>

	<!-- Hero Content -->
	<div class="relative z-10 container mx-auto px-4 text-center text-white">
		<div class="mx-auto max-w-4xl">
			<h1 class="animate-fade-in-down mb-6 text-5xl font-bold md:text-7xl">
				Discover
				<span class="bg-gradient-to-r from-indigo-400 to-blue-400 bg-clip-text text-transparent">
					Amazing
				</span>
				Events
			</h1>

			<p class="animate-fade-in-up mb-8 text-xl text-gray-200 md:text-2xl">
				Your gateway to unforgettable experiences. Find, book, and enjoy the best events in your
				city.
			</p>

			<div class="animate-fade-in-up flex flex-col justify-center gap-4 sm:flex-row">
				<button
					on:click={() => goto('/events')}
					class="rounded-xl bg-gradient-to-r from-indigo-500 to-blue-600 px-8 py-4 text-lg font-semibold text-white transition-all duration-300 hover:scale-105 hover:from-indigo-600 hover:to-blue-700 hover:shadow-2xl"
				>
					Explore Events
				</button>
				<a
					href="#features"
					class="rounded-xl border-2 border-white/30 px-8 py-4 text-lg font-semibold text-white transition-all duration-300 hover:border-white/60 hover:bg-white/10"
				>
					Learn More
				</a>
			</div>

			<!-- Stats -->
			<div class="animate-fade-in-up mt-16 grid grid-cols-3 gap-8">
				<div class="text-center">
					<div class="text-3xl font-bold text-indigo-400 md:text-4xl">1000+</div>
					<div class="text-gray-300">Events</div>
				</div>
				<div class="text-center">
					<div class="text-3xl font-bold text-blue-400 md:text-4xl">50K+</div>
					<div class="text-gray-300">Happy Users</div>
				</div>
				<div class="text-center">
					<div class="text-3xl font-bold text-indigo-400 md:text-4xl">15+</div>
					<div class="text-gray-300">Cities</div>
				</div>
			</div>
		</div>
	</div>

	<!-- Scroll Indicator -->
	<div class="absolute bottom-8 left-1/2 -translate-x-1/2 transform animate-bounce">
		<div class="flex h-10 w-6 justify-center rounded-full border-2 border-white/50">
			<div class="mt-2 h-3 w-1 animate-pulse rounded-full bg-white/70"></div>
		</div>
	</div>
</section>
