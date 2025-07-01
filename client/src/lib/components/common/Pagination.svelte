<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import { ArrowLeft, ArrowRight } from '@lucide/svelte';
	import Button from '$lib/components/ui/button/button.svelte';

	export let page: number;
	export let totalPages: number;

	const dispatch = createEventDispatcher();

	function goToPage(newPage: number) {
		if (newPage >= 1 && newPage <= totalPages) {
			dispatch('change', newPage);
		}
	}
</script>

<div class="mt-6 flex items-center justify-center gap-4 text-sm">
	<Button
		variant="outline"
		disabled={page <= 1}
		onclick={() => goToPage(page - 1)}
		class="flex items-center gap-1"
	>
		<span class="text-xs"><ArrowLeft /></span> Prev
	</Button>

	<div>
		Page {page} of {totalPages}
	</div>

	<Button
		variant="outline"
		disabled={page >= totalPages}
		onclick={() => goToPage(page + 1)}
		class="flex items-center gap-1"
	>
		Next <span class="text-xs"><ArrowRight /></span>
	</Button>
</div>
