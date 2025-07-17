<!-- src/lib/components/TextArea.svelte -->
<script>
	import BaseField from '$lib/components/form-input/BaseField.svelte';

	export let label = '';
	export let error = '';
	export let required = false;
	export let disabled = false;
	export let id = '';
	export let description = '';
	export let value = '';
	export let rows = 4;
	export let maxlength = 500;
	export let minlength = 10;
	export let readonly = false;
	export let autoResize = false;
	export let showCharCount = true;
	export let placeholder = 'Write something...';

	/**
	 * @type {HTMLTextAreaElement}
	 */
	let textareaElement;

	$: charCount = value ? value.length : 0;
	$: isOverLimit = maxlength && charCount > maxlength;

	function resizeTextarea() {
		if (!textareaElement || !autoResize) return;
		textareaElement.style.height = 'auto';
		textareaElement.style.height = textareaElement.scrollHeight + 'px';
	}

	// Auto-resize on value change if enabled
	$: if (autoResize && value !== undefined) {
		setTimeout(resizeTextarea, 0);
	}
</script>

<BaseField {label} {error} {required} {disabled} {id} {description}>
	<div class="relative">
		<textarea
			bind:this={textareaElement}
			bind:value
			{id}
			{rows}
			{disabled}
			{readonly}
			{maxlength}
			{minlength}
			class:error
			{placeholder}
			on:input={resizeTextarea}
			class:auto-resize={autoResize}
			class="border-border bg-background text-foreground placeholder-muted-foreground focus:ring-ring min-h-20 w-full resize-none rounded-lg border px-3 py-2 transition-all duration-200 focus:border-transparent focus:ring-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-60"
		></textarea>

		{#if showCharCount}
			<div
				class="text-muted-foreground over-limit:text-red-500 absolute right-3 bottom-2 text-xs"
				class:over-limit={isOverLimit}
			>
				{charCount}{#if maxlength}/{maxlength}{/if}
			</div>
		{/if}
	</div>
</BaseField>
