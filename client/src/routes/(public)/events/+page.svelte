<script lang="ts">
	import { formatDate } from '$lib/utils.js';
	import { Input } from '$lib/components/ui/input';
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Select from '$lib/components/ui/select/index.js';
	import * as Pagination from '$lib/components/ui/pagination/index.js';
	import { ArrowRight, CalendarDays, MapPin, Search } from '@lucide/svelte';

	let search = '';
	let status = '';

	const data = {
		page: 1,
		total: 10,
		events: [
			{
				id: '6e890af2-2f0f-41d2-bcc3-d9450fcd73f1',
				title: 'Konser Musik Senja',
				image: 'https://placehold.co/400x400?text=Konser+Musik+Senja',
				description: 'Nikmati malam dengan alunan musik akustik dari musisi ternama Indonesia.',
				location: 'Jakarta, Indonesia',
				date: '2025-07-12T19:00:00Z',
				status: 'active'
			},
			{
				id: 'b7de28d4-1b3d-4a8d-bf69-03b1715c23a4',
				title: 'Tech Conference 2025',
				image: 'https://placehold.co/400x400?text=Tech+Conference',
				description:
					'Konferensi teknologi terbesar tahun ini dengan pembicara dari Google, AWS, dan Microsoft.',
				location: 'Bandung Convention Center',
				date: '2025-08-02T09:00:00Z',
				status: 'active'
			},
			{
				id: '3f79e3bb-2b98-429e-b1e0-b861f9fcdbbe',
				title: 'Workshop UI/UX Design',
				image: 'https://placehold.co/400x400?text=UI%2FUX+Workshop',
				description:
					'Pelatihan intensif untuk belajar membuat desain digital yang menarik dan efisien.',
				location: 'Yogyakarta, Indonesia',
				date: '2025-07-28T13:00:00Z',
				status: 'active'
			},
			{
				id: 'df347ec2-58a1-4b8a-b16d-4829ef7ffaaa',
				title: 'Festival Kuliner Nusantara',
				image: 'https://placehold.co/400x400?text=Festival+Kuliner',
				description: 'Rasakan cita rasa dari Sabang sampai Merauke dalam satu tempat!',
				location: 'Surabaya Food Park',
				date: '2025-09-10T10:00:00Z',
				status: 'active'
			},
			{
				id: 'fa125d23-dccc-4e70-b00c-3b0c7255a3a2',
				title: 'Seminar Finansial Milenial',
				image: 'https://placehold.co/400x400?text=Seminar+Finansial',
				description: 'Pelajari cara mengatur keuangan dan berinvestasi di usia muda.',
				location: 'Online (Zoom Webinar)',
				date: '2025-07-20T18:00:00Z',
				status: 'active'
			}
		]
	};
	function handleSearch() {
		// Implement logic pencarian jika pakai query string
		console.log('search:', search, 'status:', status);
	}
</script>

<svelte:head>
	<title>Events | MyBrand</title>
	<meta
		name="description"
		content="See all the events available and purchase tickets for your favorite events with MyBrand."
	/>
</svelte:head>

<section class="mx-auto max-w-7xl px-4 py-10">
	<h1 class="mb-6 text-3xl font-bold">🎉 Daftar Event Populer</h1>

	<!-- Filter & Search -->
	<div class="mb-6 flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
		<div class="flex gap-2">
			<Input
				class="w-64"
				placeholder="Cari judul event..."
				bind:value={search}
				onkeydown={(e) => e.key === 'Enter' && handleSearch()}
			/>

			<Select.Root bind:value={status} type="single">
				<Select.Trigger class="w-[180px]">Filter Status</Select.Trigger>
				<Select.Content>
					<Select.Item value="active">Active</Select.Item>
					<Select.Item value="done">Done</Select.Item>
					<Select.Item value="cancelled">Cancelled</Select.Item>
				</Select.Content>
			</Select.Root>
		</div>
	</div>

	<!-- Event Grid -->
	<div class="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
		{#each data.events as event}
			<div
				class="overflow-hidden rounded-2xl bg-white shadow transition hover:shadow-lg dark:bg-zinc-900"
			>
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
						<Button class="w-full justify-between" href={`/events/${event.id}`} variant="outline">
							Lihat Detail <ArrowRight class="ml-2 h-4 w-4" />
						</Button>
					</div>
				</div>
			</div>
		{/each}
	</div>

	<!-- Pagination -->
	<div class="mt-10 flex justify-center">
		<Pagination.Root count={data.events.length} perPage={data.total}>
			{#snippet children({ pages, currentPage })}
				<Pagination.Content>
					<Pagination.Item>
						<Pagination.PrevButton />
					</Pagination.Item>

					{#each pages as page (page.key)}
						{#if page.type === 'ellipsis'}
							<Pagination.Item>
								<Pagination.Ellipsis />
							</Pagination.Item>
						{:else}
							<Pagination.Item>
								<Pagination.Link {page} isActive={currentPage === page.value}>
									{page.value}
								</Pagination.Link>
							</Pagination.Item>
						{/if}
					{/each}

					<Pagination.Item>
						<Pagination.NextButton />
					</Pagination.Item>
				</Pagination.Content>
			{/snippet}
		</Pagination.Root>
	</div>
</section>
