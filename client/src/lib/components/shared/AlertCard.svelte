<script lang="ts">
	import { AlertCircle, RefreshCcw } from '@lucide/svelte';

	export let onAction: () => void;
	export let message: string | null = 'An error occurred has occurred. Please try again later.';
</script>

<div class="mb-6 rounded-lg border border-red-200 bg-red-50 p-6">
	<div class="flex items-center">
		<AlertCircle class="mr-3 h-5 w-5 text-red-400" />
		<div>
			<h3 class="text-sm font-medium text-red-800">Error</h3>
			<p class="mt-1 text-sm text-red-700">{message}</p>
		</div>
	</div>
	<button
		on:click={onAction}
		class="mt-4 inline-flex items-center rounded-md bg-red-100 px-4 py-2 text-sm font-medium text-red-800 transition-colors hover:bg-red-200"
	>
		<RefreshCcw class="mr-2 h-4 w-4" />
		Try Again
	</button>
</div>
