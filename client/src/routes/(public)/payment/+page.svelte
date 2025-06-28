<script lang="ts">
	import { formatDate } from '$lib/utils';
	import Button from '$lib/components/ui/button/button.svelte';
	import * as RadioGroup from '$lib/components/ui/radio-group/index.js';
	import { CheckCircle2, CalendarDays, MapPin } from '@lucide/svelte';

	let event = {
		title: 'Konser Musik Senja',
		location: 'Jakarta, Indonesia',
		date: '2025-07-12T19:00:00Z',
		image: 'https://placehold.co/600x300?text=Konser'
	};

	let selectedTickets = [
		{ category: 'VIP', price: 500000, quantity: 2 },
		{ category: 'Reguler', price: 250000, quantity: 1 }
	];

	let paymentMethod = 'credit_card';

	// Hitung total
	$: total = selectedTickets.reduce((acc, t) => acc + t.price * t.quantity, 0);
</script>

<svelte:head>
	<title>Payment | MyBrand</title>
	<meta
		name="description"
		content="Manage your payments effectively with MyBrand's payment dashboard."
	/>
</svelte:head>

<section class="mx-auto max-w-3xl px-4 py-10">
	<h1 class="mb-6 text-2xl font-bold">💳 Pembayaran Tiket</h1>

	<!-- Ringkasan Event -->
	<div class="mb-6 rounded-xl border bg-white p-4 dark:bg-zinc-900">
		<div class="flex gap-4">
			<img src={event.image} alt="event" class="h-28 w-40 rounded object-cover" />
			<div class="flex flex-col justify-between">
				<div>
					<h2 class="text-lg font-semibold">{event.title}</h2>
					<p class="text-muted-foreground flex items-center gap-1 text-sm">
						<CalendarDays class="h-4 w-4" />
						{formatDate(event.date)}
					</p>
					<p class="text-muted-foreground flex items-center gap-1 text-sm">
						<MapPin class="h-4 w-4" />
						{event.location}
					</p>
				</div>
			</div>
		</div>
	</div>

	<!-- Tiket yang Dibeli -->
	<div class="mb-6">
		<h2 class="mb-2 text-lg font-semibold">🎟️ Tiket yang Dibeli</h2>
		<ul class="space-y-2">
			{#each selectedTickets as t}
				<li class="flex justify-between border-b pb-2">
					<span>{t.quantity}x {t.category}</span>
					<span>Rp {(t.price * t.quantity).toLocaleString('id-ID')}</span>
				</li>
			{/each}
			<li class="flex justify-between pt-2 font-bold">
				<span>Total</span>
				<span>Rp {total.toLocaleString('id-ID')}</span>
			</li>
		</ul>
	</div>

	<!-- Pilih Metode Pembayaran -->
	<div class="mb-6">
		<h2 class="mb-2 text-lg font-semibold">💼 Metode Pembayaran</h2>
		<RadioGroup.Root bind:value={paymentMethod}>
			<div class="flex flex-col gap-2">
				<label class="flex items-center gap-2">
					<RadioGroup.Item value="credit_card" />
					<span>Kartu Kredit / Debit</span>
				</label>
				<label class="flex items-center gap-2">
					<RadioGroup.Item value="bank_transfer" />
					<span>Transfer Bank</span>
				</label>
				<label class="flex items-center gap-2">
					<RadioGroup.Item value="ewallet" />
					<span>E-Wallet (Gopay, Dana, OVO)</span>
				</label>
			</div>
		</RadioGroup.Root>
	</div>

	<!-- Tombol Bayar -->
	<Button class="mt-4 w-full" size="lg">
		<CheckCircle2 class="mr-2 h-5 w-5" />
		Bayar Sekarang (Rp {total.toLocaleString('id-ID')})
	</Button>
</section>
