<!-- BaseField.svelte -->
<script lang="ts">
	export let label: string = '';
	export let id: string = '';
	export let error: string = '';
	export let required: boolean = false;
	export let disabled: boolean = false;
	export let description: string = '';
</script>

<div class="flex w-full flex-col gap-2">
	{#if label}
		<label for={id} class="field-label" class:required>
			{label}
		</label>
	{/if}

	{#if description}
		<p class="text-muted-foreground -mt-1 text-xs">{description}</p>
	{/if}

	<div class="relative disabled:pointer-events-none" class:disabled>
		<slot />
	</div>

	{#if error}
		<span class="text-sm text-red-500">{error}</span>
	{/if}
</div>
