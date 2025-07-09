<script lang="ts">
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';

	export let field: string;
	export let label: string = '';
	export let placeholder = '';
	export let type: string = 'text';

	export let form: Record<string, any>;
	export let errors: Record<string, string>;
	export let touched: Record<string, boolean>;
</script>

<div class="space-y-1">
	<Label class="text-muted-foreground text-sm font-medium capitalize" for={field}
		>{label || field}</Label
	>
	<Input bind:value={form[field]} {placeholder} {type} />
	{#if touched[field] && errors[field]}
		<p class="text-sm text-red-500">{errors[field]}</p>
	{/if}
</div>
