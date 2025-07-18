<script>
	import { formatDate } from '$lib/utils/formatter';
	import { MapPin, Calendar, Users } from '@lucide/svelte';
	import Button from '$lib/components/ui/button/button.svelte';

	export let data;
</script>

<div
	class="group overflow-hidden rounded-xl border border-gray-200 bg-white shadow-sm transition-all hover:-translate-y-1 hover:shadow-lg hover:shadow-blue-100/50"
>
	<!-- Event Image  -->
	{#if data.image}
		<div class="aspect-video overflow-hidden">
			<img
				src={data.image}
				alt={data.title}
				class="h-full w-full object-cover transition-transform group-hover:scale-105"
			/>
		</div>
	{/if}

	<div class="p-6">
		<!-- Event Title -->
		<h3
			class="mb-3 line-clamp-1 text-xl font-semibold text-gray-800 transition-colors group-hover:text-blue-600"
		>
			{data.title}
		</h3>

		<!-- Event Description -->
		<p class="mb-4 line-clamp-3 leading-relaxed text-gray-600">
			{data.description}
		</p>

		<!-- Event Details -->
		<div class="mb-4 space-y-2">
			{#if data.date}
				<div class="flex items-center text-sm">
					<Calendar class="mr-2 h-4 w-4 text-blue-500" />
					{formatDate(data.date)}
				</div>
			{/if}

			{#if data.location}
				<div class="flex items-center text-sm text-gray-500">
					<MapPin class="mr-2 h-4 w-4 text-indigo-500" />
					{data.location}
				</div>
			{/if}

			{#if data.maxCapacity}
				<div class="flex items-center text-sm text-gray-500">
					<Users class="mr-2 h-4 w-4 text-blue-500" />
					{data.currentAttendees || 0} / {data.maxCapacity} participants
				</div>
			{/if}
		</div>

		<!-- Action Button -->
		<div class="flex items-center justify-between">
			<Button
				href={`/events/${data.id}`}
				variant="event"
				class="w-full"
				disabled={data.status === 'full'}
			>
				{data.status === 'full' ? 'Event Full' : 'Get Ticket'}
			</Button>

			{#if data.price}
				<div class="ml-3 text-right">
					<span class="text-lg font-bold text-blue-600">
						{data.price === 0 ? 'Free' : `$${data.price}`}
					</span>
				</div>
			{/if}
		</div>
	</div>
</div>
