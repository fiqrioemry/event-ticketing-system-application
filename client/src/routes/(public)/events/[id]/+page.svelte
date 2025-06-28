<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import { formatDate } from '$lib/utils';
	import { CalendarDays, MapPin, User } from '@lucide/svelte';

	const event = {
		id: '6e890af2-2f0f-41d2-bcc3-d9450fcd73f1',
		title: 'Konser Musik Senja',
		image: 'https://placehold.co/400x400?text=Konser+Musik+Senja',
		description: 'Nikmati malam dengan alunan musik akustik dari musisi ternama Indonesia.',
		location: 'Jakarta, Indonesia',
		date: '2025-07-12T19:00:00Z',
		status: 'active',
		created_by: {
			id: 'user-123',
			fullname: 'Ahmad Fiqri',
			avatar_url: 'https://placehold.co/100x100?text=User',
			role: 'admin'
		},
		categories: [
			{
				id: 'cat-vip',
				name: 'VIP',
				price: 500000,
				quota: 50,
				booked: 20,
				available: 30
			},
			{
				id: 'cat-reg',
				name: 'Reguler',
				price: 250000,
				quota: 100,
				booked: 60,
				available: 40
			}
		],
		report: {
			total_sold: 80,
			total_revenue: 29000000,
			generated_at: '2025-06-27T00:00:00Z'
		},
		created_at: '2025-05-01T10:00:00Z',
		updated_at: '2025-06-01T12:00:00Z'
	};
</script>

<svelte:head>
	<title>Events Detail | MyBrand</title>
	<meta
		name="description"
		content="display the details of the event, including categories, report, and more."
	/>
</svelte:head>

<section class="mx-auto max-w-5xl px-4 py-10">
	<!-- HEADER -->
	<img src={event.image} alt={event.title} class="mb-6 h-64 w-full rounded-xl object-cover" />

	<div class="mb-4 flex flex-col gap-2">
		<h1 class="text-3xl font-bold">{event.title}</h1>
		<p class="text-muted-foreground">{event.description}</p>

		<div class="text-muted-foreground mt-2 flex items-center gap-4 text-sm">
			<span class="flex items-center gap-1">
				<CalendarDays class="h-4 w-4" />

				{formatDate(event.date)}
			</span>
			<span class="flex items-center gap-1">
				<MapPin class="h-4 w-4" />
				{event.location}
			</span>
			<span class="flex items-center gap-1">
				<User class="h-4 w-4" />
				{event.created_by.fullname}
			</span>
		</div>
	</div>

	<!-- TICKET CATEGORIES -->
	<div class="mt-6 border-t pt-6">
		<h2 class="mb-4 text-xl font-semibold">🎟️ Tiket Tersedia</h2>

		<div class="grid gap-4 sm:grid-cols-2">
			{#each event.categories as cat}
				<div class="rounded-xl border bg-white p-4 dark:bg-zinc-900">
					<h3 class="text-lg font-semibold">{cat.name}</h3>
					<p class="text-muted-foreground text-sm">Rp {cat.price.toLocaleString('id-ID')}</p>

					<div class="text-muted-foreground mt-2 text-sm">
						Kuota: {cat.quota} <br />
						Booked: {cat.booked} <br />
						Tersisa: <strong>{cat.available}</strong>
					</div>

					<Button class="mt-4 w-full" disabled={cat.available === 0}>Beli Tiket</Button>
				</div>
			{/each}
		</div>
	</div>

	<!-- REPORT (optional admin or stats view) -->
	<div class="mt-10">
		<h2 class="mb-2 text-xl font-semibold">📊 Statistik Penjualan</h2>
		<div class="text-muted-foreground text-sm">
			Total Terjual: {event.report.total_sold} tiket <br />
			Total Pendapatan: Rp {event.report.total_revenue.toLocaleString('id-ID')} <br />
			Update Terakhir: {formatDate(event.report.generated_at)}
		</div>
	</div>
</section>
