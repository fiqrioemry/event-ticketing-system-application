<!-- src/lib/components/ui/ToastContainer.svelte -->
<script>
	import Toast from './Toast.svelte';
	import { toastStore } from '$lib/stores/toast.store';

	$: toasts = $toastStore.toasts;

	function handleDismiss(toastId) {
		toastStore.removeToast(toastId);
	}

	// Group toasts by position for better rendering
	$: toastsByPosition = toasts.reduce((acc, toast) => {
		const position = toast.position || 'top-right';
		if (!acc[position]) {
			acc[position] = [];
		}
		acc[position].push(toast);
		return acc;
	}, {});
</script>

<!-- Render toasts grouped by position -->
{#each Object.entries(toastsByPosition) as [position, positionToasts]}
	<div
		class="pointer-events-none fixed z-50 {position.includes('top')
			? 'top-0'
			: 'bottom-0'} {position.includes('left')
			? 'left-0'
			: position.includes('right')
				? 'right-0'
				: 'left-1/2 -translate-x-1/2 transform'} p-4"
	>
		<div class="space-y-3">
			{#each positionToasts as toast (toast.id)}
				<div class="pointer-events-auto">
					<Toast
						type={toast.type}
						title={toast.title}
						message={toast.message}
						duration={toast.duration}
						dismissible={toast.dismissible}
						position={toast.position}
						on:dismiss={() => handleDismiss(toast.id)}
					/>
				</div>
			{/each}
		</div>
	</div>
{/each}
