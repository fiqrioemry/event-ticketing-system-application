<!-- src/lib/components/EmailInput.svelte -->
<script lang="ts">
	import { Mail, X, Check, AlertCircle } from '@lucide/svelte';
	import BaseField from '$lib/components/form-input/BaseField.svelte';

	export let id: string = '';
	export let error: string = '';
	export let value: string = '';
	export let description: string = '';
	export let readonly: boolean = false;
	export let required: boolean = false;
	export let disabled: boolean = false;
	export let clearable: boolean = true;
	export let maxlength: number | null = 254; // RFC 5321 limit
	export let label: string = 'Email Address';
	export let placeholder: string = 'Enter your email address';
	export let suggestions: string[] = ['gmail.com', 'yahoo.com', 'outlook.com', 'hotmail.com'];

	let inputElement: HTMLInputElement;
	let isValidEmail = false;
	let showSuggestions = false;
	let filteredSuggestions: string[] = [];

	$: showClear = clearable && value && value.length > 0 && !readonly && !disabled;
	$: isValidEmail = validateEmail(value);
	$: if (value && value.includes('@') && !value.includes('@.')) {
		const domain = value.split('@')[1];
		if (domain && domain.length > 0) {
			filteredSuggestions = suggestions
				.filter((suggestion) => suggestion.toLowerCase().startsWith(domain.toLowerCase()))
				.slice(0, 3);
			showSuggestions = filteredSuggestions.length > 0 && !suggestions.includes(domain);
		} else {
			showSuggestions = false;
		}
	} else {
		showSuggestions = false;
	}

	function validateEmail(email: string): boolean {
		if (!email) return false;
		const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
		return emailRegex.test(email);
	}

	function handleClear() {
		value = '';
		inputElement.focus();
		showSuggestions = false;
	}

	function applySuggestion(suggestion: string) {
		const localPart = value.split('@')[0];
		value = `${localPart}@${suggestion}`;
		showSuggestions = false;
		inputElement.focus();
	}

	function handleBlur() {
		setTimeout(() => {
			showSuggestions = false;
		}, 150);
	}

	function handleFocus() {
		if (value && value.includes('@')) {
			const domain = value.split('@')[1];
			if (domain && domain.length > 0) {
				filteredSuggestions = suggestions
					.filter((suggestion) => suggestion.toLowerCase().startsWith(domain.toLowerCase()))
					.slice(0, 3);
				showSuggestions = filteredSuggestions.length > 0 && !suggestions.includes(domain);
			}
		}
	}
</script>

<BaseField {label} {error} {required} {disabled} {id} {description}>
	<div class="relative">
		<!-- Mail Icon -->
		<div
			class="text-muted-foreground pointer-events-none absolute top-1/2 left-3 -translate-y-1/2 transform"
		>
			<Mail size={18} />
		</div>

		<input
			bind:this={inputElement}
			bind:value
			type="email"
			{id}
			{placeholder}
			{disabled}
			{readonly}
			{maxlength}
			class:error
			on:blur={handleBlur}
			on:focus={handleFocus}
			class:valid={isValidEmail && !error}
			class="bg-background border-border text-foreground placeholder-muted-foreground focus:ring-ring h-12 w-full rounded-lg
            border px-3 py-2 pr-10 pl-10 transition-all
            duration-200 focus:border-transparent
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

			<!-- Validation Indicator -->
			{#if value && !error}
				<div class="justify-center; flex h-6 w-6 items-center">
					{#if isValidEmail}
						<Check size={16} class="text-green-500" />
					{:else}
						<AlertCircle size={16} class="text-yellow-500" />
					{/if}
				</div>
			{/if}
		</div>

		<!-- Email Suggestions -->
		{#if showSuggestions}
			<div
				class="bg-popover border-border absolute top-full right-0 left-0 z-50 mt-1 overflow-hidden rounded-lg border shadow-lg"
			>
				<div class="max-h-32 overflow-y-auto">
					{#each filteredSuggestions as suggestion}
						<button
							type="button"
							class="text-muted-foreground hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground flex w-full items-center gap-2 px-3 py-2 text-left text-sm transition-colors focus:outline-none"
							on:click={() => applySuggestion(suggestion)}
						>
							<Mail size={14} />
							<span>{value.split('@')[0]}@{suggestion}</span>
						</button>
					{/each}
				</div>
			</div>
		{/if}
	</div>
</BaseField>
