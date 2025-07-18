<!-- src/routes/(public)/events/+page.svelte -->
<script lang="ts">
	import { goto } from '$app/navigation';
	import { Filter, FilterX } from '@lucide/svelte';
	import { formatDate } from '$lib/utils/formatter.js';
	import { locationOptions } from '$lib/utils/state.js';
	import Button from '$lib/components/ui/button/button.svelte';
	import Pagination from '$lib/components/shared/Pagination.svelte';
	import EventCard from '$lib/components/events/EventCard.svelte';
	import EventsLoading from '$lib/components/events/EventsLoading.svelte';
	import EventNotFound from '$lib/components/events/EventNotFound.svelte';
	import InputDateElement from '$lib/components/form-input/InputDateElement.svelte';
	import SearchInputElement from '$lib/components/form-input/SearchInputElement.svelte';
	import SelectOptionsElement from '$lib/components/form-input/SelectOptionsElement.svelte';

	export let data;

	$: events = data.events;
	$: pagination = data.pagination;
	let currentParams = data.params;

	// Filter state
	let searchQuery = currentParams.search || '';
	let endDateFilter = currentParams.endDate || '';
	let locationFilter = currentParams.location || '';
	let startDateFilter = currentParams.startDate || '';
	let perPage = currentParams.limit || 5;
	let currentPage = currentParams.page || 1;

	// Show/hide filters
	let showFilters = false;

	// Loading state
	let loading = false;

	// Update URL with current filters
	function updateURL() {
		const params = new URLSearchParams();

		if (searchQuery) params.set('search', searchQuery);
		if (endDateFilter) params.set('endDate', endDateFilter);
		if (locationFilter) params.set('location', locationFilter);
		if (perPage !== 10) params.set('limit', perPage.toString());
		if (startDateFilter) params.set('startDate', startDateFilter);
		if (currentPage > 1) params.set('page', currentPage.toString());

		const newURL = params.toString() ? `?${params.toString()}` : '/events';
		goto(newURL, { replaceState: true });
	}

	// Reset all filters
	function resetFilters() {
		searchQuery = '';
		startDateFilter = '';
		endDateFilter = '';
		locationFilter = '';
		currentPage = 1;
		perPage = 10;
		updateURL();
	}

	// Handle search
	function handleSearch() {
		currentPage = 1;
		updateURL();
	}

	// Handle filter change
	function handleFilterChange() {
		currentPage = 1;
		updateURL();
	}

	// Handle pagination
	function onPageChange(page: number) {
		currentPage = page;
		updateURL();
	}

	// Check if any filters are active
	$: hasActiveFilters = searchQuery || startDateFilter || endDateFilter || locationFilter;

	// Reactive statements for handle date dependencies
	$: if (startDateFilter && endDateFilter && endDateFilter < startDateFilter) {
		endDateFilter = '';
		handleFilterChange();
	}

	// Watch for date filter changes and trigger URL update
	$: if (startDateFilter || endDateFilter) {
		handleFilterChange();
	}
</script>

<svelte:head>
	<title>Tiketku - Event Discovery</title>
	<meta
		name="description"
		content="Discover and book amazing events in Indonesia. Find concerts, festivals, workshops, and more!"
	/>
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<meta property="og:title" content="Tiketku - Book Amazing Events" />
	<meta
		property="og:description"
		content="Book concerts, festivals, workshops & sports events in Indonesia"
	/>
	<meta property="og:image" content="/og-tiketku.jpg" />
</svelte:head>

<div class="min-h-screen bg-gray-50">
	<div class="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
		<!-- Header Section -->
		<div class="mb-8 text-center">
			<h1
				class="mb-4 bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-4xl font-bold text-transparent"
			>
				Discover Events
			</h1>
			<p class="text-lg text-gray-600">Find and join amazing events happening around you</p>
		</div>

		<!-- Search and Filters -->
		<div class="mb-8 rounded-xl border border-gray-200 bg-white p-6 shadow-sm">
			<!-- Search Bar -->
			<div class="mb-4 flex flex-col gap-4 sm:flex-row">
				<div class="flex-1">
					<SearchInputElement
						bind:value={searchQuery}
						onKeyDown={(e) => {
							if (e.key === 'Enter') {
								handleSearch();
							}
						}}
						placeholder="Search events by title, description..."
					/>
				</div>
				<div class="flex gap-3">
					<Button variant="event" size="xl" onclick={handleSearch}>Search</Button>
					<Button onclick={() => (showFilters = !showFilters)} variant="event-secondary" size="xl">
						<Filter class="h-4 w-4" />
						Filters
						{#if hasActiveFilters}
							<span class="ml-1 rounded-full bg-blue-500 px-2 py-0.5 text-xs text-white">
								{[searchQuery, startDateFilter, endDateFilter, locationFilter].filter(Boolean)
									.length}
							</span>
						{/if}
					</Button>
				</div>
			</div>

			<!-- Advanced Filters -->
			{#if showFilters}
				<div class="pt-6">
					<div class="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3">
						<!-- Start Date Filter -->
						<InputDateElement
							label="Start Date"
							id="start-date-filter"
							bind:value={startDateFilter}
							placeholder="Select start date"
						/>

						<!-- End Date Filter  -->
						<InputDateElement
							label="End Date"
							isEndDate={true}
							id="end-date-filter"
							bind:value={endDateFilter}
							startDate={startDateFilter}
							placeholder="Select end date"
						/>

						<!-- Location Filter -->
						<SelectOptionsElement
							id="location-filter"
							clearable={true}
							label="Location Filter"
							options={locationOptions}
							bind:value={locationFilter}
							onchange={handleFilterChange}
							placeholder="Select location"
						/>
					</div>

					<!-- Active Filters Display -->
					{#if hasActiveFilters}
						<div class="mt-4 mb-4">
							<h4 class="mb-2 text-sm font-medium text-gray-700">Active Filters:</h4>
							<div class="flex flex-wrap gap-2">
								{#if searchQuery}
									<span
										class="inline-flex items-center gap-1 rounded-full bg-blue-100 px-3 py-1 text-xs text-blue-700"
									>
										Search: {searchQuery}
										<button
											on:click={() => {
												searchQuery = '';
												handleFilterChange();
											}}
											class="hover:text-blue-900"
										>
											×
										</button>
									</span>
								{/if}
								{#if startDateFilter}
									<span
										class="inline-flex items-center gap-1 rounded-full bg-blue-100 px-3 py-1 text-xs text-blue-700"
									>
										From: {formatDate(new Date(startDateFilter + 'T00:00:00'))}
										<button
											on:click={() => {
												startDateFilter = '';
											}}
											class="hover:text-blue-900"
										>
											×
										</button>
									</span>
								{/if}
								{#if endDateFilter}
									<span
										class="inline-flex items-center gap-1 rounded-full bg-green-100 px-3 py-1 text-xs text-green-700"
									>
										To: {formatDate(new Date(endDateFilter + 'T00:00:00'))}
										<button
											on:click={() => {
												endDateFilter = '';
											}}
											class="hover:text-green-900"
										>
											×
										</button>
									</span>
								{/if}
								{#if locationFilter}
									<span
										class="inline-flex items-center gap-1 rounded-full bg-indigo-100 px-3 py-1 text-xs text-indigo-700"
									>
										Location: {locationOptions.find((opt) => opt.value === locationFilter)?.label}
										<button
											on:click={() => {
												locationFilter = '';
												handleFilterChange();
											}}
											class="hover:text-indigo-900"
										>
											×
										</button>
									</span>
								{/if}
							</div>
						</div>
					{/if}

					<div class="mt-6 flex items-center justify-between">
						<div class="text-sm text-gray-500">
							{#if hasActiveFilters}
								Showing filtered results
							{:else}
								Showing all events
							{/if}
						</div>
						{#if hasActiveFilters}
							<Button onclick={resetFilters} variant="outline" class="w-40">
								<FilterX class="h-4 w-4" />
								Reset All Filters
							</Button>
						{/if}
					</div>
				</div>
			{/if}
		</div>

		<!-- Loading State -->
		{#if loading}
			<EventsLoading />
		{:else if events.length === 0}
			<!-- Empty State -->
			<EventNotFound {resetFilters} />
		{:else}
			<!-- Events Grid -->
			<div class="mb-8 grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
				{#each events as event}
					<EventCard data={event} />
				{/each}
			</div>

			<!-- Pagination -->
			<div class="rounded-xl shadow-sm">
				<Pagination
					{perPage}
					{currentPage}
					{onPageChange}
					total={pagination.total}
					totalPages={pagination.totalPages}
				/>
			</div>
		{/if}
	</div>
</div>
