<!-- NoResult.svelte -->
<script lang="ts">
	export let title = 'Data tidak ditemukan';
	export let description = 'Silakan ubah pencarian atau filter Anda.';
</script>

<div
	class="text-muted-foreground flex min-h-[calc(70vh-200px)] flex-col items-center justify-center py-12 text-center"
>
	<p class="text-lg font-medium">{title}</p>
	<p class="mt-1 text-sm">{description}</p>
</div>
