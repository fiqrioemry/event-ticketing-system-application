<!-- src/lib/components/layout/Header.svelte -->
<script lang="ts">
	import AppLogo from '$lib/components/common/AppLogo.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import UserMenu from '$lib/components/common/UserMenu.svelte';

	export let session: any = null;

	const user = session?.user || null;
</script>

<header class="border-b py-2">
	<div class="mx-auto flex max-w-7xl items-center justify-between px-4">
		<AppLogo />
		<nav class="flex items-center gap-4">
			<a href="/" class="text-md">Home</a>
			<a href="/events" class="text-md">Event</a>
			<a href="/about" class="text-md">About</a>
		</nav>
		<div>
			{#if user || user?.length > 0}
				<UserMenu {user} />
			{:else}
				<Button class="rounded-full" href="/signin">LOGIN</Button>
			{/if}
		</div>
	</div>
</header>
