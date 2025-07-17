<!-- src/lib/components/dashboard/DashboardSidebar.svelte -->
<script lang="ts">
	import { page } from '$app/state';
	import AppLogo from '$lib/components/common/AppLogo.svelte';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import { LayoutDashboard, User, FolderOpen, Settings } from '@lucide/svelte';

	const navItems = [
		{ title: 'Dashboard', url: '/dashboard', icon: LayoutDashboard },
		{ title: 'Profile', url: '/dashboard/profile', icon: User },
		{ title: 'Assets', url: '/dashboard/assets', icon: FolderOpen },
		{ title: 'Settings', url: '/dashboard/options', icon: Settings }
	];

	function isActiveRoute(itemUrl: string, currentPath: string): boolean {
		if (itemUrl === '/dashboard') {
			return currentPath === '/dashboard';
		}
		return currentPath.startsWith(itemUrl);
	}
</script>

<Sidebar.Root>
	<Sidebar.Content class="bg-background border-border border-r ">
		<Sidebar.Group class="h-full p-0">
			<!-- Logo Section -->
			<div class="border-border flex h-16 items-center border-b px-4">
				<AppLogo />
			</div>

			<!-- Navigation Menu -->
			<Sidebar.GroupContent class="mt-4 p-2">
				<Sidebar.Menu class="h-full space-y-1">
					{#each navItems as item (item.title)}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton>
								<a
									href={item.url}
									class="hover:bg-accent hover:text-accent-foreground flex w-full items-center gap-3 rounded-lg px-3 py-2.5 text-sm transition-all duration-200"
									class:bg-primary={isActiveRoute(item.url, page.url.pathname)}
									class:text-primary-foreground={isActiveRoute(item.url, page.url.pathname)}
									class:shadow-sm={isActiveRoute(item.url, page.url.pathname)}
									class:font-medium={isActiveRoute(item.url, page.url.pathname)}
								>
									<svelte:component this={item.icon} class="h-4 w-4 flex-shrink-0" />
									<span class="truncate">{item.title}</span>
								</a>
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>

			<!-- Bottom Section -->
			<div class="border-border mt-auto border-t p-4">
				<div class="text-muted-foreground text-center text-xs">
					<p>© {new Date().getFullYear()}</p>
					Asset Management System
				</div>
			</div>
		</Sidebar.Group>
	</Sidebar.Content>
</Sidebar.Root>
