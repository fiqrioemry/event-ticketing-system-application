<script lang="ts">
	const orderDetail = [
		{
			id: '41a940d8-933c-400b-8091-6cfc759e686c',
			ticketName: 'Gold',
			ticketId: 'd39ff313-db85-4a3f-9f33-fa8fb0c68019',
			quantity: 1,
			price: 120000,
			createdAt: '2025-06-28'
		},
		{
			id: '8b6b8df9-cbf9-47e0-8349-e6000b366731',
			ticketName: 'Silver',
			ticketId: '802a2b63-d941-4557-8dc7-0d0580d0580f',
			quantity: 2,
			price: 90000,
			createdAt: '2025-06-28'
		}
	];

	function formatCurrency(value: number) {
		return `Rp ${value.toLocaleString('id-ID')}`;
	}

	const total = orderDetail.reduce((sum, item) => sum + item.price * item.quantity, 0);
</script>

<svelte:head>
	<title>Purchase Detail | MyBrand</title>
	<meta name="description" content="See your ticket from your purchase detail" />
</svelte:head>

<section class="mx-auto max-w-3xl p-6">
	<h1 class="mb-2 text-2xl font-bold">🧾 Purchase Detail</h1>
	<p class="text-muted-foreground mb-6">Rincian tiket dari pesanan kamu</p>

	<div class="space-y-4">
		{#each orderDetail as item (item.id)}
			<div class="bg-muted/50 flex items-center justify-between rounded-lg border p-4 shadow-sm">
				<div>
					<p class="text-lg font-semibold">{item.ticketName}</p>
					<p class="text-muted-foreground text-sm">
						Qty: {item.quantity} × {formatCurrency(item.price)}
					</p>
					<p class="text-muted-foreground text-sm">ID: {item.ticketId.slice(0, 8)}...</p>
				</div>
				<div class="text-right font-semibold">
					{formatCurrency(item.price * item.quantity)}
				</div>
			</div>
		{/each}
	</div>

	<hr class="border-border my-6" />

	<div class="flex items-center justify-between">
		<p class="text-lg font-semibold">Total Pembayaran</p>
		<p class="text-primary text-xl font-bold">{formatCurrency(total)}</p>
	</div>
</section>
