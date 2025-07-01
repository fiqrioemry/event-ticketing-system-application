<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
</script>

<svelte:head>
	<title>404 - Page Not Found</title>
</svelte:head>

<div class="flex min-h-screen flex-col items-center justify-center p-8 text-center">
	<h1 class="text-5xl font-bold">404</h1>
	<p class="mt-4 text-lg">Oops! The page you are looking for does not exist.</p>
	<Button href="/" class="mt-4">Go back home</Button>
</div>
