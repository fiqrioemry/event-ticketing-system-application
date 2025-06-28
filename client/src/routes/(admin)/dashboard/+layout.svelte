<script lang="ts">
	import UserMenu from '$lib/components/common/UserMenu.svelte';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import AppSidebar from '$lib/components/layout/AppSidebar.svelte';

	export let data: any = null;
	let user = data?.user?.user_metadata || null;
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
			<UserMenu {user} />
		</header>
		<main class="bg-muted">
			<div class="mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-7xl flex-1 flex-col py-12">
				<slot />
			</div>
		</main>
	</Sidebar.Inset>
</Sidebar.Provider>
