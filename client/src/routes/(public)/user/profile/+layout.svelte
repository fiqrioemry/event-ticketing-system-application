<script lang="ts">
	import { onMount } from 'svelte';
	import { toast } from 'svelte-sonner';
	import { goto } from '$app/navigation';
	import { Lock, RefreshCw } from '@lucide/svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import AlertCard from '$lib/components/shared/AlertCard.svelte';
	import SubmitButton from '$lib/components/form-input/SubmitButton.svelte';
	import ProfileHeader from '$lib/components/profile/ProfileHeader.svelte';
	import ProfileDetail from '$lib/components/profile/ProfileDetail.svelte';
	import ProfileLoading from '$lib/components/profile/ProfileLoading.svelte';
	import { userStore, isLoading, userError } from '$lib/stores/user.store';

	onMount(async () => {
		await userStore.getMyProfile();
	});

	async function handleRefresh() {
		const result = await userStore.getMyProfile();

		if (result.success) {
			toast.success('Profile refreshed!');
		} else {
			toast.error('Failed to refresh profile');
		}
	}

	async function handleErrorRetry() {
		await userStore.getMyProfile();
	}
</script>

<div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
	<!-- Error State -->
	{#if $userError}
		<AlertCard onAction={handleErrorRetry} message={$userError} />
	{/if}
	<!-- Loading State -->
	{#if $isLoading}
		<ProfileLoading />
	{:else}
		<div class="overflow-hidden rounded-lg border border-gray-200 bg-white shadow-sm">
			<!-- Profile Header -->
			<ProfileHeader />

			<!-- Profile Details -->
			<ProfileDetail />

			<!-- Action Buttons -->
			<div class="bg-muted border-muted border-t p-6">
				<div class="flex flex-wrap gap-3">
					<Button
						variant="outline"
						class="w-full md:w-60"
						onclick={() => goto('/user/profile/change-password')}
					>
						<Lock class="mr-2 h-4 w-4" />
						Change Password
					</Button>

					<SubmitButton
						buttonText="Refresh"
						isLoading={$isLoading}
						onsubmit={handleRefresh}
						className="md:w-60 w-full"
						buttonLoadingText="Refreshing..."
					>
						<RefreshCw class="mr-2 h-4 w-4" />
					</SubmitButton>
				</div>
			</div>
		</div>
	{/if}
</div>
<slot />
