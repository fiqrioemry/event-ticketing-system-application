<script lang="ts">
	import '../app.css';
	import { onMount } from 'svelte';
	import { Toaster } from '$lib/components/ui/sonner/index.js';
	import LayoutLoading from '$lib/components/shared/LayoutLoading.svelte';

	let isLoading = true;

	// Simulate loading delay (authentication already handled on layout.server)
	onMount(() => {
		setTimeout(() => {
			isLoading = false;
		}, 100); //
	});
</script>

{#if isLoading}
	<LayoutLoading />
{:else}
	<Toaster position="top-right" richColors theme="light" />
	<slot />
{/if}
