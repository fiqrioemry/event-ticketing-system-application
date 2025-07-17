<!-- src/lib/components/forms/NumberInput.svelte -->
<script lang="ts">
	import { ChevronUp, ChevronDown, PhoneCall, DollarSign } from '@lucide/svelte';

	// Basic props
	export let id: string = '';
	export let label: string = '';
	export let value: number | null = null;
	export let placeholder: string = '';
	export let error: string = '';
	export let disabled: boolean = false;
	export let required: boolean = false;
	export let readonly: boolean = false;
	export let min: number | null = null;
	export let max: number | null = null;
	export let step: number = 1;
	export let showSpinButtons: boolean = true;

	// Formatting props
	export let format: 'none' | 'currency' | 'phone' | 'decimal' | 'percentage' = 'none';
	export let currencySymbol: string = 'Rp';
	export let decimals: number = 0;
	export let allowNegative: boolean = true;

	// Validation props
	export let validationRules: Array<(value: number | null) => string | null> = [];
	export let validateOnInput: boolean = false;
	export let validateOnBlur: boolean = true;

	// Callback for live calculation (replacement for dispatch)
	export let onValueChange: ((value: number | null) => void) | null = null;
	export let onValidation: ((error: string | null) => void) | null = null;

	let inputElement: HTMLInputElement;
	let displayValue = '';
	let isFocused = false;
	let validationError = '';

	// Reactive: Update display value when value changes or focus state changes
	$: updateDisplayValue(value, isFocused);

	function updateDisplayValue(val: number | null, focused: boolean) {
		if (focused) {
			// Show raw value when focused for easier editing
			displayValue = val !== null ? val.toString() : '';
		} else {
			// Show formatted value when not focused
			displayValue = val !== null ? formatNumber(val) : '';
		}
	}

	function formatNumber(num: number): string {
		switch (format) {
			case 'currency':
				return formatCurrency(num);
			case 'phone':
				return formatPhone(num);
			case 'decimal':
				return formatDecimal(num);
			case 'percentage':
				return formatPercentage(num);
			default:
				return num.toString();
		}
	}

	function formatCurrency(num: number): string {
		const formatted = new Intl.NumberFormat('id-ID', {
			minimumFractionDigits: decimals,
			maximumFractionDigits: decimals
		}).format(Math.abs(num));

		const sign = num < 0 ? '-' : '';
		return `${sign}${currencySymbol} ${formatted}`;
	}

	function formatPhone(num: number): string {
		const phoneStr = Math.abs(num).toString();
		// Format: 0812-3456-7890 or +62-812-3456-7890
		if (phoneStr.length >= 10) {
			if (phoneStr.startsWith('62')) {
				return `+${phoneStr.slice(0, 2)}-${phoneStr.slice(2, 5)}-${phoneStr.slice(5, 9)}-${phoneStr.slice(9)}`;
			} else if (phoneStr.startsWith('0')) {
				return `${phoneStr.slice(0, 4)}-${phoneStr.slice(4, 8)}-${phoneStr.slice(8)}`;
			}
		}
		return phoneStr;
	}

	function formatDecimal(num: number): string {
		return new Intl.NumberFormat('id-ID', {
			minimumFractionDigits: decimals,
			maximumFractionDigits: decimals
		}).format(num);
	}

	function formatPercentage(num: number): string {
		return `${num.toFixed(decimals)}%`;
	}

	function parseInputValue(inputValue: string): number | null {
		if (!inputValue.trim()) return null;

		// Remove formatting characters
		let cleanValue = inputValue
			.replace(new RegExp(currencySymbol, 'g'), '')
			.replace(/[\s+\-()]/g, '')
			.replace(/[.,]/g, (match) => (match === ',' ? '.' : ''));

		// Remove percentage sign
		if (format === 'percentage') {
			cleanValue = cleanValue.replace('%', '');
		}

		const parsed = parseFloat(cleanValue);
		return isNaN(parsed) ? null : parsed;
	}

	function runValidation(val: number | null): string | null {
		// Built-in validations
		if (required && (val === null || val === undefined)) {
			return 'This field is required';
		}

		if (val !== null) {
			if (min !== null && val < min) {
				return `Value must be at least ${formatNumber(min)}`;
			}
			if (max !== null && val > max) {
				return `Value must not exceed ${formatNumber(max)}`;
			}
			if (!allowNegative && val < 0) {
				return 'Negative values are not allowed';
			}
		}

		// Custom validations
		for (const rule of validationRules) {
			const customError = rule(val);
			if (customError) return customError;
		}

		return null;
	}

	function handleInput(event: Event) {
		const target = event.target as HTMLInputElement;
		const inputValue = target.value;

		const parsedValue = parseInputValue(inputValue);
		value = parsedValue;

		// Live calculation callback
		if (onValueChange) {
			onValueChange(value);
		}

		// Validate on input if enabled
		if (validateOnInput) {
			validationError = runValidation(value) || '';
			if (onValidation) {
				onValidation(validationError || null);
			}
		}
	}

	function handleFocus() {
		isFocused = true;
	}

	function handleBlur() {
		isFocused = false;

		// Validate on blur if enabled
		if (validateOnBlur) {
			validationError = runValidation(value) || '';
			if (onValidation) {
				onValidation(validationError || null);
			}
		}
	}

	function handleKeydown(event: KeyboardEvent) {
		const allowedKeys = [
			'Backspace',
			'Delete',
			'Tab',
			'Escape',
			'Enter',
			'ArrowLeft',
			'ArrowRight',
			'ArrowUp',
			'ArrowDown'
		];
		const isNumber = /^[0-9]$/.test(event.key);
		const isDecimal = event.key === '.' || event.key === ',';
		const isMinus = event.key === '-' && allowNegative;
		const isPlus = event.key === '+' && format === 'phone';

		if (!allowedKeys.includes(event.key) && !isNumber && !isDecimal && !isMinus && !isPlus) {
			event.preventDefault();
		}
	}

	function increment() {
		if (disabled || readonly) return;

		const currentValue = value || 0;
		const newValue = currentValue + step;

		if (max === null || newValue <= max) {
			value = newValue;
			if (onValueChange) onValueChange(value);
		}
	}

	function decrement() {
		if (disabled || readonly) return;

		const currentValue = value || 0;
		const newValue = currentValue - step;

		if ((min === null || newValue >= min) && (allowNegative || newValue >= 0)) {
			value = newValue;
			if (onValueChange) onValueChange(value);
		}
	}

	// Reactive: Update error display
	$: displayError = error || validationError;
</script>

<div class="flex w-full flex-col gap-2">
	<!-- Label -->
	{#if label}
		<label for={id} class="text-foreground text-sm font-medium">
			{label}
			{#if required}
				<span class="text-destructive">*</span>
			{/if}
		</label>
	{/if}

	<!-- Input Container -->
	<div class="relative">
		<input
			bind:this={inputElement}
			type="text"
			inputmode={format === 'phone' ? 'tel' : 'numeric'}
			{id}
			{placeholder}
			{disabled}
			{readonly}
			{required}
			bind:value={displayValue}
			class="bg-input border-border text-foreground placeholder:text-muted-foreground focus:ring-ring h-12 w-full rounded-lg border px-3 py-2 font-mono text-sm transition-colors focus:border-transparent focus:ring-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
			class:pr-8={showSpinButtons && format !== 'phone'}
			class:border-destructive={displayError}
			class:focus:ring-destructive={displayError}
			class:bg-muted={readonly}
			class:cursor-default={readonly}
			on:input={handleInput}
			on:keydown={handleKeydown}
			on:focus={handleFocus}
			on:blur={handleBlur}
		/>

		<!-- Spin Buttons (not for phone format) -->
		{#if showSpinButtons && !readonly && format !== 'phone'}
			<div class="absolute top-1/2 right-1 flex -translate-y-1/2 transform flex-col">
				<button
					type="button"
					class="bg-muted hover:bg-muted-foreground/20 text-muted-foreground border-border focus:ring-ring flex h-4 w-6 items-center justify-center rounded-t border border-b-0 text-xs transition-colors focus:ring-1 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
					on:click={increment}
					{disabled}
					tabindex="-1"
				>
					<ChevronUp size={12} />
				</button>
				<button
					type="button"
					class="bg-muted hover:bg-muted-foreground/20 text-muted-foreground border-border focus:ring-ring flex h-4 w-6 items-center justify-center rounded-b border text-xs transition-colors focus:ring-1 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
					on:click={decrement}
					{disabled}
					tabindex="-1"
				>
					<ChevronDown size={12} />
				</button>
			</div>
		{/if}

		<!-- Format indicator -->
		{#if format !== 'none' && !isFocused && value !== null}
			<div class="pointer-events-none absolute top-1/2 right-2 -translate-y-1/2 transform">
				{#if format === 'currency'}
					<DollarSign class="h-4 w-4" />
				{:else if format === 'phone'}
					<PhoneCall class="h-4 w-4" />
				{:else if format === 'percentage'}
					<span></span>
				{/if}
			</div>
		{/if}
	</div>

	<!-- Error Message -->
	{#if displayError}
		<span class="text-destructive text-sm">{displayError}</span>
	{/if}

	<!-- Format helper text -->
	{#if format === 'phone' && !displayError}
		<span class="text-muted-foreground text-xs">Format: 08123456789 atau 628123456789</span>
	{/if}
</div>

<style>
	/* Custom styling yang tidak bisa dilakukan Tailwind */
	input::-webkit-outer-spin-button,
	input::-webkit-inner-spin-button {
		-webkit-appearance: none;
		margin: 0;
	}
</style>
