<!-- src/lib/components/PasswordInput.svelte -->
<script lang="ts">
	import { Eye, EyeOff, Lock, X } from '@lucide/svelte';
	import BaseField from '$lib/components/form-input/BaseField.svelte';

	export let id: string = '';
	export let error: string = '';
	export let value: string = '';
	export let description: string = '';
	export let label: string = 'Password';
	export let readonly: boolean = false;
	export let required: boolean = false;
	export let disabled: boolean = false;
	export let clearable: boolean = false;
	export let minlength: number | null = 6;
	export let maxlength: number | null = 100;
	export let showStrengthIndicator: boolean = false;
	export let placeholder: string = 'Enter your password';

	let inputElement: HTMLInputElement;
	let showPassword = false;
	let passwordStrength = 0;

	$: inputType = showPassword ? 'text' : 'password';
	$: showClear = clearable && value && value.length > 0 && !readonly && !disabled;
	$: if (showStrengthIndicator && value) {
		passwordStrength = calculatePasswordStrength(value);
	} else {
		passwordStrength = 0;
	}

	function togglePasswordVisibility() {
		showPassword = !showPassword;
		inputElement.focus();
	}

	function handleClear() {
		value = '';
		passwordStrength = 0;
		inputElement.focus();
	}

	function calculatePasswordStrength(password: string): number {
		let strength = 0;

		// Length check
		if (password.length >= 8) strength += 25;
		if (password.length >= 12) strength += 25;

		// Character variety checks
		if (/[a-z]/.test(password)) strength += 10;
		if (/[A-Z]/.test(password)) strength += 10;
		if (/[0-9]/.test(password)) strength += 15;
		if (/[^A-Za-z0-9]/.test(password)) strength += 15;

		return Math.min(strength, 100);
	}

	function getStrengthLabel(strength: number): string {
		if (strength < 30) return 'Weak';
		if (strength < 60) return 'Fair';
		if (strength < 80) return 'Good';
		return 'Strong';
	}

	function getStrengthColor(strength: number): string {
		if (strength < 30) return 'bg-red-500';
		if (strength < 60) return 'bg-yellow-500';
		if (strength < 80) return 'bg-blue-500';
		return 'bg-green-500';
	}
</script>

<BaseField {label} {error} {required} {disabled} {id} {description}>
	<div class="relative">
		<!-- Lock Icon -->
		<div
			class=" text-muted-foreground pointer-events-none absolute top-1/2 left-3 -translate-y-1/2 transform"
		>
			<Lock size={18} />
		</div>

		<input
			bind:this={inputElement}
			bind:value
			type={inputType}
			{id}
			{placeholder}
			{disabled}
			{readonly}
			{maxlength}
			{minlength}
			class:error
			class="bg-background border-border text-foreground placeholder-muted-foreground focus:ring-ring h-12 w-full rounded-lg border
            px-3 py-2
            pr-10
            pl-10 transition-all duration-200 focus:border-transparent
            focus:ring-2 focus:outline-none
            disabled:cursor-not-allowed disabled:opacity-60"
		/>

		<!-- Right Side Controls -->
		<div class="absolute top-1/2 right-3 flex -translate-y-1/2 transform items-center gap-1">
			{#if showClear}
				<button
					type="button"
					class="focus:ring-ring text-muted-foreground hover:text-foreground flex h-6 w-6 items-center justify-center rounded transition-colors focus:ring-2 focus:ring-offset-1 focus:outline-none"
					on:click={handleClear}
					tabindex="-1"
				>
					<X size={16} />
				</button>
			{/if}

			<button
				type="button"
				class="focus:ring-ring text-muted-foreground hover:text-foreground flex h-6 w-6 items-center justify-center rounded transition-colors focus:ring-2 focus:ring-offset-1 focus:outline-none"
				on:click={togglePasswordVisibility}
				tabindex="-1"
			>
				{#if showPassword}
					<EyeOff size={16} />
				{:else}
					<Eye size={16} />
				{/if}
			</button>
		</div>
	</div>

	<!-- Password Strength Indicator -->
	{#if showStrengthIndicator && value}
		<div class=" mt-2 flex items-center gap-2">
			<div class="bg-muted h-2 flex-1 overflow-hidden rounded-full">
				<div
					class="h-full rounded-full transition-all duration-300 {getStrengthColor(
						passwordStrength
					)}"
					style="width: {passwordStrength}%"
				></div>
			</div>
			<span
				class="text-xs font-medium {getStrengthColor(passwordStrength).replace('bg-', 'text-')}"
			>
				{getStrengthLabel(passwordStrength)}
			</span>
		</div>
	{/if}
</BaseField>
