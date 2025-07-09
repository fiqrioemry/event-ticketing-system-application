<script lang="ts">
	import Label from '$lib/components/ui/label/label.svelte';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';

	export let field: string;
	export let label: string = '';
	export let placeholder = '';
	export let form: Record<string, any>;
	export let errors: Record<string, string>;
	export let touched: Record<string, boolean>;
	export let maxLength: number = 200;
	export let maxRows: number = 5;
	export let customClass: string = 'h-24';
</script>

<div class="space-y-1">
	<Label for={field} class="font-muted-foreground text-sm">
		{label || field}
	</Label>
	<Textarea
		bind:value={form[field]}
		{placeholder}
		id={field}
		maxlength={maxLength}
		rows={maxRows}
		class={`${customClass} resize-none`}
	/>
	{#if touched[field] && errors[field]}
		<p class="text-sm text-red-500">{errors[field]}</p>
	{/if}
</div>
