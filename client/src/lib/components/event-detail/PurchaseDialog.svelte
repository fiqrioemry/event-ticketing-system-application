<!-- /lib/components/event-detail/PurchaseDialog.svelte -->
<script lang="ts">
	import type { EventDetail } from '$lib/types/api';
	import { formatPrice } from '$lib/utils/formatter';
	import * as Dialog from '$lib/components/ui/dialog/index';
	import Button from '$lib/components/ui/button/button.svelte';
	import { purchaseValidationRules } from '$lib/utils/dto.validation';
	import DialogContainer from '$lib/components/shared/DialogContainer.svelte';
	import { validateForm, type ValidationErrors } from '$lib/utils/validation';
	import InputTextElement from '$lib/components/form-input/InputTextElement.svelte';
	import InputEmailElement from '$lib/components/form-input/InputEmailElement.svelte';
	import { orderStore } from '$lib/stores/order.store';

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

		// errors = validateForm(purchaseForm, purchaseValidationRules);
		// if (Object.keys(errors).length > 0) return;

		const orderData = {
			eventId: event.id,
			orderDetails: Object.entries(selectedTickets)
				.filter(([_, qty]) => qty > 0)
				.map(([ticketId, qty]) => ({ ticketId, quantity: qty })),
			fullname: purchaseForm.fullname,
			email: purchaseForm.email,
			phone: purchaseForm.phone
		};
		console.log('Submitting purchase form:', orderData);

		await orderStore.createOrder(orderData);
	}
</script>

<DialogContainer open={showPurchaseDialog} size="sm" onCloseDialog={closeDialog}>
	<div class="p-6">
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
				<Button type="button" variant="secondary" class="w-1/2" onclick={closeDialog}>Cancel</Button
				>
				<Button type="submit" variant="primary" class="w-1/2">Proceed to Payment</Button>
			</div>
		</form>
	</div>
</DialogContainer>
