<!-- LeftSideContent.svelte -->
<script lang="ts">
	import { onMount } from 'svelte';
	import { fly, fade } from 'svelte/transition';
	import { Ticket, Lock } from '@lucide/svelte';

	// Props
	export let variant: 'login' | 'signup' | 'forgot-password' = 'login';
	export let title: string = '';
	export let subtitle: string = '';
	export let features: Array<{ text: string }> = [];

	let mounted = false;

	// Default configurations
	const configs = {
		login: {
			title:
				'Welcome to <span class="bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-transparent">TiketKu</span>',
			subtitle:
				'Your ultimate destination for seamless ticket booking and event management. Experience the convenience of booking tickets for concerts, movies, sports, and more.',
			features: [
				{ text: 'Secure and fast booking process' },
				{ text: 'Thousands of events available' },
				{ text: '24/7 customer support' }
			],
			showCTA: false,
			ctaText: ''
		},
		signup: {
			title:
				'Join <span class="bg-gradient-to-r from-blue-600 via-indigo-600 to-blue-700 bg-clip-text font-extrabold text-transparent">TiketKu</span> Today',
			subtitle:
				'Start your journey with us and unlock access to thousands of amazing events. Create your account and discover concerts, movies, sports, and exclusive experiences.',
			features: [
				{ text: 'Free account creation' },
				{ text: 'Access to exclusive events' },
				{ text: 'Personalized recommendations' }
			],
			showCTA: true,
			ctaText: 'Get started in less than 2 minutes'
		},
		'forgot-password': {
			title:
				'Reset Your <span class="bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-transparent">Password</span>',
			subtitle:
				"Don't worry, it happens to everyone. We'll help you get back into your TiketKu account quickly and securely.",
			features: [
				{ text: 'Secure password reset process' },
				{ text: 'Email link expires in 24 hours' },
				{ text: 'Back to booking in minutes' }
			],
			showCTA: false,
			ctaText: ''
		}
	};

	// Use props or default config
	$: config = configs[variant];
	$: displayTitle = title || config.title;
	$: displaySubtitle = subtitle || config.subtitle;
	$: displayFeatures = features.length > 0 ? features : config.features;
	$: displayShowCTA = config.showCTA;
	$: displayCTAText = config.ctaText;

	// Get main icon based on variant
	$: mainIcon = variant === 'forgot-password' ? Lock : Ticket;

	// Get gradient background based on variant
	$: backgroundGradient =
		variant === 'signup'
			? 'bg-gradient-to-br from-blue-50/80 via-indigo-50/90 to-cyan-50/70'
			: 'bg-gradient-to-br from-blue-50/80 via-indigo-50/90 to-blue-50/80';

	onMount(() => {
		mounted = true;
	});
</script>

<div class="relative hidden overflow-hidden {backgroundGradient} lg:flex lg:w-1/2">
	<!-- Background Pattern with enhanced animation -->
	<div class="absolute inset-0 opacity-10">
		<div
			class="absolute -top-40 -left-40 h-96 w-96 animate-pulse rounded-full bg-gradient-to-r from-blue-500 to-indigo-500 blur-3xl"
			style="animation-duration: 4s;"
		></div>
		<div
			class="absolute -right-40 -bottom-40 h-96 w-96 animate-pulse rounded-full bg-gradient-to-r from-indigo-500 to-blue-500 blur-3xl"
			style="animation-duration: 3s; animation-delay: 1s;"
		></div>
		{#if variant !== 'forgot-password'}
			<div
				class="absolute top-1/2 left-1/2 h-64 w-64 -translate-x-1/2 -translate-y-1/2 animate-pulse rounded-full bg-gradient-to-r from-cyan-400 to-blue-400 opacity-30 blur-2xl"
				style="animation-duration: 5s; animation-delay: 2s;"
			></div>
		{/if}
	</div>

	<!-- Additional decorative elements -->
	<div class="absolute inset-0 opacity-5">
		{#if variant === 'signup'}
			<div
				class="absolute top-1/4 right-1/4 h-48 w-48 animate-bounce rounded-full bg-gradient-to-r from-cyan-400 to-blue-500 blur-2xl"
				style="animation-duration: 8s; animation-delay: 0.8s;"
			></div>
			<div
				class="absolute bottom-1/4 left-1/4 h-32 w-32 animate-bounce rounded-full bg-gradient-to-r from-blue-400 to-indigo-500 blur-xl"
				style="animation-duration: 6.5s; animation-delay: 2s;"
			></div>
		{:else}
			<div
				class="absolute top-20 right-20 h-32 w-32 animate-bounce rounded-full bg-blue-400 blur-xl"
				style="animation-duration: 6s; animation-delay: 0.5s;"
			></div>
			<div
				class="absolute bottom-32 left-20 h-24 w-24 animate-bounce rounded-full bg-indigo-400 blur-lg"
				style="animation-duration: 7s; animation-delay: 1.5s;"
			></div>
		{/if}
	</div>

	<!-- Content -->
	<div class="relative z-10 flex flex-col justify-center px-12 xl:px-16">
		{#if mounted}
			<div class="mb-8">
				<!-- Icon with entrance animation -->
				<div
					class="mb-8 inline-flex h-16 w-16 items-center justify-center rounded-2xl border border-blue-200/50 bg-gradient-to-r from-blue-100 to-indigo-100 shadow-lg transition-all duration-300 hover:scale-110 hover:shadow-xl {variant ===
					'login'
						? 'hover:rotate-3'
						: variant === 'signup'
							? 'hover:-rotate-2 hover:bg-gradient-to-tr hover:from-blue-200 hover:to-indigo-200'
							: ''}"
					in:fly={{ y: -30, duration: 800, delay: 200 }}
				>
					<svelte:component
						this={mainIcon}
						class="h-8 w-8 text-blue-600 {variant === 'signup'
							? 'transition-transform duration-300 hover:rotate-12'
							: ''}"
					/>
				</div>

				<!-- Title with staggered animation -->
				<h1
					class="mb-6 text-4xl leading-tight font-bold text-slate-800 xl:text-5xl"
					in:fly={{ x: -50, duration: 800, delay: 400 }}
				>
					{@html displayTitle}
				</h1>

				<!-- Subtitle with fade in -->
				<p
					class="mb-8 text-lg leading-relaxed text-slate-600 xl:text-xl"
					in:fade={{ duration: 800, delay: 800 }}
				>
					{displaySubtitle}
				</p>
			</div>

			<!-- Features with enhanced animations -->
			<div class="space-y-{variant === 'signup' ? '5' : '4'}">
				{#each displayFeatures as feature, index}
					<div
						class="group flex items-center space-x-4 {variant === 'signup'
							? 'cursor-pointer hover:-m-2 hover:translate-x-3 hover:transform hover:rounded-lg hover:bg-blue-50/50 hover:p-2'
							: 'hover:translate-x-2 hover:transform'} transition-all duration-300"
						in:fly={{ x: -30, duration: 500, delay: 1000 + index * 200 }}
					>
						<div
							class="flex h-6 w-6 items-center justify-center rounded-full bg-gradient-to-r from-blue-500 to-indigo-500 shadow-sm transition-all duration-200 group-hover:scale-110 group-hover:shadow-md {variant ===
							'signup'
								? 'group-hover:scale-125 group-hover:rotate-12'
								: ''}"
						>
							<div class="h-2 w-2 rounded-full bg-white"></div>
						</div>
						<span
							class="text-slate-600 transition-colors duration-200 group-hover:text-slate-700 {variant ===
							'signup'
								? 'font-medium group-hover:font-semibold group-hover:text-slate-800'
								: ''}"
						>
							{feature.text}
						</span>
					</div>
				{/each}
			</div>

			<!-- CTA or animated dots -->
			{#if displayShowCTA}
				<!-- Call to Action hint -->
				<div class="mt-8 opacity-60" in:fade={{ duration: 800, delay: 1600 }}>
					<div
						class="flex items-center space-x-2 text-sm text-slate-500 transition-opacity duration-200 hover:opacity-80"
					>
						<div
							class="h-1 w-8 animate-pulse rounded-full bg-gradient-to-r from-blue-400 to-indigo-400 transition-colors duration-200 hover:from-blue-500 hover:to-indigo-500"
							style="animation-duration: 2.5s;"
						></div>
						<span class="transition-colors duration-200 hover:text-slate-600">{displayCTAText}</span
						>
					</div>
				</div>
			{:else}
				<!-- Animated dots -->
				<div class="mt-8 opacity-20" in:fade={{ duration: 600, delay: 1600 }}>
					<div class="flex space-x-2">
						<div
							class="h-2 w-2 animate-pulse cursor-pointer rounded-full bg-blue-400 transition-colors duration-200 hover:bg-blue-500"
							style="animation-duration: 2s;"
						></div>
						<div
							class="h-2 w-2 animate-pulse cursor-pointer rounded-full bg-indigo-400 transition-colors duration-200 hover:bg-indigo-500"
							style="animation-duration: 2s; animation-delay: 0.5s;"
						></div>
						<div
							class="h-2 w-2 animate-pulse cursor-pointer rounded-full bg-blue-400 transition-colors duration-200 hover:bg-blue-500"
							style="animation-duration: 2s; animation-delay: 1s;"
						></div>
					</div>
				</div>
			{/if}
		{/if}
	</div>
</div>
