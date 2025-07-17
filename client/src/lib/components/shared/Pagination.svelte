<script lang="ts">
	import { ArrowLeft, ArrowRight } from '@lucide/svelte';
	import Button from '$lib/components/ui/button/button.svelte';

	// props from parent component
	export let totalPages: number = 1;
	export let currentPage: number = 1;
	export let onPageChange: (page: number) => void = () => {};

	// optiosn for custom
	export let showInfo: boolean = true;
	export let total: number | null = null;
	export let perPage: number | null = null;

	$: canGoPrevious = currentPage > 1;
	$: canGoNext = currentPage < totalPages;

	function handlePrevious() {
		if (canGoPrevious) {
			onPageChange(currentPage - 1);
		}
	}

	function handleNext() {
		if (canGoNext) {
			onPageChange(currentPage + 1);
		}
	}
</script>

{#if totalPages > 1}
	<div class="flex items-center justify-between border-t border-gray-200 px-6 py-4">
		<!-- Info Section -->
		{#if showInfo}
			<div class="text-muted-foreground text-sm">
				{#if total && perPage}
					Showing {(currentPage - 1) * perPage + 1} to {Math.min(currentPage * perPage, total)} of {total}
					entries
				{:else}
					Page {currentPage} of {totalPages}
				{/if}
			</div>
		{:else}
			<div></div>
		{/if}

		<!-- Pagination Controls -->
		<div class="flex items-center space-x-2">
			<Button size="sm" variant="outline" onclick={handlePrevious} disabled={!canGoPrevious}>
				<ArrowLeft class="h-4 w-4" />
			</Button>

			<span class="text-muted-foreground text-sm">
				{currentPage} / {totalPages}
			</span>

			<Button size="sm" variant="outline" onclick={handleNext} disabled={!canGoNext}>
				<ArrowRight class="h-4 w-4" />
			</Button>
		</div>
	</div>
{/if}
