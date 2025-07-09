<script lang="ts">
	import { page } from '$app/state';
	import Logo from '$lib/assets/TIKETKU.svg';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import { Cloud, Brain, CreditCard, UserRound, Settings } from '@lucide/svelte';

	const navItems = [
		{ title: 'Summary', url: '/dashboard/summary', icon: Cloud },
		{ title: 'events', url: '/dashboard/events', icon: Brain },
		{ title: 'payment', url: '/dashboard/payment', icon: CreditCard },
		{ title: 'ticket', url: '/dashboard/ticket', icon: UserRound },
		{ title: 'user', url: '/dashboard/user', icon: Settings }
	];
</script>

<Sidebar.Root>
	<Sidebar.Content class="bg-background">
		<Sidebar.Group class="p-0">
			<div class="flex h-16 items-center border-b px-4">
				<div>
					<img class="h-8" src={Logo} alt="TIKETKU" />
				</div>
			</div>

			<Sidebar.GroupContent class="mt-4 p-2">
				<Sidebar.Menu class="space-y-1">
					{#each navItems as item (item.title)}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton>
								<a
									href={item.url}
									class="hover:bg-muted flex w-full items-center gap-2 rounded-md px-3 py-2 text-sm transition"
									class:bg-muted={page.url.pathname.startsWith(item.url)}
									class:text-primary={page.url.pathname.startsWith(item.url)}
									class:font-semibold={page.url.pathname.startsWith(item.url)}
								>
									<item.icon class="h-4 w-4" />
									<span>{item.title}</span>
								</a>
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>
		</Sidebar.Group>
	</Sidebar.Content>
</Sidebar.Root>
