<script lang="ts">
	import { Search, X } from '@lucide/svelte';
	import { onDestroy } from 'svelte';

	// Props
	export let value: string = '';
	export let debounceMs: number = 500;
	export let disabled: boolean = false;
	export let className: string = '';
	export let size: 'sm' | 'md' | 'lg' = 'md';
	export let showClearButton: boolean = true;
	export let placeholder: string = 'Search...';
	export let onSearch: (value: string) => void = () => {};
	export let onKeyDown: (event: KeyboardEvent) => void = () => {};
	export let onClear: (value: string) => void = () => {}; // ← New prop untuk clear action

	// State
	let debounceTimer: number | null = null;
	let previousValue = value;

	// Size classes
	const sizeClasses = {
		sm: 'h-9 px-3 text-sm',
		md: 'h-11 px-4 text-base',
		lg: 'h-12 px-4 text-lg'
	};

	const iconSizes = {
		sm: 'h-4 w-4',
		md: 'h-5 w-5',
		lg: 'h-6 w-6'
	};

	// Watch for external value changes (dari parent)
	$: if (value !== previousValue) {
		previousValue = value;
		// Jika value di-set dari luar (misal reset), jangan trigger onSearch
	}

	// Debounced search function
	function debouncedSearch(searchTerm: string) {
		if (debounceTimer) {
			clearTimeout(debounceTimer);
		}

		debounceTimer = setTimeout(() => {
			onSearch(searchTerm);
		}, debounceMs);
	}

	// Handle input
	function handleInput(event: Event) {
		const input = event.target as HTMLInputElement;
		const newValue = input.value;
		value = newValue;
		previousValue = newValue;
		debouncedSearch(newValue);
	}

	// Handle clear - use separate callback
	function handleClear() {
		value = '';
		previousValue = '';

		// Use onClear if provided, otherwise fallback to onSearch
		if (onClear && onClear !== (() => {})) {
			onClear(value);
		} else {
			onSearch(value);
		}
	}

	// Cleanup
	onDestroy(() => {
		if (debounceTimer) {
			clearTimeout(debounceTimer);
		}
	});
</script>

<div class="relative flex items-center gap-2 {className}">
	<label for="search" class="flex-shrink-0">
		<Search
			class="{iconSizes[
				size
			]} absolute top-1/2 left-5 h-5 w-5 -translate-y-1/2 transform cursor-pointer text-gray-400"
		/>
	</label>

	<div class="relative flex-1">
		<input
			id="search"
			type="text"
			{placeholder}
			{disabled}
			bind:value
			on:input={handleInput}
			on:keydown={onKeyDown}
			class="
        h-12 w-full rounded-md border border-gray-300 pl-10
        {sizeClasses[size]}
        {showClearButton && value ? 'pr-10' : ''}
        focus:ring-2 focus:ring-indigo-500 focus:outline-none
        disabled:cursor-not-allowed disabled:opacity-50
      "
		/>

		{#if showClearButton && value && !disabled}
			<button
				type="button"
				class="absolute inset-y-0 right-0 flex items-center rounded-r-md pr-3 hover:bg-gray-50"
				on:click={handleClear}
			>
				<X class="{iconSizes[size]} text-gray-400 hover:text-gray-600" />
			</button>
		{/if}
	</div>
</div>
