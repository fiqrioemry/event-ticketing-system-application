<script lang="ts">
	import Label from '$lib/components/ui/label/label.svelte';
	import Button from '$lib/components/ui/button/button.svelte';

	const user = {
		name: 'Ahmad Fiqri',
		email: 'ahmad@example.com',
		avatar: 'https://placehold.co/100x100?text=AF',
		balance: 150000,
		role: 'user'
	};
</script>

<svelte:head>
	<title>Profile | MyBrand</title>
	<meta name="description" content="Manage your profile and account settings with MyBrand." />
</svelte:head>

<section class="bg-background rounded-lg border p-4">
	<h1 class="mb-4 text-2xl font-bold">👤 Profile Dashboard</h1>

	<div class="space-y-4">
		<div>
			<Label class="text-sm font-medium">Full Name</Label>
			<p class="bg-muted rounded-md px-3 py-2 text-sm">{user.name}</p>
		</div>
		<div>
			<Label class="text-sm font-medium">Email</Label>
			<p class="bg-muted rounded-md px-3 py-2 text-sm">{user.email}</p>
		</div>
		<Button class="mt-6" variant="default" size="lg">Update Profile</Button>
	</div>
</section>
