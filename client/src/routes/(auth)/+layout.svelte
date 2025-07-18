<script lang="ts">
	import { goto } from '$app/navigation';
	import { isAuthenticated } from '$lib/stores/auth.store';
	if ($isAuthenticated) {
		goto('/user/profile');
	}
</script>

<section class="flex h-screen overflow-hidden">
	<slot />
</section>
