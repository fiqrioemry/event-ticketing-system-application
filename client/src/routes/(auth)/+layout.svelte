<!-- src/routes/(auth)/+layout.svelte -->
<script lang="ts">
	import { goto } from '$app/navigation';
	const { children } = $props<{ children: any }>();
</script>

<section class="flex h-screen overflow-hidden">
	{@render children?.()}
</section>
