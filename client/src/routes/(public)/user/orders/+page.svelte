<script lang="ts">
	import { goto } from '$app/navigation';
	import Pagination from '$lib/components/shared/Pagination.svelte';
	import PageHeading from '$lib/components/shared/PageHeading.svelte';
	import OrdersFilters from '$lib/components/orders/OrdersFilters.svelte';
	import OrdersList from '$lib/components/orders/OrdersList.svelte';

	export let data;
	$: orders = data?.orders || [];
	$: pagination = data?.pagination || { page: 1, limit: 10, total: 0 };
	let currentParams = data.params;

	// Filter state
	let searchQuery = currentParams.search || '';
	let statusFilter = currentParams.status || '';
	let perPage = currentParams.limit || 5;
	let currentPage = currentParams.page || 1;

	function updateURL() {
		const params = new URLSearchParams();
		if (searchQuery) params.set('search', searchQuery);
		if (statusFilter) params.set('status', statusFilter);
		if (perPage !== 10) params.set('limit', perPage.toString());
		if (currentPage > 1) params.set('page', currentPage.toString());

		const newURL = params.toString() ? `?${params.toString()}` : '/user/orders';
		goto(newURL, { replaceState: true });
	}

	function resetFilters() {
		searchQuery = '';
		statusFilter = '';
		currentPage = 1;
		perPage = 10;
		updateURL();
	}

	function handleSearch() {
		currentPage = 1;
		updateURL();
	}

	function handleFilterChange() {
		currentPage = 1;
		updateURL();
	}

	function onPageChange(page: number) {
		currentPage = page;
		updateURL();
	}

	function viewOrderDetail(orderId: string) {
		goto(`/user/orders/${orderId}`);
	}

	function viewTickets(orderId: string) {
		goto(`/user/orders/${orderId}/tickets`);
	}

	$: hasActiveFilters = !!(searchQuery || statusFilter);
</script>

<svelte:head>
	<title>Tiketku - Your Orders</title>
	<meta name="description" content="View and manage your orders in Tiketku." />
</svelte:head>

<PageHeading title="Your Orders" description="Manage your orders and view your purchase history." />

<OrdersFilters
	bind:searchQuery
	bind:statusFilter
	{hasActiveFilters}
	onSearch={handleSearch}
	onFilterChange={handleFilterChange}
	onReset={resetFilters}
/>

<OrdersList
	{orders}
	{hasActiveFilters}
	onViewDetail={viewOrderDetail}
	onViewTickets={viewTickets}
	onResetFilters={resetFilters}
/>

{#if orders.length > 0}
	<div class="mt-8">
		<Pagination
			{perPage}
			{currentPage}
			{onPageChange}
			total={pagination.total}
			totalPages={pagination.totalPages}
		/>
	</div>
{/if}
