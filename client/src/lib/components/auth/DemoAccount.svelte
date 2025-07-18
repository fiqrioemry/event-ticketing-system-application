<script lang="ts">
	import { slide, fade } from 'svelte/transition';
	import { quintOut } from 'svelte/easing';

	let isOpen = false;
	let showCopiedEmail = false;
	let showCopiedPassword = false;

	const demoAccount = {
		email: 'alice@event.com',
		password: '123456'
	};

	function toggleDemo() {
		isOpen = !isOpen;
	}

	async function copyToClipboard(text: string, type: 'email' | 'password') {
		try {
			await navigator.clipboard.writeText(text);

			if (type === 'email') {
				showCopiedEmail = true;
				setTimeout(() => (showCopiedEmail = false), 2000);
			} else {
				showCopiedPassword = true;
				setTimeout(() => (showCopiedPassword = false), 2000);
			}
		} catch (err) {
			console.error('Failed to copy: ', err);
		}
	}

	function fillDemoData() {
		// Dispatch custom event to parent component
		const event = new CustomEvent('fillDemo', {
			detail: {
				email: demoAccount.email,
				password: demoAccount.password
			}
		});
		window.dispatchEvent(event);
		isOpen = false;
	}
</script>

<!-- Floating Button Container -->
<div class="absolute right-14 bottom-14 z-50">
	<!-- Demo Info Card -->
	{#if isOpen}
		<div
			class="absolute right-0 bottom-16 w-72 overflow-hidden rounded-lg border border-gray-200 bg-white shadow-2xl"
			transition:slide={{ duration: 300, easing: quintOut }}
		>
			<!-- Header -->
			<div class="bg-gradient-to-r from-blue-500 to-purple-600 px-4 py-3">
				<div class="flex items-center gap-2">
					<div class="h-2 w-2 animate-pulse rounded-full bg-white"></div>
					<h3 class="text-sm font-semibold text-white">Demo Account</h3>
				</div>
			</div>

			<!-- Content -->
			<div class="space-y-4 p-4">
				<p class="text-xs leading-relaxed text-gray-600">
					Use this demo account to explore the application features without creating a real account.
				</p>

				<!-- Email Section -->
				<div class="space-y-2">
					<label for="email" class="text-xs font-medium tracking-wide text-gray-700 uppercase">
						Email Address
					</label>
					<div class="flex items-center gap-2">
						<div
							class="flex-1 rounded-md border bg-gray-50 px-3 py-2 font-mono text-sm text-gray-800"
						>
							{demoAccount.email}
						</div>
						<button
							on:click={() => copyToClipboard(demoAccount.email, 'email')}
							class="rounded-md p-2 text-gray-500 transition-colors duration-200 hover:bg-blue-50 hover:text-blue-600"
							title="Copy email"
						>
							{#if showCopiedEmail}
								<svg class="h-4 w-4 text-green-600" fill="currentColor" viewBox="0 0 20 20">
									<path
										fill-rule="evenodd"
										d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
										clip-rule="evenodd"
									/>
								</svg>
							{:else}
								<svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"
									/>
								</svg>
							{/if}
						</button>
					</div>
					{#if showCopiedEmail}
						<p class="text-xs text-green-600" transition:fade>Email copied!</p>
					{/if}
				</div>

				<!-- Password Section -->
				<div class="space-y-2">
					<label for="password" class="text-xs font-medium tracking-wide text-gray-700 uppercase">
						Password
					</label>
					<div class="flex items-center gap-2">
						<div
							class="flex-1 rounded-md border bg-gray-50 px-3 py-2 font-mono text-sm text-gray-800"
						>
							{demoAccount.password}
						</div>
						<button
							on:click={() => copyToClipboard(demoAccount.password, 'password')}
							class="rounded-md p-2 text-gray-500 transition-colors duration-200 hover:bg-blue-50 hover:text-blue-600"
							title="Copy password"
						>
							{#if showCopiedPassword}
								<svg class="h-4 w-4 text-green-600" fill="currentColor" viewBox="0 0 20 20">
									<path
										fill-rule="evenodd"
										d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
										clip-rule="evenodd"
									/>
								</svg>
							{:else}
								<svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"
									/>
								</svg>
							{/if}
						</button>
					</div>
					{#if showCopiedPassword}
						<p class="text-xs text-green-600" transition:fade>Password copied!</p>
					{/if}
				</div>

				<!-- Warning -->
				<div class="rounded-md border border-amber-200 bg-amber-50 p-3">
					<div class="flex items-center gap-2">
						<svg
							class="h-4 w-4 flex-shrink-0 text-amber-600"
							fill="currentColor"
							viewBox="0 0 20 20"
						>
							<path
								fill-rule="evenodd"
								d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
								clip-rule="evenodd"
							/>
						</svg>
						<p class="text-xs text-amber-800">This is a demo account for testing purposes only.</p>
					</div>
				</div>
			</div>
		</div>
	{/if}

	<!-- Floating Button -->
	<button
		on:click={toggleDemo}
		class="group relative flex h-14 w-14 transform items-center justify-center overflow-hidden rounded-full bg-gradient-to-r from-blue-500 to-purple-600 text-white shadow-lg transition-all duration-300 hover:scale-110 hover:shadow-xl active:scale-95"
		class:rotate-45={isOpen}
		title="Demo Account Info"
	>
		<!-- Background Animation -->
		<div
			class="absolute inset-0 rounded-full bg-gradient-to-r from-purple-600 to-blue-500 opacity-0 transition-opacity duration-300 group-hover:opacity-100"
		></div>

		<!-- Icon -->
		<div class="relative z-10 transition-transform duration-300" class:rotate-180={isOpen}>
			{#if isOpen}
				<svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M6 18L18 6M6 6l12 12"
					/>
				</svg>
			{:else}
				<svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
					/>
				</svg>
			{/if}
		</div>

		<!-- Ripple Effect -->
		<div
			class="absolute inset-0 scale-0 rounded-full bg-white opacity-20 transition-transform duration-200 group-active:scale-100"
		></div>
	</button>

	<!-- Floating Dots Animation -->
	{#if !isOpen}
		<div class="absolute -top-1 -right-1 h-3 w-3">
			<div class="h-2 w-2 animate-ping rounded-full bg-red-500"></div>
			<div class="absolute top-0 h-2 w-2 rounded-full bg-red-500"></div>
		</div>
	{/if}
</div>

<style>
	@keyframes float {
		0%,
		100% {
			transform: translateY(0px);
		}
		50% {
			transform: translateY(-10px);
		}
	}
</style>
