<!-- src/routes/(public)/+page.svelte -->
<script lang="ts">
	import CtaContent from '$lib/components/home/CTAContent.svelte';
	import HeroContent from '$lib/components/home/HeroContent.svelte';
	import EventContent from '$lib/components/home/EventContent.svelte';
	import FeatureContent from '$lib/components/home/FeatureContent.svelte';
</script>

<svelte:head>
	<title>Tiketku - Book Amazing Events in Indonesia</title>
	<meta
		name="description"
		content="Book concerts, festivals, workshops & sports events. Best prices, secure payment, instant e-tickets. Join 100K+ happy customers!"
	/>
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<meta property="og:title" content="Tiketku - Book Amazing Events" />
	<meta
		property="og:description"
		content="Book concerts, festivals, workshops & sports events in Indonesia"
	/>
	<meta property="og:image" content="/og-tiketku.jpg" />
</svelte:head>

<!-- Hero Section -->
<HeroContent />

<!-- Features Section -->
<FeatureContent />

<!-- CTA Section -->
<CtaContent />

<!-- Featured Events Preview -->
<EventContent />

<style>
	@keyframes fade-in-down {
		0% {
			opacity: 0;
			transform: translateY(-30px);
		}
		100% {
			opacity: 1;
			transform: translateY(0);
		}
	}

	@keyframes fade-in-up {
		0% {
			opacity: 0;
			transform: translateY(30px);
		}
		100% {
			opacity: 1;
			transform: translateY(0);
		}
	}

	/* Smooth scrolling */
	:global(html) {
		scroll-behavior: smooth;
	}
</style>
