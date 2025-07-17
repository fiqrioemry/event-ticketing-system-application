<script lang="ts">
	import { X } from '@lucide/svelte';
	export let message: string | null = '';
	export let onclearError: () => void = () => {};
</script>

<div class="bg-destructive/10 border-destructive/20 rounded-lg border p-4">
	<div class="flex items-center justify-between">
		<span class="text-destructive flex-1 text-sm">{message}</span>
		<button
			type="button"
			on:click={onclearError}
			class="text-destructive hover:text-destructive/80 focus:ring-destructive/20 -m-1 rounded p-1 transition-colors focus:ring-2 focus:outline-none"
		>
			<X size={16} />
		</button>
	</div>
</div>
