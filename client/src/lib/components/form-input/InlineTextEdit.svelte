<!-- InlineEdit.svelte -->
<script lang="ts">
	import { Check, Edit, Loader2, X } from '@lucide/svelte';

	export let name: string = '';
	export let value: string = '';
	export let label: string = '';
	export let loading = false;
	export let minlength: number = 3;
	export let maxlength: number = 100;
	export let onSave: (value: string) => Promise<boolean | void> = () => Promise.resolve();

	let editing = false;
	let tempValue = value;
	let previousValue = value;

	function startEdit() {
		editing = true;
		tempValue = value;
	}

	function cancel() {
		editing = false;
		tempValue = value;
	}

	async function save() {
		tempValue = tempValue.trim();
		if (tempValue.length === 0 || tempValue === previousValue) {
			return cancel();
		}

		await onSave(tempValue);
		editing = false;
	}
</script>

<div>
	<label for={label} class="mb-1 block text-sm font-medium">{label}</label>

	{#if editing}
		<div class="flex gap-2">
			<input
				{name}
				{minlength}
				{maxlength}
				bind:value={tempValue}
				class="flex-1 rounded-md border border-gray-300 px-3 py-2"
				on:keydown={(e) => {
					if (e.key === 'Enter') save();
					if (e.key === 'Escape') cancel();
				}}
			/>
			<button class="btn-action-save" on:click={save} disabled={loading}>
				{#if loading}<Loader2 class="h-4 w-4 animate-spin" />{:else}<Check class="h-4 w-4" />{/if}
			</button>
			<button class="btn-action-cancel" on:click={cancel}>
				<X class="h-4 w-4" />
			</button>
		</div>
	{:else}
		<div class="bg-muted flex items-center justify-between rounded-md px-3 py-2">
			<span>{value}</span>
			<button class="btn-edit" on:click={startEdit}>
				<Edit class="h-4 w-4" />
			</button>
		</div>
	{/if}
</div>
