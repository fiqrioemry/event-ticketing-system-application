<!-- /lib/components/event-detail/PurchaseDialog.svelte -->
<script lang="ts">
	import { fly, fade } from 'svelte/transition';
	import type { EventDetail } from '$lib/types/api';
	import { formatPrice } from '$lib/utils/formatter';
	import { orderStore, isCreating } from '$lib/stores/order.store';
	import * as Dialog from '$lib/components/ui/dialog/index';
	import Button from '$lib/components/ui/button/button.svelte';
	import { purchaseValidationRules } from '$lib/utils/dto.validation';
	import DialogContainer from '$lib/components/shared/DialogContainer.svelte';
	import { validateForm, type ValidationErrors } from '$lib/utils/validation';
	import InputTextElement from '$lib/components/form-input/InputTextElement.svelte';
	import InputEmailElement from '$lib/components/form-input/InputEmailElement.svelte';
	import { CreditCard, CheckCircle, Clock } from '@lucide/svelte';

	export let event: EventDetail;
	export let totalAmount: number = 0;
	export let onCloseDialog: () => void;
	export let showPurchaseDialog: boolean = false;
	export let selectedTickets: { [key: string]: number };

	let purchaseForm = {
		fullname: '',
		email: '',
		phone: ''
	};

	function closeDialog() {
		purchaseForm = {
			fullname: '',
			email: '',
			phone: ''
		};
		onCloseDialog();
	}

	let errors: ValidationErrors = {};

	async function handleFormSubmit() {
		errors = {};

		errors = validateForm(purchaseForm, purchaseValidationRules);
		if (Object.keys(errors).length > 0) return;

		const orderData = {
			eventId: event.id,
			orderDetails: Object.entries(selectedTickets)
				.filter(([_, qty]) => qty > 0)
				.map(([ticketId, qty]) => ({ ticketId, quantity: qty })),
			fullname: purchaseForm.fullname,
			email: purchaseForm.email,
			phone: purchaseForm.phone
		};

		await orderStore.createOrder(orderData);
	}
</script>

<DialogContainer open={showPurchaseDialog} size="sm" onCloseDialog={closeDialog}>
	{#if $isCreating}
		<!-- Loading State -->
		<div class="p-8 text-center" in:fade={{ duration: 300 }}>
			<div class="mx-auto mb-6 flex h-20 w-20 items-center justify-center">
				<!-- Animated payment processing icon -->
				<div class="relative">
					<div
						class="h-16 w-16 animate-spin rounded-full border-4 border-blue-100 border-t-blue-600"
					></div>
					<div class="absolute inset-0 flex items-center justify-center">
						<CreditCard class="h-6 w-6 text-blue-600" />
					</div>
				</div>
			</div>

			<div class="mb-4">
				<h3 class="mb-2 text-xl font-semibold text-gray-900">Processing Your Payment</h3>
				<p class="text-sm text-gray-600">
					Please wait while we create your order and redirect you to payment...
				</p>
			</div>

			<!-- Processing steps -->
			<div class="space-y-3 text-left" in:fly={{ y: 20, duration: 400, delay: 200 }}>
				<div class="flex items-center gap-3 text-sm">
					<CheckCircle class="h-4 w-4 text-green-500" />
					<span class="text-gray-700">Validating ticket availability</span>
				</div>
				<div class="flex items-center gap-3 text-sm">
					<div class="h-4 w-4 animate-pulse rounded-full bg-blue-500"></div>
					<span class="text-gray-700">Creating your order</span>
				</div>
				<div class="flex items-center gap-3 text-sm">
					<Clock class="h-4 w-4 text-gray-400" />
					<span class="text-gray-500">Preparing payment gateway</span>
				</div>
			</div>

			<!-- Progress bar -->
			<div class="mt-6">
				<div class="h-2 overflow-hidden rounded-full bg-gray-200">
					<div
						class="h-full animate-pulse rounded-full bg-gradient-to-r from-blue-500 to-indigo-500"
						style="width: 60%;"
					></div>
				</div>
				<p class="mt-2 text-xs text-gray-500">This usually takes a few seconds...</p>
			</div>

			<!-- Cancel button (disabled during loading) -->
			<div class="mt-6">
				<Button type="button" variant="ghost" class="text-gray-500" disabled>Please wait...</Button>
			</div>
		</div>
	{:else}
		<!-- Normal Purchase Form -->
		<div class="p-6" in:fade={{ duration: 300 }}>
			<Dialog.Title class="mb-4">Purchase Information</Dialog.Title>

			<!-- Order Summary -->
			<div class="mb-4 rounded-lg bg-gray-50 p-4">
				<h4 class="mb-2 font-semibold text-gray-800">Order Summary</h4>
				{#each event.tickets as ticket}
					{#if selectedTickets[ticket.id] > 0}
						<div class="mb-1 flex justify-between text-sm">
							<span>{ticket.name} x{selectedTickets[ticket.id]}</span>
							<span>{formatPrice(ticket.price * selectedTickets[ticket.id])}</span>
						</div>
					{/if}
				{/each}
				<div class="mt-2 border-t pt-2">
					<div class="flex justify-between font-semibold">
						<span>Total:</span>
						<span>{formatPrice(totalAmount)}</span>
					</div>
				</div>
			</div>

			<form on:submit|preventDefault={handleFormSubmit} class="space-y-4">
				<InputTextElement
					id="fullname"
					label="Full Name"
					bind:value={purchaseForm.fullname}
					placeholder="Enter your full name"
					error={errors.fullname}
					required
				/>
				<InputEmailElement
					id="email"
					label="Email Address"
					bind:value={purchaseForm.email}
					error={errors.email}
					placeholder="Enter your email address"
					required
				/>

				<InputTextElement
					required
					id="phone"
					type="tel"
					error={errors.phone}
					label="Phone Number"
					maxlength={13}
					bind:value={purchaseForm.phone}
					placeholder="Enter your phone number"
				/>

				<div class="flex gap-3 pt-4">
					<Button type="button" variant="secondary" class="w-1/2" onclick={closeDialog}>
						Cancel
					</Button>
					<Button type="submit" variant="primary" class="w-1/2" disabled={$isCreating}>
						{#if $isCreating}
							<div class="flex items-center gap-2">
								<div
									class="h-4 w-4 animate-spin rounded-full border-2 border-white border-t-transparent"
								></div>
								Processing...
							</div>
						{:else}
							Proceed to Payment
						{/if}
					</Button>
				</div>
			</form>
		</div>
	{/if}
</DialogContainer>
