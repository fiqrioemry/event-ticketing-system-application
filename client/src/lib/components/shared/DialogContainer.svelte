<!-- lib/components/shared/DialogContainer.svelte -->
<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog/index';

	export let open: boolean = true;
	export let onCloseDialog: () => void = () => {};
	export let size: 'default' | 'sm' | 'md' | 'lg' = 'default';

	const sizes = {
		default: 'w-60  p-0',
		sm: 'max-w-3xl max-h-[90vh] p-0',
		md: 'max-w-4xl max-h-[90vh] p-2',
		lg: 'max-w-5xl max-h-[90vh] p-4'
	};

	$: currentSize = sizes[size];
</script>

<Dialog.Root bind:open onOpenChange={(open) => !open && onCloseDialog()}>
	<Dialog.Content class="{currentSize} overflow-hidden">
		<div><slot /></div>
	</Dialog.Content>
</Dialog.Root>
