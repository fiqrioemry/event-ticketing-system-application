<script lang="ts">
	import { page } from '$app/state';
	import Avatar from './Avatar.svelte';
	import { goto } from '$app/navigation';
	import { authStore } from '$lib/stores/auth.store';
	import { LogOut, User, Settings } from '@lucide/svelte';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';

	export let user;
</script>

<DropdownMenu.Root>
	<DropdownMenu.Trigger>
		<Avatar src={user.avatar} name={user.fullname} />
	</DropdownMenu.Trigger>
	<DropdownMenu.Content class="w-56" align="end" sideOffset={8}>
		<DropdownMenu.Label>
			<span class="font-medium">{user.fullname}</span>
			<span class="text-muted-foreground text-sm">{user.email}</span>
		</DropdownMenu.Label>
		<DropdownMenu.Separator />
		<DropdownMenu.Item onclick={() => goto('/user/profile')}>
			<User class="h-4 w-4" />
			Profile
		</DropdownMenu.Item>
		<DropdownMenu.Item
			onclick={() => goto('/user/orders')}
			class={page.url.pathname === '/user/orders' ? 'bg-blue-500' : ''}
		>
			<Settings class="h-4 w-4" />
			Orders
		</DropdownMenu.Item>
		<DropdownMenu.Separator />
		<DropdownMenu.Item onclick={() => authStore.logout()}>
			<LogOut class="h-4 w-4" />
			Logout
		</DropdownMenu.Item>
	</DropdownMenu.Content>
</DropdownMenu.Root>
