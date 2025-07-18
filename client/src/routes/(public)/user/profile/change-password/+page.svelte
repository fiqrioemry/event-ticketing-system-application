<script lang="ts">
	import { goto } from '$app/navigation';
	import * as Dialog from '$lib/components/ui/dialog/index';
	import DialogContainer from '$lib/components/shared/DialogContainer.svelte';
	import ChangePasswordForm from '$lib/components/profile/ChangePasswordForm.svelte';

	let open = true;

	function onCloseDialog() {
		goto('/user/profile');
	}
</script>

<svelte:head>
	<title>Tiketku - Change Password</title>
	<meta name="description" content="Change your password securely in the Tiketku." />
</svelte:head>

<DialogContainer {open} size="sm" {onCloseDialog}>
	<div class="p-4">
		<Dialog.Header class="mb-4 flex items-center justify-center">
			<Dialog.Title>Update Password Form</Dialog.Title>
			<Dialog.Description>Submit when you are ready to update your password.</Dialog.Description>
		</Dialog.Header>
		<ChangePasswordForm onSuccess={onCloseDialog} onCancel={onCloseDialog} />
	</div>
</DialogContainer>
