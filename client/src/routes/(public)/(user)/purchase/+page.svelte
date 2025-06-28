<script lang="ts">
	import * as Table from '$lib/components/ui/table/index.js';
	import * as Pagination from '$lib/components/ui/pagination/index.js';
	import ChevronLeftIcon from '@lucide/svelte/icons/chevron-left';
	import ChevronRightIcon from '@lucide/svelte/icons/chevron-right';
	import { MediaQuery } from 'svelte/reactivity';
	import EyeIcon from '@lucide/svelte/icons/eye';
	import Button from '$lib/components/ui/button/button.svelte';
	import { goto } from '$app/navigation';

	const isDesktop = new MediaQuery('(min-width: 768px)');
	const perPage = $derived(isDesktop.current ? 3 : 8);
	const siblingCount = $derived(isDesktop.current ? 1 : 0);

	const purchases = {
		data: [
			{
				id: '2b9ad08e-c087-4702-8b18-6b6c655ddabf',
				eventId: 'ddaf8eb0-a68e-4316-8dc7-834d183faaf6',
				eventName: 'Event 3',
				eventImage: 'https://placehold.co/400x400?text=Acoustic+Night',
				fullname: 'John Doe',
				email: 'john@example.com',
				phone: '081234567890',
				totalPrice: 300000,
				status: 'failed',
				createdAt: '2025-06-28'
			},
			{
				id: '93263d98-fce1-4297-988b-3f2b4594653c',
				eventId: 'ddaf8eb0-a68e-4316-8dc7-834d183faaf6',
				eventName: 'Event 3',
				eventImage: 'https://placehold.co/400x400?text=Acoustic+Night',
				fullname: 'John Doe',
				email: 'john@example.com',
				phone: '081234567890',
				totalPrice: 300000,
				status: 'paid',
				createdAt: '2025-06-28'
			},
			{
				id: 'ab2cdf9c-2803-4be8-9800-c5644a9124b1',
				eventId: 'ddaf8eb0-a68e-4316-8dc7-834d183faaf6',
				eventName: 'Event 3',
				eventImage: 'https://placehold.co/400x400?text=Acoustic+Night',
				fullname: 'John Doe',
				email: 'john@example.com',
				phone: '081234567890',
				totalPrice: 300000,
				status: 'failed',
				createdAt: '2025-06-28'
			},
			{
				id: '7211b4b2-bd46-48a5-83f4-924e246c917b',
				eventId: 'ddaf8eb0-a68e-4316-8dc7-834d183faaf6',
				eventName: 'Event 3',
				eventImage: 'https://placehold.co/400x400?text=Acoustic+Night',
				fullname: 'John Doe',
				email: 'john@example.com',
				phone: '081234567890',
				totalPrice: 300000,
				status: 'refunded',
				createdAt: '2025-06-28'
			},
			{
				id: '24525c25-1944-4d57-8bbd-c4169cc46f04',
				eventId: 'ddaf8eb0-a68e-4316-8dc7-834d183faaf6',
				eventName: 'Event 3',
				eventImage: 'https://placehold.co/400x400?text=Acoustic+Night',
				fullname: 'Charlie Customer',
				email: 'charlie@event.com',
				phone: '08123456788',
				totalPrice: 210000,
				status: 'paid',
				createdAt: '2025-06-28'
			}
		],
		pagination: {
			page: 1,
			limit: 10,
			totalRows: 5,
			totalPages: 1
		}
	};

	function formatDate(dateStr: string) {
		return new Date(dateStr).toLocaleDateString('id-ID', {
			day: '2-digit',
			month: 'short',
			year: 'numeric'
		});
	}

	function formatCurrency(amount: number) {
		return `Rp ${amount.toLocaleString('id-ID')}`;
	}
</script>

<svelte:head>
	<title>Purchase | MyBrand</title>
	<meta name="description" content="See your purchase history with MyBrand." />
</svelte:head>

<section class="bg-background mx-auto max-w-5xl rounded-lg border p-6">
	<h1 class="mb-6 text-center text-2xl font-bold">🧾 Riwayat Pembelian Tiket</h1>

	<Table.Root>
		<Table.Caption class="text-muted-foreground mb-2"
			>Daftar lengkap pembelian tiket kamu</Table.Caption
		>
		<Table.Header>
			<Table.Row>
				<Table.Head class="min-w-[200px]">Event</Table.Head>
				<Table.Head>Status</Table.Head>
				<Table.Head class="text-right">Total Bayar</Table.Head>
				<Table.Head class="text-right">Tanggal</Table.Head>
				<Table.Head class="text-right">Detail</Table.Head>
			</Table.Row>
		</Table.Header>
		<Table.Body>
			{#each purchases.data as p (p.id)}
				<Table.Row class="hover:bg-muted/50">
					<Table.Cell>
						<div class="flex items-center gap-3">
							<img
								src={p.eventImage}
								alt="event image"
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
					<Table.Cell class="text-right">{formatCurrency(p.totalPrice)}</Table.Cell>
					<Table.Cell class="text-right">{formatDate(p.createdAt)}</Table.Cell>
					<Table.Cell class="text-right">
						<Button variant="outline" size="icon" href={`/purchase/${p.id}`}>
							<EyeIcon class="size-4" />
						</Button>
					</Table.Cell>
				</Table.Row>
			{/each}
		</Table.Body>
	</Table.Root>

	<!-- Pagination -->
	<div class="mt-6">
		<Pagination.Root
			count={purchases.pagination.totalRows}
			{perPage}
			{siblingCount}
			page={purchases.pagination.page}
		>
			{#snippet children({ pages, currentPage })}
				<Pagination.Content>
					<Pagination.Item>
						<Pagination.PrevButton>
							<ChevronLeftIcon class="size-4" />
							<span class="hidden sm:block">Previous</span>
						</Pagination.PrevButton>
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
						<Pagination.NextButton>
							<span class="hidden sm:block">Next</span>
							<ChevronRightIcon class="size-4" />
						</Pagination.NextButton>
					</Pagination.Item>
				</Pagination.Content>
			{/snippet}
		</Pagination.Root>
	</div>
</section>
