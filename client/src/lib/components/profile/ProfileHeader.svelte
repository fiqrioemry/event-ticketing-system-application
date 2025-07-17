<script lang="ts">
	import { cn } from '$lib/utils';
	import { toast } from 'svelte-sonner';
	import { Calendar1, Camera, Loader2 } from '@lucide/svelte';
	import { formatDate, getAvatarInitials } from '$lib/utils/formatter';
	import { userStore, userProfile, isUpdating } from '$lib/stores/user.store';
	import Avatar from '../shared/Avatar.svelte';

	let avatarFile: File | null = null;
	let fileInput: HTMLInputElement | null = null;

	async function handleAvatarChange(event: any) {
		const file = event.target.files[0];
		if (!file) return;

		if (!file.type.startsWith('image/')) {
			toast.error('Please select an image file');
			return;
		}

		// Validate file size (max 1MB)
		if (file.size > 1 * 1024 * 1024) {
			toast.error('File size must be less than 1MB');
			return;
		}

		avatarFile = file;
		await uploadAvatar(file);
	}

	async function uploadAvatar(file: any) {
		if (!file || !$userProfile?.fullname) return;

		const updateData = {
			fullname: $userProfile?.fullname,
			avatar: file
		};

		await userStore.updateProfile(updateData);

		avatarFile = null;
	}

	function triggerFileInput() {
		if (fileInput) {
			fileInput.click();
		}
	}
</script>

<!-- Profile Header -->

<div class="bg-gradient-to-r from-blue-600 to-purple-600 px-6 py-8">
	<div class="flex items-center space-x-6">
		<!-- Avatar with Edit -->
		<div class="relative flex-shrink-0">
			<!-- Loading overlay for avatar updates -->
			{#if $isUpdating && avatarFile}
				<div
					class="absolute inset-0 z-10 flex items-center justify-center rounded-full bg-black/50"
				>
					<Loader2 class="h-8 w-8 animate-spin text-white" />
				</div>
			{/if}

			<Avatar src={$userProfile?.avatar} size={80} name={$userProfile?.fullname} />

			<!-- Avatar Edit Button -->
			<button
				class="absolute -right-1 -bottom-1 flex h-8 w-8 items-center justify-center rounded-full bg-white shadow-lg transition-colors hover:bg-gray-100"
				aria-label="Change avatar"
				on:click={triggerFileInput}
				disabled={$isUpdating}
			>
				<Camera class="h-4 w-4" />
			</button>

			<!-- Hidden File Input -->
			<input
				type="file"
				accept="image/*"
				class="hidden"
				bind:this={fileInput}
				on:change={handleAvatarChange}
			/>
		</div>

		<!-- User Headline Info -->
		<div class="flex-1">
			<div class="flex items-center space-x-2">
				<h2 class="text-2xl font-bold text-white">{$userProfile?.fullname}</h2>
			</div>
			<p class="text-blue-100">{$userProfile?.email}</p>
			<div class="mt-2 flex items-center text-blue-100">
				<Calendar1 class="mr-2 h-4 w-4" />

				<span class="text-sm">Member since {formatDate($userProfile?.joinedAt)}</span>
			</div>
		</div>
	</div>
</div>
