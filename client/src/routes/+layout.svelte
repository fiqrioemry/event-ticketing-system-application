<!-- src/routes/+layout.svelte -->
<script lang="ts">
	import '../app.css';
	import { onMount } from 'svelte';
	import { Toaster } from '$lib/components/ui/sonner/index.js';
	import LayoutLoading from '$lib/components/shared/LayoutLoading.svelte';

	let isLoading = true;

	onMount(() => {
		setTimeout(() => {
			isLoading = false;
		}, 150); //
	});
	export let children;
</script>

{#if isLoading}
	<LayoutLoading />
{:else}
	<Toaster position="top-right" richColors theme="light" />
	{@render children?.()}
{/if}
