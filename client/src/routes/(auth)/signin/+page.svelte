<script lang="ts">
	import { useLogin } from '$lib/hooks/useAuth';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { loginForm, isLoading, error } from '$lib/stores/auth.store';

	$: isFormValid = $loginForm.email && $loginForm.password;
</script>

<svelte:head>
	<title>Login | Tiketku</title>
	<meta name="description" content="Login page for Tiketku." />
</svelte:head>

<section class="flex h-screen overflow-hidden">
	<!-- Left Side - Welcome Content -->
	<div
		class="from-primary/5 via-primary/10 to-primary/5 relative hidden overflow-hidden bg-gradient-to-br lg:flex lg:w-1/2"
	>
		<!-- Background Pattern -->
		<div class="absolute inset-0 opacity-5">
			<div class="bg-primary absolute -top-40 -left-40 h-96 w-96 rounded-full blur-3xl"></div>
			<div class="bg-primary absolute -right-40 -bottom-40 h-96 w-96 rounded-full blur-3xl"></div>
		</div>

		<!-- Content -->
		<div class="relative z-10 flex flex-col justify-center px-12 xl:px-16">
			<div class="mb-8">
				<div
					class="bg-primary/10 mb-8 inline-flex h-16 w-16 items-center justify-center rounded-2xl"
				>
					<svg class="text-primary h-8 w-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M15 5v2m0 4v2m0 4v2M5 5a2 2 0 00-2 2v3a2 2 0 110 4v3a2 2 0 002 2h14a2 2 0 002-2v-3a2 2 0 110-4V7a2 2 0 00-2-2H5z"
						></path>
					</svg>
				</div>

				<h1 class="text-foreground mb-6 text-4xl leading-tight font-bold xl:text-5xl">
					Welcome to <span class="text-primary">Tiketku</span>
				</h1>

				<p class="text-muted-foreground mb-8 text-lg leading-relaxed xl:text-xl">
					Your ultimate destination for seamless ticket booking and event management. Experience the
					convenience of booking tickets for concerts, movies, sports, and more.
				</p>
			</div>

			<!-- Features -->
			<div class="space-y-4">
				<div class="flex items-center space-x-4">
					<div class="bg-primary h-2 w-2 rounded-full"></div>
					<span class="text-muted-foreground">Secure and fast booking process</span>
				</div>
				<div class="flex items-center space-x-4">
					<div class="bg-primary h-2 w-2 rounded-full"></div>
					<span class="text-muted-foreground">Thousands of events available</span>
				</div>
				<div class="flex items-center space-x-4">
					<div class="bg-primary h-2 w-2 rounded-full"></div>
					<span class="text-muted-foreground">24/7 customer support</span>
				</div>
			</div>
		</div>
	</div>

	<!-- Right Side - Login Form -->
	<div class="bg-background relative flex w-full items-center justify-center lg:w-1/2">
		<!-- Header Navigation -->
		<div class="absolute top-8 right-8 left-8 flex items-center justify-between">
			<a
				href="/"
				class="group inline-flex items-center space-x-2 transition-all duration-200 hover:scale-105"
			>
				<div
					class="bg-primary/10 group-hover:bg-primary/20 inline-flex h-10 w-10 items-center justify-center rounded-xl transition-colors"
				>
					<svg class="text-primary h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M15 5v2m0 4v2m0 4v2M5 5a2 2 0 00-2 2v3a2 2 0 110 4v3a2 2 0 002 2h14a2 2 0 002-2v-3a2 2 0 110-4V7a2 2 0 00-2-2H5z"
						></path>
					</svg>
				</div>
				<span
					class="text-foreground group-hover:text-primary hidden font-semibold transition-colors lg:block"
					>Tiketku</span
				>
			</a>

			<!-- Back to Home Button (Alternative Style) -->
			<a
				href="/"
				class="text-muted-foreground hover:text-foreground group inline-flex items-center space-x-2 px-4 py-2 text-sm transition-colors"
			>
				<svg
					class="h-4 w-4 transition-transform group-hover:-translate-x-0.5"
					fill="none"
					stroke="currentColor"
					viewBox="0 0 24 24"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M10 19l-7-7m0 0l7-7m-7 7h18"
					></path>
				</svg>
				<span>Back to Home</span>
			</a>
		</div>

		<div class="w-full max-w-md px-8 lg:px-0">
			<!-- Header -->
			<div class="mb-8 text-center lg:text-left">
				<h2 class="text-foreground mb-2 text-3xl font-bold">Welcome Back</h2>
				<p class="text-muted-foreground">Sign in to your account to continue your journey</p>
			</div>

			<!-- Form -->
			<form on:submit|preventDefault={useLogin} class="space-y-6">
				{#if $error}
					<div class="rounded-md border border-red-200 bg-red-50 px-4 py-3 text-red-700">
						{$error}
					</div>
				{/if}

				<div class="space-y-2">
					<Label for="email" class="text-foreground text-sm font-medium">Email Address</Label>
					<Input
						type="email"
						id="email"
						bind:value={$loginForm.email}
						placeholder="Enter your email address"
						class="border-input focus:border-primary focus:ring-primary/20 h-12 border-2 px-4 text-base transition-all duration-200 focus:ring-2"
						disabled={$isLoading}
					/>
				</div>

				<div class="space-y-2">
					<div class="flex items-center justify-between">
						<Label for="password" class="text-foreground text-sm font-medium">Password</Label>
						<a
							href="/forgot-password"
							class="text-primary hover:text-primary/80 text-sm transition-colors"
						>
							Forgot password?
						</a>
					</div>
					<Input
						type="password"
						id="password"
						bind:value={$loginForm.password}
						placeholder="Enter your password"
						class="border-input focus:border-primary focus:ring-primary/20 h-12 border-2 px-4 text-base transition-all duration-200 focus:ring-2"
						disabled={$isLoading}
					/>
				</div>

				<!-- Remember Me -->
				<div class="flex items-center space-x-2">
					<input
						bind:checked={$loginForm.rememberMe}
						type="checkbox"
						id="remember"
						class="text-primary border-input focus:ring-primary h-4 w-4 rounded border-2 focus:ring-2"
						disabled={$isLoading}
					/>
					<label for="remember" class="text-muted-foreground text-sm">Remember me</label>
				</div>

				<Button
					class="bg-primary hover:bg-primary/90 text-primary-foreground h-12 w-full transform text-base font-medium shadow-lg transition-all duration-200 hover:-translate-y-0.5 hover:shadow-xl disabled:cursor-not-allowed disabled:opacity-50"
					type="submit"
					disabled={$isLoading || !isFormValid}
				>
					{$isLoading ? 'Signing in...' : 'Sign In'}
				</Button>
			</form>
			<!-- Divider -->
			<div class="my-8 flex items-center">
				<div class="bg-border h-px flex-1"></div>
				<span class="text-muted-foreground px-4 text-sm">or</span>
				<div class="bg-border h-px flex-1"></div>
			</div>

			<!-- Social Login -->
			<div class="space-y-3">
				<Button
					variant="outline"
					class="border-input hover:border-primary/50 hover:bg-primary/5 h-12 w-full border-2 transition-all duration-200"
					type="button"
				>
					<svg class="mr-3 h-5 w-5" viewBox="0 0 24 24">
						<path
							fill="currentColor"
							d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
						/>
						<path
							fill="currentColor"
							d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
						/>
						<path
							fill="currentColor"
							d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
						/>
						<path
							fill="currentColor"
							d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
						/>
					</svg>
					Continue with Google
				</Button>
			</div>

			<!-- Footer -->
			<div class="mt-8 text-center">
				<p class="text-muted-foreground text-sm">
					Don't have an account?
					<a
						href="/signup"
						class="text-primary hover:text-primary/80 font-medium transition-colors"
					>
						Sign up
					</a>
				</p>
			</div>
		</div>
	</div>
</section>
