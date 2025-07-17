<script lang="ts">
	// Types
	type Variant =
		| 'default'
		| 'primary'
		| 'secondary'
		| 'success'
		| 'warning'
		| 'error'
		| 'info'
		| 'neutral';

	type Size = 'sm' | 'md' | 'lg';

	// Props
	export let variant: Variant = 'default';
	export let size: Size = 'md';
	export let className: string = '';

	// Badge variants
	const variants: Record<Variant, string> = {
		default: 'bg-gray-100 text-gray-800',
		primary: 'bg-indigo-100 text-indigo-800',
		secondary: 'bg-blue-100 text-blue-800',
		success: 'bg-green-100 text-green-800',
		warning: 'bg-yellow-100 text-yellow-800',
		error: 'bg-red-100 text-red-800',
		info: 'bg-blue-100 text-blue-800',
		neutral: 'bg-slate-100 text-slate-800'
	};

	// Size variants
	const sizes: Record<Size, string> = {
		sm: 'px-2 py-0.5 text-xs',
		md: 'px-2.5 py-0.5 text-xs',
		lg: 'px-3 py-1 text-sm'
	};

	// Base classes (from your existing style)
	const baseClasses = 'inline-flex items-center rounded-full font-medium';

	// Computed final classes
	$: badgeClasses = `${baseClasses} ${sizes[size]} ${variants[variant]} ${className}`;
</script>

<span class={badgeClasses}>
	<slot />
</span>
