<script lang="ts">
	import { goto } from '$app/navigation';
	import { useLogout } from '$lib/hooks/useAuth';
	import { DollarSign, LogOut, UserCircle } from '@lucide/svelte';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
	import { Avatar, AvatarFallback, AvatarImage } from '$lib/components/ui/avatar';

	export let user;
</script>

<DropdownMenu.Root>
	<DropdownMenu.Trigger class="hover:bg-muted flex items-center gap-4 rounded-xl p-2 duration-300">
		<Avatar class="h-9 w-9">
			<AvatarImage src={user?.avatar} alt={user?.fullname} />
			<AvatarFallback>{user?.fullname?.[0] || 'A'}</AvatarFallback>
		</Avatar>
		<div class="hidden flex-col overflow-hidden text-left md:flex">
			<span class=" truncate text-xs font-bold">
				{user?.fullname || 'admin'}
			</span>
			<span class="truncate text-xs">
				{user?.email || 'admin@example.com'}
			</span>
		</div>
	</DropdownMenu.Trigger>

	<DropdownMenu.Content align="end" class="w-52">
		<DropdownMenu.Item onclick={() => goto('/dashboard')}>
			<UserCircle class="mr-2 h-4 w-4" /> overview
		</DropdownMenu.Item>
		<DropdownMenu.Item onclick={() => goto('/dashboard/events')}>
			<DollarSign class="mr-2 h-4 w-4" /> events
		</DropdownMenu.Item>
		<DropdownMenu.Item onclick={() => goto('/dashboard/payment')}>
			<DollarSign class="mr-2 h-4 w-4" /> payment
		</DropdownMenu.Item>
		<DropdownMenu.Item class="text-red-600" onclick={useLogout}>
			<LogOut href="/signout" class="mr-2 h-4 w-4" /> Logout
		</DropdownMenu.Item>
	</DropdownMenu.Content>
</DropdownMenu.Root>
