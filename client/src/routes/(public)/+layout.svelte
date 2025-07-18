<!-- /(public)/+layout.svelte -->
<script lang="ts">
	import { authStore } from '$lib/stores/auth.store';
	import Footer from '$lib/components/layout/Footer.svelte';
	import Header from '$lib/components/layout/Header.svelte';

	export let data;
	export let children;

	authStore.setUser(data?.user || null);
</script>

<Header />
{@render children?.()}
<Footer />
