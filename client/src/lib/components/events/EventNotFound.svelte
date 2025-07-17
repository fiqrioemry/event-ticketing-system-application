<script lang="ts">
	import { Search, RefreshCw } from '@lucide/svelte';

	export let resetFilters: () => void = () => {};
</script>

<div class="rounded-xl border border-gray-200 bg-white shadow-sm">
	<div class="px-6 py-16 text-center">
		<!-- No Results with Filters -->
		<div
			class="mx-auto mb-6 flex h-24 w-24 items-center justify-center rounded-full bg-gradient-to-r from-blue-100 to-indigo-100"
		>
			<Search class="h-10 w-10 text-blue-500" />
		</div>

		<h3 class="mb-4 text-2xl font-bold text-gray-800">No events match your filters</h3>

		<p class="mx-auto mb-8 max-w-md leading-relaxed text-gray-600">
			We couldn't find any events matching your current search criteria. Try adjusting your filters
			or clearing them to see all available events.
		</p>

		<div class="flex flex-col justify-center gap-4 sm:flex-row">
			<button
				on:click={resetFilters}
				class="inline-flex items-center gap-2 rounded-lg bg-gradient-to-r from-blue-600 to-indigo-600 px-6 py-3 font-medium text-white transition-all hover:from-blue-700 hover:to-indigo-700 hover:shadow-lg"
			>
				<RefreshCw class="h-4 w-4" />
				Clear All Filters
			</button>
		</div>
	</div>
</div>
