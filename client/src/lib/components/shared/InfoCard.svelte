<!-- InfoCard.svelte -->
<script lang="ts">
	import type { Component } from 'svelte';

	export let value: string | null = null;
	export let icon: Component;
	export let variant: 'purple' | 'blue' | 'green' | 'red' | 'yellow' | 'gray' = 'purple';
	export let size: 'sm' | 'md' | 'lg' = 'md';

	// Definisi warna untuk setiap variant
	const variants = {
		purple: {
			container: 'border-purple-100 bg-purple-50',
			icon: 'text-purple-600',
			value: 'text-purple-800'
		},
		blue: {
			container: 'border-blue-100 bg-blue-50',
			icon: 'text-blue-600',
			value: 'text-blue-800'
		},
		green: {
			container: 'border-green-100 bg-green-50',
			icon: 'text-green-600',
			value: 'text-green-800'
		},
		red: {
			container: 'border-red-100 bg-red-50',
			icon: 'text-red-600',
			value: 'text-red-800'
		},
		yellow: {
			container: 'border-yellow-100 bg-yellow-50',
			icon: 'text-yellow-600',
			value: 'text-yellow-800'
		},
		gray: {
			container: 'border-gray-100 bg-gray-50',
			icon: 'text-gray-600',
			value: 'text-gray-800'
		}
	};

	// Definisi ukuran
	const sizes = {
		sm: {
			container: 'p-3',
			icon: 'h-6 w-6',
			value: 'text-xs'
		},
		md: {
			container: 'p-4',
			icon: 'h-8 w-8',
			value: 'text-sm'
		},
		lg: {
			container: 'p-5',
			icon: 'h-10 w-10',
			label: 'text-sm',
			value: 'text-base'
		}
	};

	$: currentVariant = variants[variant];
	$: currentSize = sizes[size];
</script>

<div class="rounded-lg border bg-red-500 {currentVariant.container} {currentSize.container}">
	<div class="mb-1 flex flex-col justify-center">
		<div class="flex items-center justify-center">
			<svelte:component this={icon} class="{currentSize.icon} {currentVariant.icon}" />
		</div>
	</div>
	<p class="{currentSize.value} font-medium {currentVariant.value}">
		{value ? value : '-'}
	</p>
</div>
