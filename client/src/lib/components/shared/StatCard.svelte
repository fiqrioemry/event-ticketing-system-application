<!-- StatCard.svelte -->
<script lang="ts">
	export let icon: any;
	export let title: string;
	export let value: string;
	export let variant: 'blue' | 'green' | 'blue' | 'red' | 'yellow' = 'blue';

	const variants = {
		blue: 'bg-blue-50 text-blue-900 text-blue-700 text-blue-600',
		green: 'bg-green-50 text-green-900 text-green-700 text-green-600',
		blue: 'bg-blue-50 text-blue-900 text-blue-700 text-blue-600',
		red: 'bg-red-50 text-red-900 text-red-700 text-red-600',
		yellow: 'bg-yellow-50 text-yellow-900 text-yellow-700 text-yellow-600'
	};
</script>

<div class="rounded-lg {variants[variant].split(' ')[0]} p-4">
	<div class="flex items-center">
		<div class="flex-shrink-0">
			<svelte:component this={icon} class="h-6 w-6 {variants[variant].split(' ')[3]}" />
		</div>
		<div class="ml-3">
			<p class="text-sm font-medium {variants[variant].split(' ')[1]}">{title}</p>
			<p class="text-sm {variants[variant].split(' ')[2]}">{value}</p>
		</div>
	</div>
</div>
