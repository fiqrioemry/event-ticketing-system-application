<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let value: string;
	export let options: { value: string; label: string }[] = [];

	const dispatch = createEventDispatcher();

	function handleChange(e: Event) {
		const target = e.target as HTMLSelectElement;
		dispatch('change', target.value);
	}
</script>

<select class="rounded-md border px-3 py-2" bind:value on:change={handleChange}>
	{#each options as opt}
		<option value={opt.value}>{opt.label}</option>
	{/each}
</select>
