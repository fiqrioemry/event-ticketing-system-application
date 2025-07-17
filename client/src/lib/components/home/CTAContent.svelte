<script>
	import { goto } from '$app/navigation';
</script>

<section class="relative overflow-hidden py-20">
	<div class="absolute inset-0 z-0">
		<img src="banner2.webp" alt="Event Background" class="h-full w-full object-cover" />
		<div class="absolute inset-0 bg-gradient-to-r from-blue-900/90 to-indigo-900/90"></div>
	</div>

	<div class="relative z-10 container mx-auto px-4 text-center text-white">
		<h2 class="mb-6 text-4xl font-bold md:text-5xl">
			Ready to Start Your
			<span class="bg-gradient-to-r from-indigo-400 to-blue-400 bg-clip-text text-transparent">
				Adventure?
			</span>
		</h2>
		<p class="mx-auto mb-8 max-w-2xl text-xl text-gray-200">
			Join thousands of event-goers who trust EventTix for their entertainment needs
		</p>
		<button
			on:click={() => goto('/events')}
			class="inline-block rounded-xl bg-gradient-to-r from-indigo-500 to-blue-600 px-12 py-4 text-xl font-semibold text-white transition-all duration-300 hover:scale-105 hover:from-indigo-600 hover:to-blue-700 hover:shadow-2xl"
		>
			Browse Events Now
		</button>
	</div>
</section>
