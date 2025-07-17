<script>
	import { Loader2 } from '@lucide/svelte';

	export let loadingText = 'Loading Asset ...';
</script>

<div class="flex h-full items-center justify-center p-12">
	<div class="space-y-4">
		<div class="flex items-center justify-center">
			<Loader2 class="h-12 w-12 animate-spin" />
		</div>
		<p class="text-muted-foreground text-sm">{loadingText}</p>
	</div>
</div>
