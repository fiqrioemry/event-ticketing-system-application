<script lang="ts">
	export let message: string = 'Loading...';
	export let showLogo: boolean = true;
	export let subtitle: string = 'Please wait while we prepare your experience';
</script>

<div class="fixed inset-0 z-50 flex items-center justify-center bg-white">
	<div class="text-center">
		{#if showLogo}
			<div class="relative mb-6">
				<div
					class="mx-auto flex h-16 w-16 animate-pulse items-center justify-center rounded-xl bg-gradient-to-r from-blue-600 to-indigo-600"
				>
					<span class="text-xl font-bold text-white">T</span>
				</div>
				<div
					class="absolute inset-0 mx-auto h-16 w-16 animate-spin rounded-xl border-4 border-blue-200 border-t-blue-600"
				></div>
			</div>
		{/if}

		<h3 class="mb-2 text-xl font-semibold text-gray-800">{message}</h3>
		<p class="mb-6 text-gray-600">{subtitle}</p>

		<!-- Animated dots -->
		<div class="flex justify-center space-x-1">
			{#each [0, 1, 2] as i}
				<div
					class="h-2 w-2 animate-bounce rounded-full bg-blue-600"
					style="animation-delay: {i * 0.15}s"
				></div>
			{/each}
		</div>
	</div>

	<!-- Background decoration -->
	<div class="absolute top-10 right-10 h-32 w-32 rounded-full bg-blue-200/30 blur-xl"></div>
	<div class="absolute bottom-10 left-10 h-24 w-24 rounded-full bg-indigo-200/30 blur-lg"></div>
</div>
