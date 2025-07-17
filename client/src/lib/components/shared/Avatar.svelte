<script lang="ts">
	export let src: string | null = null;
	export let name: string = '';
	export let size: number = 40;
	export let className: string = '';

	let hasError = false;

	function getInitials(name: string): string {
		return name
			.split(' ')
			.map((word) => word[0]?.toUpperCase())
			.slice(0, 2)
			.join('');
	}
</script>

<div
	class={`bg-muted text-foreground inline-flex  items-center justify-center overflow-hidden rounded-full border border-4 border-white font-medium ${className}`}
	style={`width: ${size}px; height: ${size}px; font-size: ${size * 0.4}px`}
>
	{#if src && !hasError}
		<img
			{src}
			alt={name}
			loading="lazy"
			class="h-full w-full object-cover"
			on:error={() => (hasError = true)}
		/>
	{:else}
		<span>{getInitials(name)}</span>
	{/if}
</div>
