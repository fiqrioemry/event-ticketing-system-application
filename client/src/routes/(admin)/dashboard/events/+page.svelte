<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';

	export let data;
	const { events, pagination } = data || {};
</script>

<svelte:head>
	<title>Event | Tiketku</title>
	<meta
		name="description"
		content="Manage your events effectively with Tiketku's event dashboard."
	/>
</svelte:head>

<section class="space-y-6 px-4 py-10">
	<div class="text-center text-3xl font-bold">
		<h1>Event Lists</h1>
		<span>Manage events effectively with Tiketku's event dashboard.</span>
	</div>

	{#if !events}
		<div class="text-center">
			<p>No events found.</p>
		</div>
	{:else}
		<div>
			{#each events as event}
				<div>{event.title}</div>
			{/each}
		</div>
	{/if}
</section>
