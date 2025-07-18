<!-- src/routes/auth/google/callback/+page.svelte -->
<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { authStore, authLoading } from '$lib/stores/auth.store';

	let error = '';
	let processing = true;

	onMount(async () => {
		const urlParams = new URLSearchParams($page.url.search);
		const code = urlParams.get('code');
		const state = urlParams.get('state');
		const oauthError = urlParams.get('error');

		// Handle OAuth errors
		if (oauthError) {
			error = `OAuth Error: ${oauthError}`;
			processing = false;
			return;
		}

		// Handle missing code
		if (!code) {
			error = 'Missing authorization code';
			processing = false;
			return;
		}

		try {
			// Process Google OAuth login
			await authStore.googleOAuthLogin({
				code,
				state: state || undefined
			});
			// User will be redirected by the store after successful login
		} catch (err: any) {
			console.error('Google OAuth callback error:', err);
			error = err.message || 'Authentication failed';
			processing = false;
		}
	});

	function goBack() {
		goto('/signin');
	}
</script>

<svelte:head>
	<title>Google Authentication</title>
	<meta name="description" content="Authenticate with Google to access your account." />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
</svelte:head>

<div class="flex min-h-screen w-full items-center justify-center bg-gray-50">
	<div class="w-full max-w-md space-y-6 rounded-lg bg-white p-8 shadow-md">
		{#if processing || $authLoading}
			<div class="text-center">
				<div class="mx-auto mb-4 h-12 w-12">
					<svg class="h-12 w-12 animate-spin text-blue-600" viewBox="0 0 24 24">
						<circle
							class="opacity-25"
							cx="12"
							cy="12"
							r="10"
							stroke="currentColor"
							stroke-width="4"
							fill="none"
						/>
						<path
							class="opacity-75"
							fill="currentColor"
							d="m4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
						/>
					</svg>
				</div>
				<h2 class="text-xl font-semibold text-gray-900">Authenticating with Google</h2>
				<p class="mt-2 text-gray-600">Please wait while we complete your login...</p>
			</div>
		{:else if error}
			<div class="text-center">
				<div class="mx-auto mb-4 h-12 w-12">
					<svg class="h-12 w-12 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.082 16.5c-.77.833.192 2.5 1.732 2.5z"
						/>
					</svg>
				</div>
				<h2 class="text-xl font-semibold text-gray-900">Authentication Failed</h2>
				<p class="mt-2 text-gray-600">{error}</p>
				<button
					on:click={goBack}
					class="mt-4 rounded-md bg-blue-600 px-4 py-2 text-white hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:outline-none"
				>
					Back to Sign In
				</button>
			</div>
		{/if}
	</div>
</div>
