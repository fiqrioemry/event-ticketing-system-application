<script lang="ts">
	import { goto } from '$app/navigation';
	import FormHeader from './FormHeader.svelte';
	import FormFooter from './FormFooter.svelte';
	import type { LoginRequest } from '$lib/types/api';
	import GoogleProvider from './GoogleProvider.svelte';
	import { loginValidationRules } from '$lib/utils/dto.validation';
	import ErrorMessage from '$lib/components/shared/ErrorMessage.svelte';
	import SubmitButton from '$lib/components/form-input/SubmitButton.svelte';
	import { authStore, authLoading, authError } from '$lib/stores/auth.store';
	import { validateForm, type ValidationErrors } from '$lib/utils/validation';
	import InputEmailElement from '$lib/components/form-input/InputEmailElement.svelte';
	import InputPasswordElement from '$lib/components/form-input/InputPasswordElement.svelte';
	import DemoAccount from './DemoAccount.svelte';

	let errors: ValidationErrors = {};

	let loginForm = {
		email: '',
		password: ''
	};

	async function handleSubmit() {
		errors = {};

		errors = validateForm(loginForm, loginValidationRules);
		if (Object.keys(errors).length > 0) {
			return;
		}

		const loginRequest: LoginRequest = {
			email: loginForm.email.trim(),
			password: loginForm.password
		};

		await authStore.login(loginRequest);
	}

	function handleNavigate() {
		errors = {};

		loginForm = {
			email: '',
			password: ''
		};
		goto('/signup');
	}
	console.log($authError);
</script>

<DemoAccount />
<div class="w-full max-w-md px-8 lg:px-0">
	<!-- Header -->
	<FormHeader title="Welcome Back" subtitle="Login to continue booking your favorite events" />

	<form on:submit|preventDefault={handleSubmit} class="space-y-4">
		{#if $authError}
			<ErrorMessage message={$authError} onclearError={authStore.clearError} />
		{/if}

		<InputEmailElement
			required
			id="email"
			label="Email Address"
			error={errors.email}
			disabled={$authLoading}
			bind:value={loginForm.email}
			placeholder="Enter your email address"
		/>

		<InputPasswordElement
			id="password"
			required
			disabled={$authLoading}
			error={errors.password}
			bind:value={loginForm.password}
		/>

		<div class="flex justify-end">
			<a href="/forgot-password" class="text-sm text-blue-600 hover:underline">Forgot password?</a>
		</div>
		<SubmitButton
			className="h-12 w-full"
			variant="primary"
			isLoading={$authLoading}
			buttonText="Sign In"
			buttonLoadingText="Signing in..."
		/>
	</form>

	<!-- Divider -->
	<div class="my-8 flex items-center">
		<div class="bg-border h-px flex-1"></div>
		<span class="text-muted-foreground px-4 text-sm">or</span>
		<div class="bg-border h-px flex-1"></div>
	</div>

	<!-- Social Login -->
	<GoogleProvider />

	<!-- Footer -->
	<FormFooter buttonText="Sign up" onClick={handleNavigate} description="Don't have an account?" />
</div>
