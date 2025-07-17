<!-- src/lib/components/TextInput.svelte -->
<script lang="ts">
	import BaseField from '$lib/components/form-input/BaseField.svelte';

	export let id: string = '';
	export let type: 'text' | 'tel' = 'text';
	export let label: string = '';
	export let error: string = '';
	export let value: string = '';
	export let description: string = '';
	export let placeholder: string = '';
	export let readonly: boolean = false;
	export let required: boolean = false;
	export let disabled: boolean = false;
	export let pattern: string | null = null;
	export let maxlength: number | null = null;
	export let minlength: number | null = null;

	let inputElement: HTMLInputElement;
</script>

<BaseField {label} {error} {required} {disabled} {id} {description}>
	<div class="relative">
		<input
			{id}
			{type}
			bind:value
			{disabled}
			{readonly}
			{maxlength}
			{minlength}
			{pattern}
			class:error
			{placeholder}
			class=" focus:ring-ring bg-background border-border text-foreground placeholder-muted-foreground w-full rounded-lg border px-3 py-2 transition-all duration-200 focus:border-transparent focus:ring-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-60"
			bind:this={inputElement}
		/>
	</div>
</BaseField>
