<script lang="ts">
	import { toast } from 'svelte-sonner';
	import { PlusCircle, X } from '@lucide/svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import Button from '$lib/components/ui/button/button.svelte';

	export let field: string;
	export let label = '';
	export let maxImages: number = 5;
	export let maxSizeMB: number = 2;
	export let isSingle: boolean = false;
	export let form: Record<string, any>;
	export let errors: Record<string, string>;
	export let touched: Record<string, boolean>;

	let isDragging = false;

	function handleFiles(files: FileList) {
		const validFiles = Array.from(files).filter((file) => {
			const isValidSize = file.size / 1024 / 1024 <= maxSizeMB;
			if (!isValidSize) {
				toast.error(`${file.name} exceeds ${maxSizeMB}MB`);
			}
			return isValidSize;
		});
		if (validFiles.length === 0) return;

		if (isSingle) {
			form[field] = validFiles[0];
		} else {
			const current = form[field] || [];
			const updated = [...current, ...validFiles].slice(0, maxImages);
			form[field] = updated;
		}
	}

	function removeImage(img: File | string) {
		if (isSingle) {
			form[field] = null;
		} else {
			form[field] = (form[field] || []).filter((f: any) => f !== img);
		}
	}

	function getImageURL(img: File | string): string {
		return img instanceof File ? URL.createObjectURL(img) : img;
	}

	function onDrop(event: DragEvent) {
		event.preventDefault();
		isDragging = false;
		if (event.dataTransfer?.files.length) {
			handleFiles(event.dataTransfer.files);
			event.dataTransfer.clearData();
		}
	}
</script>

<div class="space-y-2">
	{#if label}
		<Label for={field}>{label}</Label>
	{/if}

	<div
		role="region"
		on:drop={onDrop}
		on:dragover|preventDefault={() => (isDragging = true)}
		on:dragleave|preventDefault={() => (isDragging = false)}
		class={`rounded-md border-2 transition ${
			isDragging ? 'border-primary bg-primary/10' : 'border-border bg-muted/30'
		} ${isSingle ? 'relative flex h-64 w-full items-center justify-center overflow-hidden' : 'flex flex-wrap gap-4 p-4'}`}
	>
		{#if isSingle}
			{#if form[field]}
				<div class="relative h-full w-full">
					<img
						src={getImageURL(form[field])}
						class="h-full w-full rounded-md object-cover"
						alt="preview"
					/>
					<Button
						type="button"
						variant="destructive"
						class="absolute top-1 right-1 p-1"
						onclick={() => removeImage(form[field])}
					>
						<X class="h-4 w-4" />
					</Button>
				</div>
			{:else}
				<label
					for={`${field}-upload`}
					class="hover:bg-muted flex h-full w-full cursor-pointer flex-col items-center justify-center rounded-md transition"
				>
					<PlusCircle class="text-primary mb-2" />
					<span class="text-sm">Select Image</span>
				</label>
			{/if}
		{:else}
			{#each form[field] || [] as img, idx (idx)}
				<div class="relative h-32 w-32 overflow-hidden rounded-md border">
					<img src={getImageURL(img)} class="h-full w-full object-cover" alt="preview" />
					<Button
						size="icon"
						type="button"
						variant="destructive"
						class="absolute top-1 right-1 p-1"
						onclick={() => removeImage(img)}
					>
						<X class="h-2 w-2" />
					</Button>
				</div>
			{/each}
			{#if !form[field] || form[field].length < maxImages}
				<label
					for={`${field}-upload`}
					class="border-primary hover:bg-muted flex h-32 w-32 cursor-pointer flex-col items-center justify-center rounded-md border-2 border-dashed transition"
				>
					<PlusCircle class="text-primary mb-2" />
					<span class="text-sm">Select Images</span>
				</label>
			{/if}
		{/if}
		<input
			id={`${field}-upload`}
			type="file"
			accept="image/*"
			hidden
			multiple={!isSingle}
			on:change={(e) => {
				handleFiles((e.target as HTMLInputElement).files!);
				(e.target as HTMLInputElement).value = '';
			}}
		/>
	</div>

	{#if touched[field] && errors[field]}
		<p class="text-destructive mt-1 text-xs">{errors[field]}</p>
	{/if}
</div>
