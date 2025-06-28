<script>
	import Footer from '$lib/components/layout/Footer.svelte';
	import Header from '$lib/components/layout/Header.svelte';

	export let data;
	let session = data || {};
</script>

<Header {session} />
<main>
	<slot />
</main>
<Footer />
