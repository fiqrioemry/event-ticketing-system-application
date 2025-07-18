<script lang="ts">
	import { page } from '$app/state';
	import { goto } from '$app/navigation';
	import { formatPrice } from '$lib/utils/formatter';
	import * as Sheet from '$lib/components/ui/sheet';
	import Avatar from '$lib/components/shared/Avatar.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { isAuthenticated, authUser, authStore } from '$lib/stores/auth.store';
	import { UserCircle, Package, Bell, LogOut, Menu } from '@lucide/svelte';

	if (!$isAuthenticated) {
		goto('/signin');
	}

	const navItems = [
		{
			label: 'Profile',
			href: '/user/profile',
			icon: UserCircle
		},
		{
			label: 'Orders',
			href: '/user/orders',
			icon: Package
		}
	];

	let isSheetOpen = $state(false);
	let currentPath = $state('');

	// Update currentPath saat component mount dan saat page berubah
	$effect(() => {
		currentPath = page.url.pathname;
	});

	function closeSheet() {
		isSheetOpen = false;
	}

	function openSheet() {
		isSheetOpen = true;
	}

	function handleLogout() {
		authStore.logout().then(() => {
			goto('/signin');
		});
	}
</script>

<section class="container mx-auto max-w-7xl px-4 py-6">
	<!-- Mobile Header -->
	<div class="mb-6 flex items-center justify-between rounded-xl bg-white p-4 shadow-sm lg:hidden">
		<div class="flex items-center gap-3">
			<Avatar size={40} src={$authUser?.avatar} name={$authUser?.fullname} />
			<div>
				<h2 class="text-sm font-semibold text-gray-900">{$authUser?.fullname}</h2>
				<p class="text-xs text-gray-500">{formatPrice($authUser?.balance ?? 0)}</p>
			</div>
		</div>

		<Sheet.Root bind:open={isSheetOpen}>
			<Sheet.Trigger>
				<button onclick={openSheet} class="rounded-lg p-2 transition-colors hover:bg-gray-100">
					<Menu class="h-5 w-5 text-gray-600" />
				</button>
			</Sheet.Trigger>

			<Sheet.Content side="left" class="w-80">
				<Sheet.Header class="space-y-4 border-b pb-6">
					<div class="flex items-center gap-4">
						<Avatar size={56} src={$authUser?.avatar} name={$authUser?.fullname} />
						<div class="text-left">
							<Sheet.Title class="text-base font-semibold">{$authUser?.fullname}</Sheet.Title>
							<Sheet.Description class="text-sm">{$authUser?.email}</Sheet.Description>
						</div>
					</div>
					<div class="rounded-lg bg-blue-50 p-3">
						<div class="flex items-center justify-between">
							<span class="text-sm font-medium text-blue-900">Balance</span>
							<span class="text-sm font-semibold text-blue-600">
								{formatPrice($authUser?.balance ?? 0)}
							</span>
						</div>
					</div>
				</Sheet.Header>

				<div class="space-y-2 p-2">
					{#each navItems as item}
						<Button
							href={item.href}
							onclick={closeSheet}
							variant="ghost"
							class="flex h-10 items-center justify-start gap-3 rounded-lg px-3 py-3 transition-colors {currentPath ===
							item.href
								? 'bg-blue-50 text-blue-600'
								: 'text-gray-700 hover:bg-gray-50'}"
						>
							<item.icon />
							<span class="font-medium">{item.label}</span>
							{#if currentPath === item.href}
								<div class="ml-auto h-2 w-2 rounded-full bg-blue-500"></div>
							{/if}
						</Button>
					{/each}

					<div class="mt-4 border-t pt-4">
						<Button variant="event-outline" class="h-12 w-full" onclick={handleLogout}>
							<LogOut class="h-5 w-5" />
							<span class="font-medium">Logout</span>
						</Button>
					</div>
				</div>
			</Sheet.Content>
		</Sheet.Root>
	</div>

	<div class="flex gap-6">
		<!-- Desktop Sidebar -->
		<aside class="hidden w-72 lg:block">
			<div class="sticky top-6 rounded-xl border border-gray-100 bg-white p-6 shadow-sm">
				<!-- Desktop Sidebar Header -->
				<div class="mb-6 border-b border-gray-100 pb-6">
					<div class="text-center">
						<Avatar size={80} src={$authUser?.avatar} name={$authUser?.fullname} />
						<h3 class="mt-4 text-lg font-semibold text-gray-900">{$authUser?.fullname}</h3>
						<p class="text-sm text-gray-500">{$authUser?.email}</p>

						<!-- Balance Card -->
						<div class="mt-4 rounded-xl bg-gradient-to-r from-blue-50 to-indigo-50 p-4">
							<div class="text-xs font-medium tracking-wide text-gray-600 uppercase">
								Account Balance
							</div>
							<div class="text-xl font-bold text-blue-600">
								{formatPrice($authUser?.balance ?? 0)}
							</div>
						</div>
					</div>
				</div>

				<!-- Desktop Navigation -->
				<nav class="space-y-3">
					{#each navItems as item}
						<Button
							variant="ghost"
							href={item.href}
							class="h-12 w-full justify-start hover:bg-gray-100 {currentPath === item.href
								? 'bg-gray-100 '
								: 'bg-none'}"
							><item.icon /><span class="font-medium">{item.label}</span>
						</Button>
					{/each}

					<div class="mt-4 border-t pt-4">
						<Button
							variant="event-ghost"
							class="h-12 w-full bg-red-100 text-red-500 hover:bg-none hover:text-red-600"
							onclick={handleLogout}
						>
							<LogOut class="h-5 w-5" />
							<span class="font-medium">Logout</span>
						</Button>
					</div>
				</nav>
			</div>
		</aside>

		<!-- Main Content -->
		<main class="min-h-[600px] flex-1">
			<slot />
		</main>
	</div>
</section>
