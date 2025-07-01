<script lang="ts">
	import qs from 'qs';
	import {
		isLoading,
		eventFilters,
		eventSortOptions,
		eventResponse,
		eventLocationOptions,
		eventStatusOptions
	} from '$lib/stores/event.store';
	import { goto } from '$app/navigation';
	import { formatDate } from '$lib/utils';
	import Input from '$lib/components/ui/input/input.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import NoSearchResult from '$lib/components/common/NoResult.svelte';
	import { ArrowRight, CalendarDays, MapPin } from '@lucide/svelte';
	import FilterSelect from '$lib/components/common/FilterSelect.svelte';
	import Pagination from '$lib/components/common/Pagination.svelte';
	import EventListsLoading from '$lib/components/loading/EventListsLoading.svelte';

	$: filters = $eventFilters;
	$: loading = $isLoading;
	$: events = $eventResponse.events;
	$: pagination = $eventResponse.pagination;

	function fetchEvents(newPage?: number) {
		const merged = { ...filters };

		if (newPage !== undefined) {
			merged.page = newPage;
		}

		const query = qs.stringify(merged, { skipNulls: true });
		isLoading.set(true);
		goto(`/events?${query}`);
	}
</script>

<section class="mx-auto max-w-7xl px-4 py-10">
	<h1 class="mb-6 text-2xl font-bold">Daftar Event</h1>
	<div class="mb-6 flex flex-col items-center gap-4 md:flex-row">
		<div class="flex w-full items-center md:w-1/2">
			<form on:submit={() => fetchEvents(1)} class=" flex w-full gap-4">
				<Input type="text" name="search" bind:value={filters.search} placeholder="Search" />
				<Button type="submit">Search</Button>
			</form>
		</div>

		<div class="flex w-full items-center gap-4 md:w-1/2">
			<FilterSelect
				value={filters.location}
				options={eventLocationOptions}
				on:change={(e) => {
					filters.location = e.detail;
					fetchEvents(1);
				}}
			/>

			<FilterSelect
				value={filters.status}
				options={eventStatusOptions}
				on:change={(e) => {
					filters.status = e.detail;
					fetchEvents(1);
				}}
			/>

			<FilterSelect
				value={filters.sort}
				options={eventSortOptions}
				on:change={(e) => {
					filters.sort = e.detail;
					fetchEvents(1);
				}}
			/>
		</div>
	</div>
	{#if loading}
		<EventListsLoading />
	{:else if events && events.length > 0}
		<div class="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
			{#each events as event}
				<div class="overflow-hidden rounded-2xl bg-white shadow hover:shadow-lg dark:bg-zinc-900">
					<img src={event.image} alt={event.title} class="h-48 w-full object-cover" />
					<div class="flex flex-col gap-2 p-4">
						<h2 class="line-clamp-2 text-lg font-semibold">{event.title}</h2>
						<p class="text-muted-foreground line-clamp-3 text-sm">{event.description}</p>
						<div class="text-muted-foreground mt-2 flex items-center gap-2 text-sm">
							<CalendarDays class="h-4 w-4" />
							<span>{formatDate(event.date)}</span>
						</div>
						<div class="text-muted-foreground flex items-center gap-2 text-sm">
							<MapPin class="h-4 w-4" />
							<span>{event.location}</span>
						</div>
						<div class="mt-4">
							<Button href={`/events/${event.id}`} class="w-full justify-between" variant="outline">
								Lihat Detail <ArrowRight class="ml-2 h-4 w-4" />
							</Button>
						</div>
					</div>
				</div>
			{/each}
		</div>
	{:else}
		<NoSearchResult />
	{/if}

	<!-- Pagination tombol -->

	{#if pagination}
		<Pagination
			page={pagination.page}
			totalPages={pagination.totalPages}
			on:change={(e) => fetchEvents(e.detail)}
		/>
	{/if}
</section>
