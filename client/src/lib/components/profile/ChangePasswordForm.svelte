<!-- src/lib/components/forms/ChangePasswordForm.svelte -->
<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import ErrorMessage from '$lib/components/shared/ErrorMessage.svelte';
	import SubmitButton from '$lib/components/form-input/SubmitButton.svelte';
	import { userStore, isUpdating, userError } from '$lib/stores/user.store';
	import InputPasswordElement from '$lib/components/form-input/InputPasswordElement.svelte';
	import { validateForm, type ValidationRules, type ValidationErrors } from '$lib/utils/validation';

	// callback props
	export let onSuccess: (() => void) | null = null;
	export let onCancel: (() => void) | null = null;

	let systemErrors: string | null = null;
	let errors: ValidationErrors = {};

	// set form data
	let formData = {
		currentPassword: '',
		newPassword: '',
		confirmPassword: ''
	};

	// validation rules
	const validationRules: ValidationRules = {
		currentPassword: { required: true },
		newPassword: { required: true, minLength: 8 },
		confirmPassword: { required: true, minLength: 8 }
	};

	async function handleSubmit() {
		errors = {};
		systemErrors = null;

		// Validate form
		errors = validateForm(formData, validationRules);

		// Check password confirmation
		if (formData.newPassword !== formData.confirmPassword) {
			errors.confirmPassword = 'Passwords do not match';
		}

		// Check if new password is different
		if (formData.currentPassword === formData.newPassword) {
			errors.newPassword = 'New password must be different from current password';
		}

		if (Object.keys(errors).length > 0) {
			return;
		}

		// assign to variable
		const changePasswordRequest = {
			currentPassword: formData.currentPassword,
			newPassword: formData.newPassword,
			confirmPassword: formData.confirmPassword
		};

		// send request
		const result = await userStore.changePassword(changePasswordRequest);

		if (result.success) {
			// Clear form on success
			formData = {
				currentPassword: '',
				newPassword: '',
				confirmPassword: ''
			};
			errors = {};

			// Call success callback
			if (onSuccess) onSuccess();
		} else {
			systemErrors =
				result.response.data.message || result.data?.message || 'Failed to change password';
		}
	}

	function handleCancel() {
		formData = {
			currentPassword: '',
			newPassword: '',
			confirmPassword: ''
		};
		errors = {};
		userStore.clearError();

		if (onCancel) onCancel();
	}

	function handleClearError() {
		systemErrors = '';
	}
</script>

<form on:submit|preventDefault={handleSubmit} class="space-y-4">
	<!-- Error from store -->
	{#if systemErrors}
		<ErrorMessage message={systemErrors} onclearError={handleClearError} />
	{/if}

	<InputPasswordElement
		id="current-password"
		label="Current Password"
		disabled={$isUpdating}
		error={errors.currentPassword}
		bind:value={formData.currentPassword}
		placeholder="Enter your current password"
		required
	/>

	<InputPasswordElement
		required
		id="password"
		minlength={8}
		maxlength={100}
		label="New Password"
		disabled={$isUpdating}
		error={errors.newPassword}
		showStrengthIndicator={true}
		bind:value={formData.newPassword}
		placeholder="Enter your new password"
	/>

	<InputPasswordElement
		required
		id="confirm-password"
		disabled={$isUpdating}
		label="Confirm New Password"
		error={errors.confirmPassword}
		bind:value={formData.confirmPassword}
		placeholder="Confirm your new password"
	/>

	<div class="flex gap-2 pt-4">
		<Button
			type="button"
			class="w-1/2"
			variant="outline"
			onclick={handleCancel}
			disabled={$isUpdating}
		>
			Cancel
		</Button>
		<SubmitButton
			className="w-1/2"
			isLoading={$isUpdating}
			buttonText="Update Password"
			buttonLoadingText="Updating..."
		/>
	</div>
</form>
