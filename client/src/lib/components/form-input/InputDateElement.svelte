<script lang="ts">
	import { formatDate } from '$lib/utils/formatter';
	import { Calendar, ChevronLeft, ChevronRight, X } from '@lucide/svelte';

	export let value: string = '';
	export let id: string = '';
	export let label: string = '';
	export let minDate: string = '';
	export let maxDate: string = '';
	export let required: boolean = false;
	export let disabled: boolean = false;
	export let placeholder: string = 'Select date';
	export let onchange: (value: string) => void = () => {};

	// New props for start/end date functionality
	export let isEndDate: boolean = false;
	export let startDate: string = '';

	let isOpen = false;
	let viewDate = value ? new Date(value) : new Date();

	// Get today's date string in local timezone - FIXED
	const getTodayString = (): string => {
		const today = new Date();
		const year = today.getFullYear();
		const month = String(today.getMonth() + 1).padStart(2, '0');
		const day = String(today.getDate()).padStart(2, '0');
		return `${year}-${month}-${day}`;
	};

	// Format date for input value (YYYY-MM-DD) in local timezone - FIXED
	const formatDateValue = (date: Date): string => {
		const year = date.getFullYear();
		const month = String(date.getMonth() + 1).padStart(2, '0');
		const day = String(date.getDate()).padStart(2, '0');
		return `${year}-${month}-${day}`;
	};

	// Reactive display text - FIXED
	$: displayText = (() => {
		if (value) {
			try {
				// Parse date in local timezone to avoid timezone shift
				const [year, month, day] = value.split('-').map(Number);
				const date = new Date(year, month - 1, day);
				return formatDate(date);
			} catch {
				return placeholder;
			}
		}
		// If no value selected, show placeholder
		return placeholder;
	})();

	// Calendar logic
	const getDaysInMonth = (date: Date): number => {
		return new Date(date.getFullYear(), date.getMonth() + 1, 0).getDate();
	};

	const getFirstDayOfMonth = (date: Date): number => {
		return new Date(date.getFullYear(), date.getMonth(), 1).getDay();
	};

	const isDateDisabled = (date: Date): boolean => {
		if (disabled) return true;

		const dateString = formatDateValue(date);
		const today = getTodayString();

		// Always disable dates before today
		if (dateString < today) return true;

		// For end date, disable dates before start date
		if (isEndDate && startDate && dateString < startDate) return true;

		// Check custom min/max dates
		if (minDate && dateString < minDate) return true;
		if (maxDate && dateString > maxDate) return true;

		return false;
	};

	const isDateSelected = (date: Date): boolean => {
		if (!value) return false;
		return formatDateValue(date) === value;
	};

	// Check if date is today - FIXED
	const isToday = (date: Date): boolean => {
		const today = new Date();
		return (
			date.getFullYear() === today.getFullYear() &&
			date.getMonth() === today.getMonth() &&
			date.getDate() === today.getDate()
		);
	};

	const generateCalendarDays = (): (Date | null)[] => {
		const daysInMonth = getDaysInMonth(viewDate);
		const firstDayOfMonth = getFirstDayOfMonth(viewDate);
		const days: (Date | null)[] = [];

		// Add empty cells for days before the first day of the month
		for (let i = 0; i < firstDayOfMonth; i++) {
			days.push(null);
		}

		// Add days of the month
		for (let day = 1; day <= daysInMonth; day++) {
			days.push(new Date(viewDate.getFullYear(), viewDate.getMonth(), day));
		}

		return days;
	};

	const selectDate = (date: Date) => {
		if (isDateDisabled(date)) return;

		value = formatDateValue(date);
		onchange(value);
		isOpen = false;
	};

	const navigateMonth = (direction: 'prev' | 'next') => {
		const newDate = new Date(viewDate);
		if (direction === 'prev') {
			newDate.setMonth(newDate.getMonth() - 1);
		} else {
			newDate.setMonth(newDate.getMonth() + 1);
		}
		viewDate = newDate;
	};

	const clearDate = (event: Event) => {
		event.stopPropagation();
		value = '';
		onchange('');
	};

	// Close calendar when clicking outside
	const handleClickOutside = (event: MouseEvent) => {
		if (!event.target) return;

		const target = event.target as HTMLElement;
		const calendar = document.getElementById(`calendar-${id}`);
		const trigger = document.getElementById(`trigger-${id}`);

		if (calendar && !calendar.contains(target) && trigger && !trigger.contains(target)) {
			isOpen = false;
		}
	};

	// Update viewDate when value changes - FIXED
	$: if (value) {
		try {
			const [year, month, day] = value.split('-').map(Number);
			viewDate = new Date(year, month - 1, day);
		} catch {
			// Invalid date, keep current viewDate
		}
	}

	$: if (isOpen) {
		document.addEventListener('click', handleClickOutside);
	} else {
		document.removeEventListener('click', handleClickOutside);
	}
</script>

<div class="relative">
	{#if label}
		<label for={id} class="mb-2 block text-sm font-medium text-gray-700">
			{label}
			{#if required}<span class="ml-1 text-red-500">*</span>{/if}
		</label>
	{/if}

	<!-- Trigger Button -->
	<div
		id="trigger-{id}"
		class="flex h-12 w-full items-center justify-between rounded-lg border border-gray-300 bg-white px-3 py-2 text-left text-sm transition-colors focus-within:border-blue-500 focus-within:ring-2 focus-within:ring-blue-500/20 hover:border-blue-400"
		class:border-blue-500={isOpen}
		class:ring-2={isOpen}
		class:ring-blue-500={isOpen}
		class:opacity-50={disabled}
		class:cursor-not-allowed={disabled}
	>
		<button
			{id}
			type="button"
			on:click={() => !disabled && (isOpen = !isOpen)}
			class="flex flex-1 items-center gap-2 focus:outline-none"
			{disabled}
		>
			<Calendar class="h-4 w-4 text-gray-500" />
			<span class={value ? 'text-gray-900' : 'text-gray-500'}>
				{displayText}
			</span>
		</button>

		{#if value && !disabled}
			<button
				type="button"
				on:click={clearDate}
				class="rounded p-1 hover:bg-gray-100 focus:outline-none"
				aria-label="Clear date"
			>
				<X class="h-4 w-4 text-gray-500" />
			</button>
		{/if}
	</div>

	<!-- Calendar Dropdown -->
	{#if isOpen}
		<div
			id="calendar-{id}"
			class="absolute z-50 mt-1 w-80 rounded-lg border border-gray-200 bg-white p-4 shadow-lg"
		>
			<!-- Calendar Header -->
			<div class="mb-4 flex items-center justify-between">
				<button
					type="button"
					on:click={() => navigateMonth('prev')}
					class="rounded p-2 hover:bg-gray-100 focus:outline-none"
				>
					<ChevronLeft class="h-5 w-5 text-gray-600" />
				</button>

				<h2 class="text-lg font-semibold text-gray-900">
					{viewDate.toLocaleDateString('en-US', { month: 'long', year: 'numeric' })}
				</h2>

				<button
					type="button"
					on:click={() => navigateMonth('next')}
					class="rounded p-2 hover:bg-gray-100 focus:outline-none"
				>
					<ChevronRight class="h-5 w-5 text-gray-600" />
				</button>
			</div>

			<!-- Calendar Grid -->
			<div class="mb-4 grid grid-cols-7 gap-1">
				<!-- Day headers -->
				{#each ['Su', 'Mo', 'Tu', 'We', 'Th', 'Fr', 'Sa'] as day}
					<div class="p-2 text-center text-xs font-medium text-gray-500">
						{day}
					</div>
				{/each}

				<!-- Calendar days -->
				{#each generateCalendarDays() as day}
					{#if day}
						<button
							type="button"
							on:click={() => selectDate(day)}
							class="rounded-md p-2 text-center text-sm transition-colors hover:bg-blue-50 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
							class:bg-blue-600={isDateSelected(day)}
							class:text-white={isDateSelected(day)}
							class:hover:bg-blue-700={isDateSelected(day)}
							class:bg-blue-100={isToday(day) && !isDateSelected(day)}
							class:text-blue-600={isToday(day) && !isDateSelected(day)}
							class:font-medium={isToday(day)}
							disabled={isDateDisabled(day)}
						>
							{day.getDate()}
						</button>
					{:else}
						<div class="p-2"></div>
					{/if}
				{/each}
			</div>

			<!-- Quick Actions -->
			<div class="flex items-center justify-between border-t pt-4">
				<button
					type="button"
					on:click={() => selectDate(new Date())}
					class="text-sm font-medium text-blue-600 hover:text-blue-800 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
					disabled={isDateDisabled(new Date())}
				>
					Today
				</button>

				<button
					type="button"
					on:click={() => (isOpen = false)}
					class="rounded bg-gray-100 px-3 py-1 text-sm text-gray-700 hover:bg-gray-200 focus:outline-none"
				>
					Cancel
				</button>
			</div>
		</div>
	{/if}
</div>
