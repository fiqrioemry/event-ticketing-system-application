<script lang="ts">
	import { CheckCircle, Clock, XCircle, AlertTriangle, Info, Circle } from '@lucide/svelte';

	// Types (same as above)
	type StatusConfig = {
		[key: string]: {
			bg: string;
			text: string;
			icon?: any;
		};
	};

	type Variant = 'default' | 'success' | 'warning' | 'error' | 'info' | 'neutral';
	type Size = 'sm' | 'md' | 'lg';

	// Props
	export let value: string = '';
	export let variant: Variant | null = null;
	export let size: Size = 'md';
	export let customConfig: StatusConfig = {};
	export let fallbackVariant: Variant = 'neutral';
	export let showIcon: boolean = false;
	export let className: string = '';

	// Predefined variant styles with icons
	const variantStyles: Record<Variant, { bg: string; text: string; icon: any }> = {
		success: { bg: 'bg-green-100', text: 'text-green-800', icon: CheckCircle },
		warning: { bg: 'bg-yellow-100', text: 'text-yellow-800', icon: AlertTriangle },
		error: { bg: 'bg-red-100', text: 'text-red-800', icon: XCircle },
		info: { bg: 'bg-blue-100', text: 'text-blue-800', icon: Info },
		neutral: { bg: 'bg-gray-100', text: 'text-gray-800', icon: Circle },
		default: { bg: 'bg-gray-100', text: 'text-gray-800', icon: Circle }
	};

	// Size styles
	const sizeStyles: Record<Size, { badge: string; icon: string }> = {
		sm: { badge: 'px-2 py-0.5 text-xs', icon: 'h-3 w-3' },
		md: { badge: 'px-2.5 py-0.5 text-xs', icon: 'h-3.5 w-3.5' },
		lg: { badge: 'px-3 py-1 text-sm', icon: 'h-4 w-4' }
	};

	// Auto-detect variant (same function as above)
	function autoDetectVariant(status: string): Variant {
		const statusLower = status.toLowerCase();

		if (
			[
				'excellent',
				'active',
				'approved',
				'completed',
				'success',
				'published',
				'verified',
				'confirmed'
			].includes(statusLower)
		) {
			return 'success';
		}

		if (
			['fair', 'pending', 'warning', 'draft', 'partial', 'limited', 'review'].includes(statusLower)
		) {
			return 'warning';
		}

		if (
			[
				'poor',
				'failed',
				'error',
				'rejected',
				'cancelled',
				'expired',
				'blocked',
				'disabled'
			].includes(statusLower)
		) {
			return 'error';
		}

		if (['good', 'info', 'processing', 'in progress', 'scheduled', 'new'].includes(statusLower)) {
			return 'info';
		}

		return fallbackVariant;
	}

	// Get final styles
	$: finalVariant = variant || autoDetectVariant(value);
	$: customStyle = customConfig[value];
	$: styles = customStyle || variantStyles[finalVariant];
	$: iconComponent = customStyle?.icon || styles.icon;
	$: baseClasses = 'inline-flex items-center rounded-full font-medium';
</script>

<span class="{baseClasses} {sizeStyles[size].badge} {styles.bg} {styles.text} {className}">
	{#if showIcon}
		<svelte:component this={iconComponent} class="{sizeStyles[size].icon} mr-1.5" />
	{/if}
	{value}
</span>
