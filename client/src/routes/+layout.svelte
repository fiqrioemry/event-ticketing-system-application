<script lang="ts">
	import '../app.css';
	import { onMount } from 'svelte';
	import { authStore } from '$lib/stores/auth.store';
	import * as authService from '$lib/services/auth.service';
	import { Toaster } from '$lib/components/ui/sonner/index.js';
	import LayoutLoading from '$lib/components/shared/LayoutLoading.svelte';

	let isLoading = true;

	// Simulate loading delay (authentication already handled on layout.server)
	onMount(async () => {
		try {
			// Fetch user data to ensure authStore is populated
			const response = await authService.refreshToken();
			authStore.setUser(response?.data || null);
		} catch (error) {
			authStore.setUser(null);
		} finally {
			setTimeout(() => {
				isLoading = false;
			}, 100); //
		}
	});
</script>

{#if isLoading}
	<LayoutLoading />
{:else}
	<Toaster position="top-right" richColors theme="light" />
	<slot />
{/if}
