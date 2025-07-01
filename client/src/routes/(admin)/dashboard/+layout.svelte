<script lang="ts">
	import AdminMenu from '$lib/components/common/AdminMenu.svelte';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import AppSidebar from '$lib/components/layout/AppSidebar.svelte';
	import { userSession } from '$lib/stores/auth.store.js';

	$: user = $userSession;
</script>

<Sidebar.Provider>
	<AppSidebar />
	<Sidebar.Inset>
		<header
			class="bg-background md: sticky top-0 z-50 flex h-16 items-center justify-between border-b px-2 md:justify-end"
		>
			<div class="block md:hidden">
				<Sidebar.Trigger />
			</div>
			<AdminMenu {user} />
		</header>
		<main class="bg-muted">
			<div class="mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-7xl flex-1 flex-col py-12">
				<slot />
			</div>
		</main>
	</Sidebar.Inset>
</Sidebar.Provider>
