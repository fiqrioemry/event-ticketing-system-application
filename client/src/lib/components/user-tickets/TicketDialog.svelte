<script lang="ts">
	import { getQRCodeURL } from '$lib/utils/formatter';
	import Button from '$lib/components/ui/button/button.svelte';
	import DialogContainer from '$lib/components/shared/DialogContainer.svelte';

	export let downloadQR: (
		qrData: string | null,
		type: string,
		eventName: string
	) => void = () => {};
	export let open: boolean = false;
	export let selectedQR: string | null = null;
	export let onCloseDialog: () => void = () => {};

	function closeModal() {
		open = false;
		selectedQR = null;
	}

	function handleDownload() {
		downloadQR(selectedQR, 'ticket', 'event');
		onCloseDialog();
	}
</script>

<DialogContainer {open} onCloseDialog={closeModal} size="lg">
	<!-- Modal Header -->
	<div class="border-b border-gray-200 p-6">
		<div class="flex items-center justify-between">
			<h3 class="text-lg font-semibold text-gray-900">QR Code</h3>
		</div>
	</div>

	<!-- Modal Body -->
	<div class="p-6 text-center">
		<div
			class="mx-auto mb-4 h-64 w-64 overflow-hidden rounded-lg border border-gray-200 bg-white p-4"
		>
			<img alt="QR Code" src={getQRCodeURL(selectedQR, 240)} class="h-full w-full object-contain" />
		</div>
		<p class="text-sm break-all text-gray-500">{selectedQR}</p>
		<!-- Modal Actions -->
		<div class="mt-6 flex gap-3">
			<Button class="w-full" variant="event-cyan" onclick={handleDownload}>Download</Button>
		</div>
	</div>
</DialogContainer>
