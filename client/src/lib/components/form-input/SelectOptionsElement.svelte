<script lang="ts">
	import { ChevronDown, X } from '@lucide/svelte';

	// Types
	interface Option {
		value: string;
		label: string;
		disabled?: boolean;
	}

	// Props
	export let value: string = '';
	export let id: string = '';
	export let label: string = '';
	export let options: Option[] = [];
	export let placeholder: string = 'Select option';
	export let required: boolean = false;
	export let disabled: boolean = false;
	export let clearable: boolean = false;
	export let onchange: (value: string) => void = () => {};

	// Reactive display text
	$: displayText = (() => {
		if (value) {
			const selectedOption = options.find((option) => option.value === value);
			return selectedOption ? selectedOption.label : placeholder;
		}
		return placeholder;
	})();

	// Handle change
	const handleChange = (event: Event) => {
		const target = event.target as HTMLSelectElement;
		value = target.value;
		onchange(value);
	};

	// Clear selection
	const clearSelection = (event: Event) => {
		event.stopPropagation();
		value = '';
		onchange('');
	};

	// Focus management
	let selectElement: HTMLSelectElement;

	const focusSelect = () => {
		if (!disabled) {
			selectElement?.focus();
		}
	};
</script>

<div class="relative">
	{#if label}
		<label for={id} class="mb-2 block text-sm font-medium text-gray-700">
			{label}
			{#if required}<span class="ml-1 text-red-500">*</span>{/if}
		</label>
	{/if}

	<!-- Custom Select Container -->
	<div
		class="relative flex h-12 w-full items-center justify-between rounded-lg border border-gray-300 bg-white px-3 py-2 text-left text-sm transition-colors focus-within:border-blue-500 focus-within:ring-2 focus-within:ring-blue-500/20 hover:border-blue-400"
		class:opacity-50={disabled}
		class:cursor-not-allowed={disabled}
	>
		<!-- Native Select - FIXED: adjusted positioning to avoid covering clear button -->
		<select
			bind:this={selectElement}
			{id}
			bind:value
			on:change={handleChange}
			{disabled}
			class="absolute inset-0 w-full cursor-pointer opacity-0"
			class:cursor-not-allowed={disabled}
			style="right: {clearable && value && !disabled ? '40px' : '0'};"
		>
			{#if placeholder}
				<option value="" disabled>{placeholder}</option>
			{/if}
			{#each options as option}
				<option value={option.value} disabled={option.disabled}>
					{option.label}
				</option>
			{/each}
		</select>

		<!-- Custom Display -->
		<div class="pointer-events-none flex flex-1 items-center gap-2">
			<span class={value ? 'text-gray-900' : 'text-gray-500'}>
				{displayText}
			</span>
		</div>

		<!-- Icons -->
		<div class="relative z-10 flex items-center gap-1">
			{#if clearable && value && !disabled}
				<button
					type="button"
					on:click={clearSelection}
					class="rounded p-1 transition-colors hover:bg-gray-100 focus:outline-none"
					aria-label="Clear selection"
				>
					<X class="h-4 w-4 text-gray-500 hover:text-gray-700" />
				</button>
			{/if}

			<div class="pointer-events-none">
				<ChevronDown class="h-4 w-4 text-gray-500" />
			</div>
		</div>
	</div>

	<!-- Error State (optional) -->
	{#if required && !value}
		<p class="mt-1 text-xs text-red-600">This field is required</p>
	{/if}
</div>
