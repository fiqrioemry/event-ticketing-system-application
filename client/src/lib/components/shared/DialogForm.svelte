<!-- src/lib/components/dialogs/FormDialog.svelte -->
<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';

	export let open = false;
	export let title = '';
	export let description = '';
	export let size: 'sm' | 'md' | 'lg' = 'md';
	export let onClose: (() => void) | undefined = undefined;

	const sizeClasses = {
		sm: 'max-w-sm',
		md: 'max-w-md',
		lg: 'max-w-lg'
	};

	function handleOpenChange(isOpen: boolean) {
		open = isOpen;
		if (!isOpen && onClose) {
			onClose();
		}
	}
</script>

<Dialog.Root bind:open onOpenChange={handleOpenChange}>
	<Dialog.Content class={sizeClasses[size]}>
		<Dialog.Header>
			<Dialog.Title>{title}</Dialog.Title>
			{#if description}
				<Dialog.Description>{description}</Dialog.Description>
			{/if}
		</Dialog.Header>

		<div class="py-4">
			<slot />
		</div>
	</Dialog.Content>
</Dialog.Root>
