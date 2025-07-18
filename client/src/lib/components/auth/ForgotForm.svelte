<script lang="ts">
	import { Check } from '@lucide/svelte';
	import { goto } from '$app/navigation';
	import { fly } from 'svelte/transition';
	import FormHeader from './FormHeader.svelte';
	import FormFooter from './FormFooter.svelte';
	import type { ForgotPasswordRequest } from '$lib/types/api';
	import ErrorMessage from '$lib/components/shared/ErrorMessage.svelte';
	import { forgotPasswordValidationRules } from '$lib/utils/dto.validation';
	import SubmitButton from '$lib/components/form-input/SubmitButton.svelte';
	import { authStore, authLoading, authError } from '$lib/stores/auth.store';
	import { validateForm, type ValidationErrors } from '$lib/utils/validation';
	import InputEmailElement from '$lib/components/form-input/InputEmailElement.svelte';

	let errors: ValidationErrors = {};
	let showSuccess = false;

	let forgotForm = {
		email: ''
	};

	async function handleSubmit() {
		errors = {};

		errors = validateForm(forgotForm, forgotPasswordValidationRules);
		if (Object.keys(errors).length > 0) {
			return;
		}

		const forgotRequest: ForgotPasswordRequest = {
			email: forgotForm.email.trim()
		};

		const response = await authStore.forgotPassword(forgotRequest);

		if (response.success) {
			showSuccess = true;
		}
	}

	function handleNavigate() {
		errors = {};
		showSuccess = false;
		forgotForm = {
			email: ''
		};
		goto('/signin');
	}
</script>

<div class="w-full max-w-md px-8 lg:px-0">
	<!-- Header -->
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
				<h2 class="text-2xl font-bold text-slate-900 sm:text-3xl">Check your email</h2>
				<p class="text-sm leading-relaxed text-slate-600">We've sent a password reset link to</p>
				<p class="text-base font-semibold break-words text-slate-900">
					{forgotForm.email || 'your email address'}
				</p>
			</div>

			<!-- Instructions with animation -->
			<div class="space-y-4" in:fly={{ y: 20, duration: 500, delay: 500 }}>
				<div class="rounded-lg border border-blue-200 bg-blue-50 p-4">
					<div class="flex items-start space-x-3">
						<div class="text-left">
							<p class="text-sm font-medium text-blue-900">Next steps:</p>
							<ul class="mt-2 space-y-1 text-sm text-blue-800">
								<li>• Click the link in your email to reset your password</li>
								<li>• The link will expire in 24 hours</li>
								<li>• Check your spam folder if you don't see it</li>
							</ul>
						</div>
					</div>
				</div>
			</div>
		</div>
	{:else}
		<FormHeader
			title="Forgot your password?"
			subtitle="Enter your email address and we'll send you a link to reset your password"
		/>

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
				bind:value={forgotForm.email}
				placeholder="Enter your email address"
			/>
			<SubmitButton
				className="h-12 w-full"
				variant="primary"
				isLoading={$authLoading}
				buttonText="Send reset link"
				buttonLoadingText="sending..."
			/>
		</form>
	{/if}
	<!-- Footer -->
	<FormFooter
		buttonText="Sign in instead"
		onClick={handleNavigate}
		description="Remember your password?"
	/>
</div>
