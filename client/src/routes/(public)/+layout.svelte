<!-- src/routes/(public)/+layout.svelte -->
<script lang="ts">
	import { authStore } from '$lib/stores/auth.store';
	import Footer from '$lib/components/layout/Footer.svelte';
	import Header from '$lib/components/layout/Header.svelte';

	const { children, data } = $props<{ children: any; data: any }>();
	console.log('Refresh Token:', data);
	$effect(() => {
		authStore.setUser(data?.user || null);
	});
</script>

<Header />
{@render children?.()}
<Footer />
