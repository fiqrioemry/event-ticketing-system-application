<script lang="ts">
	import { Check } from '@lucide/svelte';
	import { goto } from '$app/navigation';
	import { fly } from 'svelte/transition';
	import type { ResetPasswordRequest } from '$lib/types/api';
	import FormHeader from '$lib/components/auth/FormHeader.svelte';
	import ErrorMessage from '$lib/components/shared/ErrorMessage.svelte';
	import { resetPasswordValidationRules } from '$lib/utils/dto.validation';
	import SubmitButton from '$lib/components/form-input/SubmitButton.svelte';
	import { authStore, authLoading, authError } from '$lib/stores/auth.store';
	import { validateForm, type ValidationErrors } from '$lib/utils/validation';
	import InputPasswordElement from '$lib/components/form-input/InputPasswordElement.svelte';

	export let data: { valid: boolean; token: string };

	let errors: ValidationErrors = {};
	let showSuccess = false;
	let systemErrors: string | null = null;

	let resetForm = {
		newPassword: '',
		confirmPassword: '',
		token: data.token || ''
	};

	async function handleSubmit() {
		errors = {};
		systemErrors = null;

		errors = validateForm(resetForm, resetPasswordValidationRules);

		if (resetForm.newPassword !== resetForm.confirmPassword) {
			errors.confirmPassword = 'Passwords do not match';
		}

		if (Object.keys(errors).length > 0) {
			return;
		}

		const resetRequest: ResetPasswordRequest = {
			token: resetForm.token,
			newPassword: resetForm.newPassword.trim(),
			confirmPassword: resetForm.confirmPassword.trim()
		};

		const response = await authStore.resetPassword(resetRequest);

		if (response.success) {
			showSuccess = true;
		} else {
			systemErrors =
				response.response?.data?.message || response.data?.message || 'Failed to reset password';
		}
	}

	function handleNavigateToLogin() {
		errors = {};
		showSuccess = false;
		systemErrors = null;
		resetForm = {
			newPassword: '',
			confirmPassword: '',
			token: data.token
		};
		goto('/signin');
	}

	function handleClearError() {
		systemErrors = null;
		authStore.clearError();
	}
</script>

<svelte:head>
	<title>Reset Password | TiketKu</title>
	<meta name="description" content="Reset your password on TiketKu." />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<meta property="og:title" content="Tiketku - Book Amazing Events" />
	<meta
		property="og:description"
		content="Book concerts, festivals, workshops & sports events in Indonesia"
	/>
	<meta property="og:image" content="/og-tiketku.jpg" />
</svelte:head>

<div class="flex h-screen w-full items-center justify-center">
	{#if showSuccess}
		<!-- Success State with Animation -->
		<div class="space-y-6 text-center" in:fly={{ y: 30, duration: 600, delay: 100 }}>
			<!-- Success Icon with staggered animation -->
			<div
				class="mx-auto inline-flex h-20 w-20 items-center justify-center rounded-full bg-gradient-to-r from-green-100 to-emerald-100 shadow-lg"
				in:fly={{ y: -20, duration: 500, delay: 200 }}
			>
				<div in:fly={{ duration: 400, delay: 400 }}>
					<Check size={50} class="text-green-600" />
				</div>
			</div>

			<!-- Success Header with staggered animation -->
			<div class="space-y-2" in:fly={{ y: 20, duration: 500, delay: 300 }}>
				<h2 class="text-2xl font-bold text-slate-900 sm:text-3xl">Password Reset Successfully</h2>
				<p class="text-sm leading-relaxed text-slate-600">
					Your password has been updated successfully. You can now sign in with your new password.
				</p>
			</div>

			<!-- Instructions with animation -->
			<div class="space-y-4" in:fly={{ y: 20, duration: 500, delay: 500 }}>
				<div class="rounded-lg border border-green-200 bg-green-50 p-4">
					<div class="flex items-start space-x-3">
						<div class="text-left">
							<p class="text-sm font-medium text-green-900">What's next:</p>
							<ul class="mt-2 space-y-1 text-sm text-green-800">
								<li>• Sign in with your new password</li>
								<li>• Your account is now secure</li>
								<li>• Consider enabling two-factor authentication</li>
							</ul>
						</div>
					</div>
				</div>
			</div>

			<!-- CTA Button with animation -->
			<div class="pt-4" in:fly={{ y: 20, duration: 500, delay: 600 }}>
				<SubmitButton
					className="h-12 w-full"
					variant="primary"
					buttonText="Continue to Sign In"
					onsubmit={handleNavigateToLogin}
				/>
			</div>
		</div>
	{:else}
		<!-- Reset Password Form with exit animation -->
		<div in:fly={{ y: 0, duration: 300 }} out:fly={{ y: -30, duration: 300 }}>
			<FormHeader
				title="Reset your password"
				subtitle="Enter your new password below to complete the reset process"
			/>

			<form on:submit|preventDefault={handleSubmit} class="space-y-4">
				{#if $authError}
					<ErrorMessage message={$authError} onclearError={authStore.clearError} />
				{/if}

				{#if systemErrors}
					<ErrorMessage message={systemErrors} onclearError={handleClearError} />
				{/if}

				<InputPasswordElement
					required
					id="password"
					minlength={8}
					maxlength={100}
					label="New Password"
					disabled={$authLoading}
					error={errors.newPassword}
					showStrengthIndicator={true}
					bind:value={resetForm.newPassword}
					placeholder="Enter your new password"
				/>

				<InputPasswordElement
					required
					id="confirm-password"
					disabled={$authLoading}
					label="Confirm New Password"
					error={errors.confirmPassword}
					bind:value={resetForm.confirmPassword}
					placeholder="Confirm your new password"
				/>

				<SubmitButton
					className="h-12 w-full"
					variant="primary"
					isLoading={$authLoading}
					buttonText="Reset Password"
					buttonLoadingText="Resetting..."
				/>
			</form>
		</div>
	{/if}
</div>
