<!-- src/lib/components/ui/GlobalConfirmDialog.svelte -->
<script lang="ts">
	import { Trash2 } from '@lucide/svelte';
	import { confirmStore } from '$lib/stores/confirm.store';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Dialog from '$lib/components/ui/dialog/index.js';

	$: state = $confirmStore;

	let typedValue = '';

	$: canConfirm =
		state.type === 'simple' || (state.type === 'typed' && typedValue === state.confirmWord);

	// Reset typed value when dialog opens
	$: if (state.isOpen && state.type === 'typed') {
		typedValue = '';
	}

	function handleOpenChange(open: boolean) {
		if (!open && state.isOpen) {
			confirmStore.cancel();
		}
	}

	function handleConfirm() {
		if (canConfirm) {
			confirmStore.confirm();
		}
	}
</script>

<Dialog.Root open={state.isOpen} onOpenChange={handleOpenChange}>
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title class="flex items-center gap-3">
				<div class="flex-shrink-0 rounded-full border border-red-200 bg-red-50 p-2">
					<Trash2 class="h-5 w-5 text-red-600" />
				</div>
				<span class="text-red-800">{state.title}</span>
			</Dialog.Title>
			<Dialog.Description class="text-red-600">
				{state.message}
			</Dialog.Description>
		</Dialog.Header>

		<!-- Typed confirmation input -->
		{#if state.type === 'typed' && state.confirmWord}
			<div class="py-4">
				<p class="mb-3 text-sm text-gray-700">
					Type <code class="rounded bg-gray-100 px-2 py-1 font-mono text-red-600"
						>{state.confirmWord}</code
					> to confirm:
				</p>
				<Input bind:value={typedValue} placeholder="Type here..." class="w-full" autofocus />
			</div>
		{/if}

		<Dialog.Footer class="gap-3">
			<Button variant="outline" onclick={confirmStore.cancel}>Cancel</Button>
			<Button variant="destructive" onclick={handleConfirm} disabled={!canConfirm}>Delete</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
