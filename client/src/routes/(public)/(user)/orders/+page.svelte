<script lang="ts">
	import { get } from 'svelte/store';
	import { onMount } from 'svelte';
	import { TicketIcon } from '@lucide/svelte';
	import EyeIcon from '@lucide/svelte/icons/eye';
	import { useGetOrders } from '$lib/hooks/useOrder';
	import { formatDate, formatRupiah } from '$lib/utils.js';
	import * as Table from '$lib/components/ui/table/index.js';
	import Button from '$lib/components/ui/button/button.svelte';
	import Pagination from '$lib/components/common/Pagination.svelte';
	import { orderFilters, ordersResponse, isLoading } from '$lib/stores/order.store.js';
	import Input from '$lib/components/ui/input/input.svelte';
	import NoResult from '$lib/components/common/NoSearchResult.svelte';
	import OrdersListLoading from '$lib/components/loading/OrdersListLoading.svelte';
	import NoSearchResult from '$lib/components/common/NoSearchResult.svelte';

	$: orders = $ordersResponse.orders || [];
	$: pagination = $ordersResponse.pagination || {};

	let timeout: ReturnType<typeof setTimeout>;

	function handleSearchInput(val: string) {
		clearTimeout(timeout);
		orderFilters.update((f) => ({ ...f, search: val, page: 1 }));
		timeout = setTimeout(() => {
			useGetOrders(get(orderFilters));
		}, 300);
	}
	function handlerOrderChange(newPage: number) {
		orderFilters.set({ ...get(orderFilters), page: newPage });
		useGetOrders(get(orderFilters));
	}

	onMount(() => {
		useGetOrders(get(orderFilters));
	});
</script>

<svelte:head>
	<title>Purchase | Tiketku</title>
	<meta name="description" content="See your purchase history with Tiketku." />
</svelte:head>

<section class="bg-background mx-auto max-w-5xl rounded-lg border p-6">
	<h1 class="mb-6 text-center text-2xl font-bold">Purchasement History</h1>
	<Input
		bind:value={$orderFilters.search}
		oninput={(e) => handleSearchInput((e.target as HTMLInputElement).value)}
		class="mb-4 w-full md:w-1/2"
		placeholder="Search by event name or email"
	/>

	{#if $isLoading}
		<OrdersListLoading />
	{:else if orders.length === 0}
		<NoSearchResult title="Pesanan tidak ditemukan" description="Coba ubah kata kunci pencarian." />
	{:else}
		<Table.Root>
			<Table.Header>
				<Table.Row>
					<Table.Head class="min-w-[200px]">Event</Table.Head>
					<Table.Head>Status</Table.Head>
					<Table.Head>Total Payment</Table.Head>
					<Table.Head>Date</Table.Head>
					<Table.Head class="text-center">Detail</Table.Head>
				</Table.Row>
			</Table.Header>
			<Table.Body>
				{#each orders as p (p.id)}
					<Table.Row class="hover:bg-muted/50">
						<Table.Cell>
							<div class="flex items-center gap-3">
								<img
									src={p.eventImage}
									alt="this_event"
									class="h-10 w-10 rounded border object-cover"
								/>
								<div>
									<div class="font-medium">{p.eventName}</div>
									<div class="text-muted-foreground text-xs">{p.email}</div>
								</div>
							</div>
						</Table.Cell>
						<Table.Cell>
							<span
								class={`font-medium capitalize ${p.status === 'paid' ? 'text-green-600' : p.status === 'failed' ? 'text-red-500' : 'text-yellow-600'}`}
							>
								{p.status}
							</span>
						</Table.Cell>
						<Table.Cell>{formatRupiah(p.totalPrice)}</Table.Cell>
						<Table.Cell>{formatDate(p.createdAt)}</Table.Cell>
						<Table.Cell class="space-x-2 text-center">
							{#if p.status === 'paid'}
								<Button variant="outline" size="icon" href={`/orders/${p.id}/tickets`}>
									<TicketIcon />
								</Button>
							{/if}
							<Button variant="outline" size="icon" href={`/orders/${p.id}`}>
								<EyeIcon class="size-4" />
							</Button>
						</Table.Cell>
					</Table.Row>
				{/each}
			</Table.Body>
		</Table.Root>
	{/if}

	<!-- Pagination -->
	{#if !$isLoading && pagination.totalPages > 1}
		<Pagination
			page={pagination.page}
			totalPages={pagination.totalPages}
			on:change={(e) => handlerOrderChange(e.detail)}
		/>
	{/if}
</section>
