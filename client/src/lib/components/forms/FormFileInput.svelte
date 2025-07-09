<script lang="ts">
	import Label from '$lib/components/ui/label/label.svelte';

	export let field: string;
	export let label: string = '';
	export let form: Record<string, any>;
	export let errors: Record<string, string>;
	export let touched: Record<string, boolean>;

	function handleChange(event: Event) {
		const file = (event.target as HTMLInputElement).files?.[0];
		if (file) {
			form[field] = file;
		}
	}
</script>

<div class="space-y-1">
	<Label for={field} class="font-muted-foreground text-sm">{label || field}</Label>
	<input type="file" id={field} on:change={handleChange} />
	{#if touched[field] && errors[field]}
		<p class="text-sm text-red-500">{errors[field]}</p>
	{/if}
</div>
