<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import { authStore, authLoading } from '$lib/stores/auth.store';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';

	export let buttonText = 'Continue with Google';

	// Handle Google OAuth callback when component mounts
	onMount(() => {
		const urlParams = new URLSearchParams($page.url.search);
		const code = urlParams.get('code');
		const state = urlParams.get('state');
		const error = urlParams.get('error');

		// If there's an error from Google OAuth
		if (error) {
			console.error('Google OAuth error:', error);
			return;
		}

		// If we have a code, process the OAuth callback
		if (code) {
			handleGoogleCallback(code, state);
		}
	});

	async function handleGoogleLogin() {
		try {
			// Get Google OAuth URL and redirect user
			const googleOAuthUrl = authStore.getGoogleOAuthUrl();
			window.location.href = googleOAuthUrl;
		} catch (error) {
			console.error('Failed to initiate Google login:', error);
		}
	}

	async function handleGoogleCallback(code: string, state: string | null) {
		try {
			await authStore.googleOAuthLogin({
				code,
				state: state || undefined
			});
			// User will be redirected by the store after successful login
		} catch (error) {
			console.error('Google login failed:', error);
			// Clean up URL params on error
			goto($page.url.pathname, { replaceState: true });
		}
	}
</script>

<div class="space-y-3">
	<Button
		variant="event-outline"
		class="h-12 w-full"
		type="button"
		disabled={$authLoading}
		onclick={handleGoogleLogin}
	>
		{#if $authLoading}
			<svg class="mr-3 h-5 w-5 animate-spin" viewBox="0 0 24 24">
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
			Processing...
		{:else}
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
			{buttonText}
		{/if}
	</Button>
</div>
