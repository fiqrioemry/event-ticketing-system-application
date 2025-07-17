<script lang="ts">
	import { goto } from '$app/navigation';
	import FormFooter from './FormFooter.svelte';
	import FormHeader from './FormHeader.svelte';
	import { Loader2, ArrowLeft } from '@lucide/svelte';
	import GoogleProvider from './GoogleProvider.svelte';
	import { registerValidationRules } from '$lib/utils/dto.validation';
	import ErrorMessage from '$lib/components/shared/ErrorMessage.svelte';
	import type { RegisterRequest, VerifyOTPRequest } from '$lib/types/api';
	import SubmitButton from '$lib/components/form-input/SubmitButton.svelte';
	import { authStore, authLoading, authError } from '$lib/stores/auth.store';
	import { validateForm, type ValidationErrors } from '$lib/utils/validation';
	import InputTextElement from '$lib/components/form-input/InputTextElement.svelte';
	import InputEmailElement from '$lib/components/form-input/InputEmailElement.svelte';
	import InputPasswordElement from '$lib/components/form-input/InputPasswordElement.svelte';

	// Registration steps
	const STEPS = {
		REGISTER: 1,
		VERIFY_OTP: 2
	} as const;

	type StepType = (typeof STEPS)[keyof typeof STEPS];

	// State management
	let currentStep = $state<StepType>(STEPS.REGISTER);
	let errors: ValidationErrors = $state({});
	let termsAccepted = $state(false);
	let resendCooldown = $state(0);
	let resendTimer: ReturnType<typeof setInterval> | null = null;

	// Form data
	let formData = $state({
		fullname: '',
		email: '',
		password: '',
		otp: ''
	});

	// Computed properties
	const isRegistrationStep = $derived(currentStep === STEPS.REGISTER);
	const isVerificationStep = $derived(currentStep === STEPS.VERIFY_OTP);
	const canSubmit = $derived(
		isRegistrationStep
			? formData.fullname && formData.email && formData.password && termsAccepted
			: formData.otp.length === 6
	);
	const canResend = $derived(resendCooldown === 0 && !$authLoading);

	// Clear errors when switching steps or input changes
	$effect(() => {
		errors = {};
	});

	// Cleanup timer on component destroy
	$effect(() => {
		return () => {
			if (resendTimer) {
				clearInterval(resendTimer);
			}
		};
	});

	// Form validation rules
	const otpValidationRules = {
		otp: { required: true, minLength: 6, maxLength: 6 }
	};

	// Start resend cooldown timer
	function startResendCooldown() {
		resendCooldown = 60; // 60 seconds cooldown
		resendTimer = setInterval(() => {
			resendCooldown--;
			if (resendCooldown <= 0) {
				if (resendTimer) {
					clearInterval(resendTimer);
					resendTimer = null;
				}
			}
		}, 1000);
	}

	// Event handlers
	async function handleRegistration() {
		// Validate registration form
		errors = validateForm(formData, registerValidationRules);
		if (Object.keys(errors).length > 0) return;

		const registerRequest: RegisterRequest = {
			fullname: formData.fullname.trim(),
			email: formData.email.trim().toLowerCase(),
			password: formData.password
		};

		try {
			const response = await authStore.register(registerRequest);
			if (response.success) {
				currentStep = STEPS.VERIFY_OTP;
				// Clear sensitive data
				formData.password = '';
				// Start cooldown for initial OTP
				startResendCooldown();
			}
		} catch (error) {
			console.error('Registration failed:', error);
		}
	}

	async function handleResendOtp() {
		if (!canResend) return;

		try {
			await authStore.resendOtp(formData.email);
			startResendCooldown();
		} catch (error) {
			console.error('Resend OTP failed:', error);
		}
	}

	async function handleOtpVerification() {
		// Validate OTP
		errors = validateForm({ otp: formData.otp }, otpValidationRules);
		if (Object.keys(errors).length > 0) return;

		const verifyRequest: VerifyOTPRequest = {
			email: formData.email.trim().toLowerCase(),
			otp: formData.otp.trim()
		};

		try {
			await authStore.verifyOtp(verifyRequest);
			// Success is handled in the store (redirect to profile)
		} catch (error) {
			console.error('OTP verification failed:', error);
		}
	}

	async function handleSubmit(event: Event) {
		event.preventDefault();
		if (!canSubmit) return;

		if (isRegistrationStep) {
			await handleRegistration();
		} else {
			await handleOtpVerification();
		}
	}

	function handleBackToRegistration() {
		currentStep = STEPS.REGISTER;
		formData.otp = '';
		authStore.clearError();

		// Clear resend timer
		if (resendTimer) {
			clearInterval(resendTimer);
			resendTimer = null;
		}
		resendCooldown = 0;
	}

	function resetAndNavigateToSignIn() {
		// Reset form
		formData = {
			fullname: '',
			email: '',
			password: '',
			otp: ''
		};
		errors = {};
		termsAccepted = false;
		authStore.clearError();

		// Clear resend timer
		if (resendTimer) {
			clearInterval(resendTimer);
			resendTimer = null;
		}
		resendCooldown = 0;

		goto('/signin');
	}

	// OTP input formatting
	function handleOtpInput(event: Event) {
		const target = event.target as HTMLInputElement;
		const value = target.value.replace(/\D/g, ''); // Only allow digits
		formData.otp = value.slice(0, 6); // Limit to 6 digits
	}
</script>

<div class="w-full max-w-md px-8 lg:px-0">
	{#if $authLoading}
		<div class="flex h-96 items-center justify-center">
			<div class="flex flex-col items-center justify-center gap-4">
				<Loader2 class="h-10 w-10 animate-spin text-blue-500" />
				<p class="text-muted-foreground">
					{isRegistrationStep
						? 'Creating account...'
						: resendCooldown > 0
							? 'Sending OTP...'
							: 'Verifying OTP...'}
				</p>
			</div>
		</div>
	{:else}
		<!-- Dynamic Header -->
		<FormHeader
			title={isRegistrationStep ? 'Create Account' : 'Verify Your Email'}
			subtitle={isRegistrationStep
				? 'Sign up to start booking your favorite events'
				: `We've sent a verification code to ${formData.email}`}
		/>

		<!-- Back button for OTP step -->
		{#if isVerificationStep}
			<button
				type="button"
				onclick={handleBackToRegistration}
				class="mb-4 flex items-center gap-2 text-sm text-blue-600 transition-colors hover:text-blue-800"
				disabled={$authLoading}
			>
				<ArrowLeft class="h-4 w-4" />
				Back to registration
			</button>
		{/if}

		<form onsubmit={handleSubmit} class="space-y-4">
			{#if $authError}
				<ErrorMessage message={$authError.message} onclearError={authStore.clearError} />
			{/if}

			{#if isRegistrationStep}
				<!-- Registration Form -->
				<InputEmailElement
					required
					id="email"
					label="Email Address"
					disabled={$authLoading}
					error={errors.email}
					bind:value={formData.email}
					placeholder="Enter your email address"
				/>

				<InputTextElement
					required
					id="fullname"
					label="Full Name"
					disabled={$authLoading}
					error={errors.fullname}
					bind:value={formData.fullname}
					placeholder="Enter your full name"
				/>

				<InputPasswordElement
					id="password"
					required
					disabled={$authLoading}
					error={errors.password}
					showStrengthIndicator={true}
					bind:value={formData.password}
					placeholder="Enter your password"
				/>

				<!-- Terms and Conditions -->
				<div class="flex items-start space-x-3">
					<input
						id="terms"
						type="checkbox"
						bind:checked={termsAccepted}
						class="mt-1 rounded border-gray-300 text-blue-600 focus:ring-blue-500"
						disabled={$authLoading}
					/>
					<label for="terms" class="text-sm leading-relaxed text-gray-600">
						I agree to the
						<a href="/terms" class="font-medium text-blue-600 hover:text-blue-800">
							Terms of Service
						</a>
						and
						<a href="/privacy" class="font-medium text-blue-600 hover:text-blue-800">
							Privacy Policy
						</a>
					</label>
				</div>
			{:else}
				<!-- OTP Verification Form -->
				<div class="space-y-2">
					<label for="otp" class="block text-sm font-medium text-gray-700">
						Verification Code
					</label>
					<input
						id="otp"
						type="text"
						inputmode="numeric"
						maxlength="6"
						value={formData.otp}
						oninput={handleOtpInput}
						placeholder="Enter 6-digit code"
						disabled={$authLoading}
						class="w-full rounded-lg border border-gray-300 px-4 py-3 text-center font-mono text-2xl tracking-widest focus:border-blue-500 focus:ring-2 focus:ring-blue-500 disabled:bg-gray-50 disabled:text-gray-500"
						class:border-red-500={errors.otp}
					/>
					{#if errors.otp}
						<p class="text-sm text-red-600">{errors.otp}</p>
					{/if}
				</div>

				<!-- Resend OTP -->
				<div class="text-center">
					<p class="text-sm text-gray-600">
						Didn't receive the code?
						{#if canResend}
							<button
								type="button"
								class="font-medium text-blue-600 transition-colors hover:text-blue-800"
								onclick={handleResendOtp}
							>
								Resend
							</button>
						{:else}
							<span class="font-medium text-gray-400">
								Resend in {resendCooldown}s
							</span>
						{/if}
					</p>
				</div>
			{/if}

			<SubmitButton
				variant="primary"
				isLoading={$authLoading}
				disabled={!canSubmit || $authLoading}
				buttonText={isRegistrationStep ? 'Create Account' : 'Verify Email'}
				className="h-12 w-full disabled:bg-gray-100 disabled:cursor-not-allowed"
				buttonLoadingText={isRegistrationStep ? 'Creating account...' : 'Verifying...'}
			/>
		</form>

		{#if isRegistrationStep}
			<!-- Divider -->
			<div class="my-8 flex items-center">
				<div class="bg-border h-px flex-1"></div>
				<span class="text-muted-foreground px-4 text-sm">or</span>
				<div class="bg-border h-px flex-1"></div>
			</div>

			<!-- Social Signup -->
			<GoogleProvider buttonText="Sign up with Google" />
		{/if}

		<!-- Footer -->
		<FormFooter
			buttonText="Sign in"
			onClick={resetAndNavigateToSignIn}
			description="Already have an account?"
		/>
	{/if}
</div>
