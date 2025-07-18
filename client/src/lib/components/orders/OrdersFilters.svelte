<script lang="ts">
	import { statusOptions } from '$lib/utils/state.js';
	import SelectOptionsElement from '$lib/components/form-input/SelectOptionsElement.svelte';
	import SearchInputElement from '$lib/components/form-input/SearchInputElement.svelte';

	// Props
	export let searchQuery: string = '';
	export let statusFilter: string = '';
	export let hasActiveFilters: boolean = false;

	// Events
	export let onSearch: () => void;
	export let onFilterChange: () => void;
	export let onReset: () => void;
</script>

<div class="mt-6 space-y-4">
	<div class="flex flex-col gap-4 lg:flex-row lg:items-center lg:justify-between">
		<!-- Search -->
		<div class="flex-1 lg:max-w-md">
			<SearchInputElement
				bind:value={searchQuery}
				onClear={onReset}
				onKeyDown={(e) => {
					if (e.key === 'Enter') {
						onSearch();
					}
				}}
			/>
		</div>

		<!-- Filter Controls -->
		<div class="flex items-center gap-3">
			<!-- Status Filter -->
			<SelectOptionsElement
				id="status-filter"
				clearable={true}
				bind:value={statusFilter}
				onchange={onFilterChange}
				options={statusOptions}
				placeholder="All Statuses"
			/>

			<!-- Reset Filters -->
			{#if hasActiveFilters}
				<button on:click={onReset} class="text-sm text-blue-600 hover:text-blue-800">
					Reset
				</button>
			{/if}
		</div>
	</div>
</div>
