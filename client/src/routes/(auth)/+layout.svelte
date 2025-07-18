<script lang="ts">
	import { goto } from '$app/navigation';
	import { isAuthenticated } from '$lib/stores/auth.store';

	if ($isAuthenticated) {
		goto('/user/profile');
	}

	export let children;
</script>

<section class="flex h-screen overflow-hidden">
	{@render children?.()}
</section>
