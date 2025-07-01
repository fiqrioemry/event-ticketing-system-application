<script lang="ts">
	import { formatRupiah } from '$lib/utils';
	export let data;
	const orderDetail = data.order;
</script>

<svelte:head>
	<title>Purchase Detail | MyBrand</title>
	<meta name="description" content="See your ticket from your purchase detail" />
</svelte:head>

<section class="mx-auto max-w-3xl px-6 py-10">
	<div class="mb-10 text-center">
		<h1 class="text-2xl font-bold tracking-tight">🧾 Purchase Detail</h1>
		<p class="text-muted-foreground mt-2 text-sm">Rincian tiket dari pesanan kamu</p>
	</div>

	<div class="space-y-5">
		{#each orderDetail as item (item.id)}
			<div
				class="rounded-xl border bg-white/70 p-5 shadow-sm transition hover:shadow-md dark:bg-zinc-900/60"
			>
				<div class="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
					<div>
						<p class="text-base font-semibold sm:text-lg">{item.ticketName}</p>
						<p class="text-muted-foreground text-sm">
							{item.quantity} × {formatRupiah(item.price)}
						</p>
					</div>
					<div class="text-primary text-right text-sm font-semibold sm:text-base">
						{formatRupiah(item.price * item.quantity)}
					</div>
				</div>
			</div>
		{/each}
	</div>
</section>
