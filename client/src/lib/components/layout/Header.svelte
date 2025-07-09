<!-- src/lib/components/layout/Header.svelte -->
<script lang="ts">
	import { page } from '$app/state';
	import Button from '$lib/components/ui/button/button.svelte';
	let user = null;
</script>

<header class="border-b py-2">
	<div class="mx-auto flex max-w-7xl items-center justify-between px-4">
		<h1>Logo</h1>
		<nav class="flex items-center gap-4">
			<a href="/" class="text-md font-bold" class:text-blue-500={page.url.pathname === '/'}>Home</a>
			<a
				href="/events"
				class="text-md font-bold"
				class:text-blue-500={page.url.pathname === '/events'}>Event</a
			>
		</nav>
		<div>
			{#if user}
				<span>usermenu</span>
			{:else}
				<Button class="rounded-full bg-blue-500" href="/signin">LOGIN</Button>
			{/if}
		</div>
	</div>
</header>
