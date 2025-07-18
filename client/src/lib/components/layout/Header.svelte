<!-- src/lib/components/layout/Header.svelte -->
<script lang="ts">
	import { page } from '$app/state';
	import { User, Menu, X } from '@lucide/svelte';
	import { authUser } from '$lib/stores/auth.store';
	import UserMenu from '$lib/components/shared/UserMenu.svelte';
	import AppLogo from '$lib/components/shared/AppLogo.svelte';
	import Button from '$lib/components/ui/button/button.svelte';

	let mobileMenuOpen = $state(false);

	function toggleMobileMenu() {
		mobileMenuOpen = !mobileMenuOpen;
	}
</script>

<header class="sticky top-0 z-50 border-b border-blue-100 bg-white/95 shadow-sm backdrop-blur-md">
	<div class="mx-auto flex max-w-7xl items-center justify-between px-4 py-3">
		<!-- Logo -->
		<AppLogo />

		<!-- Desktop Navigation -->
		<nav class="hidden items-center gap-8 md:flex">
			<a
				href="/"
				class="relative text-lg font-medium transition-colors duration-200 hover:text-blue-600"
				class:text-blue-600={page.url.pathname === '/'}
				class:font-semibold={page.url.pathname === '/'}
			>
				Home
			</a>

			<a
				href="/events"
				class="relative text-lg font-medium transition-colors duration-200 hover:text-blue-600"
				class:text-blue-600={page.url.pathname === '/events'}
				class:font-semibold={page.url.pathname === '/events'}
			>
				Events
			</a>
		</nav>

		<!-- User Actions -->
		<div class="flex items-center gap-3">
			{#if $authUser}
				<div class="hidden items-center gap-3 md:flex">
					<UserMenu user={$authUser} />
				</div>
			{:else}
				<!-- Auth Buttons -->
				<div class="hidden items-center gap-3 md:flex">
					<Button variant="event-ghost" size="sm" href="/signin">Sign In</Button>
					<Button variant="event" size="sm" href="/signup">
						<User />
						Get Started
					</Button>
				</div>
			{/if}

			<!-- Mobile Menu Button -->
			<button
				class="flex h-9 w-9 items-center justify-center rounded-lg border border-gray-300 bg-white transition-colors duration-200 hover:bg-gray-50 md:hidden"
				onclick={toggleMobileMenu}
				aria-label="Toggle mobile menu"
			>
				{#if mobileMenuOpen}
					<X class="h-4 w-4 text-gray-600" />
				{:else}
					<Menu class="h-4 w-4 text-gray-600" />
				{/if}
			</button>
		</div>
	</div>

	<!-- Mobile Menu -->
	{#if mobileMenuOpen}
		<div class="border-t border-blue-100 bg-white/95 backdrop-blur-md md:hidden">
			<div class="space-y-4 px-4 py-4">
				<!-- Mobile Navigation -->
				<nav class="space-y-3">
					<Button
						href="/"
						variant="event-ghost"
						class="h-12 w-full text-lg"
						onclick={() => (mobileMenuOpen = false)}
					>
						Home
					</Button>

					<Button
						href="/events"
						variant="event-ghost"
						class="h-12 w-full text-lg"
						onclick={() => (mobileMenuOpen = false)}
					>
						Events
					</Button>
				</nav>

				<!-- Mobile User Actions -->
				<div class="border-t border-blue-100 pt-4">
					{#if $authUser}
						<Button
							variant="event-ghost"
							href="/user/profile"
							class="h-12 w-full text-lg"
							onclick={() => (mobileMenuOpen = false)}
						>
							Dashboard
						</Button>
					{:else}
						<div class="space-y-3">
							<Button variant="event-outline" size="default" href="/signin" class="w-full">
								Sign In
							</Button>
							<Button variant="event" size="default" href="/signup" class="w-full">
								<User />
								Get Started
							</Button>
						</div>
					{/if}
				</div>
			</div>
		</div>
	{/if}
</header>

<style>
	/* Additional smooth transitions */
	header {
		transition: all 0.3s ease;
	}

	/* Backdrop blur fallback */
	@supports not (backdrop-filter: blur(12px)) {
		header {
			background-color: white;
		}
	}

	@keyframes dropdownFadeIn {
		from {
			opacity: 0;
			transform: translateY(-8px);
		}
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}
</style>
