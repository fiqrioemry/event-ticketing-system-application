<!-- src/error.svelte -->
<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Home, ArrowLeft } from '@lucide/svelte';

	export let status: number;
	export let message: string;

	// Determine error type
	$: isNotFound = status === 404;

	function goBack() {
		if (typeof window !== 'undefined') {
			window.history.back();
		}
	}
</script>

<svelte:head>
	<title>{status || 404} - {message || 'Page not found'}</title>
	<meta name="description" content="Error {status}: {message}" />
</svelte:head>

<div
	class="flex min-h-screen items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100"
>
	<div class="mx-auto w-full max-w-md px-6 py-8 text-center">
		<div class="mb-8">
			<h1 class="mb-10 text-6xl font-bold text-gray-800">404</h1>
			<p class="mb-2 text-gray-600">The page you're looking for doesn't exist or has been moved.</p>
		</div>
		<!-- Action Buttons -->
		<div class="space-y-3">
			<Button href="/" variant="event-cyan" class="w-full">
				<Home class="mr-2 h-4 w-4" />
				Go home
			</Button>

			<Button variant="event-outline" onclick={goBack} class="w-full">
				<ArrowLeft class="mr-2 h-4 w-4" />
				Go Back
			</Button>

			{#if isNotFound}
				<Button variant="ghost" href="/sitemap" class="w-full">View Sitemap</Button>
			{/if}
		</div>
	</div>
</div>
