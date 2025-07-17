<script lang="ts" module>
	import { cn, type WithElementRef } from '$lib/utils.js';
	import type { HTMLAnchorAttributes, HTMLButtonAttributes } from 'svelte/elements';
	import { type VariantProps, tv } from 'tailwind-variants';

	export const buttonVariants = tv({
		base: "hover:-translate-y-0.5 cursor-pointer focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive inline-flex shrink-0 items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium outline-none transition-all focus-visible:ring-[3px] disabled:pointer-events-none disabled:opacity-50 aria-disabled:pointer-events-none aria-disabled:opacity-50 [&_svg:not([class*='size-'])]:size-4 [&_svg]:pointer-events-none [&_svg]:shrink-0",
		variants: {
			variant: {
				default: 'bg-primary text-primary-foreground shadow-xs hover:bg-primary/90',
				destructive:
					'bg-destructive shadow-xs hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:focus-visible:ring-destructive/40 dark:bg-destructive/60 text-white',
				outline:
					'bg-background shadow-xs hover:bg-accent hover:text-accent-foreground dark:bg-input/30 dark:border-input dark:hover:bg-input/50 border',
				secondary: 'bg-secondary text-secondary-foreground shadow-xs hover:bg-secondary/80',
				ghost: 'hover:bg-accent hover:text-accent-foreground dark:hover:bg-accent/50',
				link: 'text-primary underline-offset-4 hover:underline',

				primary:
					'bg-gradient-to-r from-blue-600 to-indigo-600 text-white shadow-xs hover:from-blue-700 hover:to-indigo-700 hover:shadow-lg hover:shadow-blue-600/20 focus-visible:ring-blue-500/30 transform transition-all duration-200 hover:drop-shadow-lg',
				event:
					'bg-gradient-to-r from-blue-600 to-indigo-600 text-white shadow-lg hover:from-blue-700 hover:to-indigo-700 hover:shadow-xl focus-visible:ring-blue-500/30 transform transition-all duration-200',
				'event-outline':
					'border-2 border-blue-600 text-blue-600 bg-white hover:bg-gradient-to-r hover:from-blue-600 hover:to-indigo-600 hover:text-white hover:border-transparent shadow-sm hover:shadow-lg transition-all duration-200',
				'event-ghost':
					'text-blue-600 hover:bg-gradient-to-r hover:from-blue-50 hover:to-indigo-50 hover:text-blue-700 transition-all duration-200',
				'event-secondary':
					'bg-blue-100 text-blue-700 hover:bg-blue-200 hover:text-blue-800 shadow-sm transition-all duration-200',
				'event-cyan':
					'bg-gradient-to-r from-cyan-500 to-blue-500 text-white shadow-lg hover:from-cyan-600 hover:to-blue-600 hover:shadow-xl focus-visible:ring-cyan-500/30 transform transition-all duration-200'
			},
			size: {
				default: 'h-9 px-4 py-2 has-[>svg]:px-3',
				sm: 'h-8 gap-1.5 rounded-md px-3 has-[>svg]:px-2.5',
				lg: 'h-10 rounded-md px-6 has-[>svg]:px-4',
				xl: 'h-12 rounded-lg px-8 text-base has-[>svg]:px-6',
				icon: 'size-9',
				'icon-sm': 'size-8',
				'icon-lg': 'size-10'
			}
		},
		defaultVariants: {
			variant: 'default',
			size: 'default'
		}
	});

	export type ButtonVariant = VariantProps<typeof buttonVariants>['variant'];
	export type ButtonSize = VariantProps<typeof buttonVariants>['size'];

	export type ButtonProps = WithElementRef<HTMLButtonAttributes> &
		WithElementRef<HTMLAnchorAttributes> & {
			variant?: ButtonVariant;
			size?: ButtonSize;
			loading?: boolean;
			leftIcon?: any;
			rightIcon?: any;
		};
</script>

<script lang="ts">
	import { LoaderCircle } from '@lucide/svelte';

	let {
		class: className,
		variant = 'default',
		size = 'default',
		ref = $bindable(null),
		href = undefined,
		type = 'button',
		disabled,
		loading = false,
		leftIcon = undefined,
		rightIcon = undefined,
		children,
		...restProps
	}: ButtonProps = $props();

	// Disable button when loading
	let isDisabled = $derived(disabled || loading);
</script>

{#if href}
	<a
		bind:this={ref}
		data-slot="button"
		class={cn(buttonVariants({ variant, size }), className)}
		href={isDisabled ? undefined : href}
		aria-disabled={isDisabled}
		role={isDisabled ? 'link' : undefined}
		tabindex={isDisabled ? -1 : undefined}
		{...restProps}
	>
		{#if loading}
			<LoaderCircle class="animate-spin" />
		{:else if leftIcon}
			{@render leftIcon()}
		{/if}

		{@render children?.()}

		{#if !loading && rightIcon}
			{@render rightIcon()}
		{/if}
	</a>
{:else}
	<button
		bind:this={ref}
		data-slot="button"
		class={cn(buttonVariants({ variant, size }), className)}
		{type}
		disabled={isDisabled}
		{...restProps}
	>
		{#if loading}
			<LoaderCircle class="animate-spin" />
		{:else if leftIcon}
			{@render leftIcon()}
		{/if}

		{@render children?.()}

		{#if !loading && rightIcon}
			{@render rightIcon()}
		{/if}
	</button>
{/if}
