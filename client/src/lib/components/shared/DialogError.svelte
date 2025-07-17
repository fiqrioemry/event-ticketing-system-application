<script lang="ts">
	import { AlertCircle } from '@lucide/svelte';
	import Button from '$lib/components/ui/button/button.svelte';

	export let title = 'Asset Not Found';
	export let description = "The asset you're looking for doesn't exist or has been removed.";
	export let onClose: () => void = () => {};
</script>

<div class="flex items-center justify-center p-12">
	<!-- Error State -->
	<div class="space-y-4 text-center">
		<AlertCircle class="mx-auto h-16 w-16 text-red-400" />
		<div>
			<h3 class="text-lg font-semibold text-gray-900">{title}</h3>
			<p class="mt-1 text-sm text-gray-500">{description}</p>
		</div>
		<Button variant="outline" onclick={onClose}>Go back</Button>
	</div>
</div>
