<script lang="ts">
	import { toast } from 'svelte-sonner';
	import { formatDate } from '$lib/utils/formatter';
	import { CheckCircle, Clock, User } from '@lucide/svelte';
	import StatCard from '$lib/components/shared/StatCard.svelte';
	import { userStore, isUpdating, userProfile } from '$lib/stores/user.store';
	import InlineTextEdit from '$lib/components/form-input/InlineTextEdit.svelte';

	async function handleUpdateFullname(newValue: string) {
		if (!newValue || newValue.trim() === '' || newValue.trim().length < 5) {
			toast.error('Fullname required and must be at least 5 characters');
			return false;
		}
		await userStore.updateProfile({ fullname: newValue });
	}
</script>

<div class="px-6 py-6">
	<div class="grid grid-cols-1 gap-6 md:grid-cols-2">
		<!-- Account Information -->
		<div>
			<h3 class="text-foreground mb-4 text-lg font-medium">Account Information</h3>
			<div class="space-y-4">
				<InlineTextEdit
					name="fullname"
					label="Fullname"
					loading={$isUpdating}
					value={$userProfile?.fullname}
					onSave={handleUpdateFullname}
				/>

				<!-- Email Address (Read Only) -->
				<div>
					<span class="block text-sm font-medium">Email Address</span>
					<div class="mt-1 flex items-center justify-between rounded-md bg-gray-100 px-3 py-2">
						<span class="text-foreground text-sm">{$userProfile?.email}</span>
						<span class="text-xs text-gray-500">Cannot be changed</span>
					</div>
				</div>
			</div>
		</div>

		<!-- Profile Statistics -->
		<div>
			<h3 class="text-foreground mb-4 text-lg font-medium">Account Statistics</h3>
			<div class="space-y-4">
				<StatCard icon={User} title="Account Status" value="Active" variant="blue" />
				<StatCard icon={CheckCircle} title="Email Verified" value="Yes" variant="green" />
				<StatCard
					icon={Clock}
					variant="green"
					title="Member since"
					value={formatDate($userProfile?.joinedAt)}
				/>
			</div>
		</div>
	</div>
</div>
