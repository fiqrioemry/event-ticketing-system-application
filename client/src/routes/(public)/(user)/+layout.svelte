<script lang="ts">
	import { page } from '$app/stores';
	import { User, ShoppingBag, Undo2, LogOut } from '@lucide/svelte';

	const user = {
		name: 'Ahmad Fiqri',
		email: 'ahmad@example.com',
		avatar: 'https://placehold.co/100x100?text=AF',
		balance: 150000,
		role: 'user'
	};

	$: currentPath = $page.url.pathname;
</script>

<svelte:head>
	<title>Profile | MyBrand</title>
	<meta name="description" content="Manage your profile and account settings with MyBrand." />
</svelte:head>

<section class="mx-auto my-8 flex max-w-7xl gap-6 px-4">
	<aside class="bg-background h-96 w-60 rounded-lg border p-4">
		<div class="mb-6 flex flex-col items-center text-center">
			<img src={user.avatar} alt="Avatar" class="h-20 w-20 rounded-full object-cover" />
			<h2 class="mt-2 text-lg font-semibold">{user.name}</h2>
			<p class="text-muted-foreground text-sm">{user.email}</p>
			<p class="mt-1 text-xs text-green-600 dark:text-green-400">
				Saldo: Rp {user.balance.toLocaleString('id-ID')}
			</p>
		</div>

		<nav class="flex flex-col gap-1 text-sm font-medium">
			<a
				href="/profile"
				class="flex items-center gap-2 rounded-md px-3 py-2 transition-colors
					{currentPath === '/profile' ? 'bg-muted text-primary' : 'hover:bg-muted hover:text-primary'}"
			>
				<User class="h-4 w-4" />
				Profile
			</a>
			<a
				href="/purchase"
				class="flex items-center gap-2 rounded-md px-3 py-2 transition-colors
					{currentPath === '/purchase' ? 'bg-muted text-primary' : 'hover:bg-muted hover:text-primary'}"
			>
				<ShoppingBag class="h-4 w-4" />
				Purchasement
			</a>
			<a
				href="/refund"
				class="flex items-center gap-2 rounded-md px-3 py-2 transition-colors
					{currentPath === '/refund' ? 'bg-muted text-primary' : 'hover:bg-muted hover:text-primary'}"
			>
				<Undo2 class="h-4 w-4" />
				Refund
			</a>
			<a
				href="/logout"
				class="mt-2 flex items-center gap-2 rounded-md px-3 py-2 text-red-500 transition-colors hover:bg-red-50 hover:underline dark:hover:bg-red-900/10"
			>
				<LogOut class="h-4 w-4" />
				Logout
			</a>
		</nav>
	</aside>

	<main class="bg-background min-h-[calc(90vh-4rem)] flex-1">
		<slot />
	</main>
</section>
