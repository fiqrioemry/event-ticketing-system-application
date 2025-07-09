<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';

	let email: string = '';
	let fullname: string = '';
	let password: string = '';

	function inputValidation() {
		return email.length > 0 && fullname.length > 0 && password.length > 0;
	}

	async function handleRegister() {
		if (!(await inputValidation())) {
			return;
		}

		const form = { email, fullname, password };
		try {
			// Simulate a registration request
			console.log('Registering with', form);
			// Here you would typically make an API call to register the user
			alert('Registration successful!');
		} catch (error) {
			console.error('Registration failed:', error);
			alert('Registration failed. Please try again.');
		}
	}
</script>

<svelte:head>
	<title>Signup | Tiketku</title>
	<meta name="description" content="Signup page for Tiketku." />
</svelte:head>

<section class="flex h-screen items-center justify-center bg-gray-100">
	<div class="bg-background w-full max-w-md rounded-lg p-4 shadow-md">
		<div class="mb-6 text-center">
			<h1 class="text-2xl font-bold">Create Your Account</h1>
			<p class="text-muted-foreground text-sm">
				Start your journey with Tiketku by creating an account.
			</p>
		</div>
		<form on:submit|preventDefault={handleRegister} class="space-y-4">
			<div class="space-y-2">
				<Label for="fullname">Fullname</Label>
				<Input
					type="text"
					id="fullname"
					bind:value={fullname}
					placeholder="Enter your full name"
					required
				/>
			</div>

			<div class="space-y-2">
				<Label for="email">Email</Label>
				<Input type="email" id="email" bind:value={email} placeholder="Enter your email" required />
			</div>

			<div class="space-y-2">
				<Label for="password">Password</Label>
				<Input
					type="password"
					id="password"
					bind:value={password}
					placeholder="Enter your password"
					required
				/>
			</div>

			<Button class="bg-primary w-full" disabled={!inputValidation()} type="submit">Signup</Button>
		</form>

		<div class="text-muted-foreground mt-4 text-center text-sm">
			<p>Already have an account? <a href="/signin" class="text-primary">Signin</a></p>
		</div>
	</div>
</section>
