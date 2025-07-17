<script lang="ts">
	import { Loader } from '@lucide/svelte';
	import Button from '$lib/components/ui/button/button.svelte';

	export let className: string = 'w-full';
	export let isLoading: boolean = false;
	export let disabled: boolean = false;
	export let buttonText: string = 'Sign In';
	export let onsubmit: () => void = () => {};
	export let buttonLoadingText: string = 'Signing in...';
	export let variant: 'primary' | 'secondary' | 'ghost' | 'link' | 'default' = 'primary';
</script>

<Button type="submit" {variant} onclick={onsubmit} {disabled} class={className}>
	{#if isLoading}
		<Loader class="h-4 w-4 animate-spin" />
		<span>{buttonLoadingText}</span>
	{:else}
		<slot />
		<span>{buttonText}</span>
	{/if}
</Button>
