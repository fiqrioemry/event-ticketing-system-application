<script lang="ts">
	import * as Select from '$lib/components/ui/select';
	import Label from '$lib/components/ui/label/label.svelte';

	export let field: string;
	export let label: string = '';
	export let placeholder = '';
	export let form: Record<string, any>;
	export let errors: Record<string, string>;
	export let touched: Record<string, boolean>;
	export let options: { label: string; value: string }[] = [];
</script>

<div class="space-y-1">
	<Label for={field}>{label || field}</Label>

	<Select.Root bind:value={form[field]} type="single">
		<Select.Trigger>
			{form[field] ? form[field] : placeholder}
		</Select.Trigger>
		<Select.Content class="w-full">
			{#each options as opt}
				<Select.Item value={opt.value}>{opt.label}</Select.Item>
			{/each}
		</Select.Content>
	</Select.Root>

	{#if touched[field] && errors[field]}
		<p class="text-sm text-red-500">{errors[field]}</p>
	{/if}
</div>
