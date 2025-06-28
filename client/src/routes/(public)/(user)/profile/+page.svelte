<script lang="ts">
	import Label from '$lib/components/ui/label/label.svelte';
	import Button from '$lib/components/ui/button/button.svelte';

	const userProfile = {
		id: 'bdb598a3-1c86-4e95-93d2-65cda21b4b33',
		email: 'charlie@event.com',
		fullname: 'Charlie Customer',
		avatar: 'https://api.dicebear.com/6.x/initials/svg?seed=charlie',
		role: 'user',
		balance: '210000.00',
		joinedAt: '2025-06-28'
	};
</script>

<svelte:head>
	<title>Profile | MyBrand</title>
	<meta name="description" content="Manage your profile and account settings with MyBrand." />
</svelte:head>

<section class="bg-background mx-auto max-w-3xl rounded-lg border p-6">
	<div class="mb-6 flex items-center gap-4">
		<img src={userProfile.avatar} alt="Avatar" class="h-16 w-16 rounded-full border" />
		<div>
			<p class="text-lg font-semibold">{userProfile.fullname}</p>
			<p class="text-muted-foreground text-sm capitalize">{userProfile.email}</p>
		</div>
	</div>

	<div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
		<div>
			<Label class="text-sm font-medium">Full Name</Label>
			<p class="bg-muted rounded-md px-3 py-2 text-sm">{userProfile.fullname}</p>
		</div>
		<div>
			<Label class="text-sm font-medium">Email</Label>
			<p class="bg-muted rounded-md px-3 py-2 text-sm">{userProfile.email}</p>
		</div>
		<div class="sm:col-span-2">
			<Label class="text-sm font-medium">Joined At</Label>
			<p class="rounded-md px-3 py-2 text-sm">{userProfile.joinedAt}</p>
		</div>
	</div>

	<Button class="mt-6" variant="default" size="lg">Update Profile</Button>
</section>
