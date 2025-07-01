<script lang="ts">
	import { onMount } from 'svelte';
	import { formatRupiah } from '$lib/utils';
	import { Avatar, AvatarFallback, AvatarImage } from '$lib/components/ui/avatar';

	export let data;
	const { events } = data;

	let current = 0;
	const banners = ['banner1.webp', 'banner2.webp', 'banner3.webp'];
	onMount(() => {
		const interval = setInterval(() => {
			current = (current + 1) % banners.length;
		}, 5000);
		return () => clearInterval(interval);
	});
</script>

<svelte:head>
	<title>Home | MyBrand</title>
	<meta name="description" content="Automate your work effortlessly with AI-powered workflows." />
</svelte:head>

<!-- banner section -->
<section class="relative mx-auto max-w-7xl px-4 pt-6">
	<div class="flex h-96 overflow-hidden rounded-xl">
		<div
			class="flex transition-transform duration-700 ease-in-out"
			style="transform: translateX({-current * 100}%)"
		>
			{#each banners as banner (banner)}
				<img src={banner} alt="Event Banner" class="w-full min-w-full object-cover" />
			{/each}
		</div>
	</div>

	<!-- Dot indicators -->
	<div class="mt-2 flex justify-center gap-2">
		{#each banners as _, i}
			<div
				class="h-2 w-2 rounded-full"
				class:bg-blue-500={i === current}
				class:bg-zinc-300={i !== current}
			></div>
		{/each}
	</div>
</section>

<!-- Recommended Events -->
<section class="mx-auto max-w-7xl px-4 py-12">
	<div class="mb-6 flex items-center justify-between">
		<div>
			<h2 class="text-muted-foreground text-2xl font-bold">Exciting Events Just for You</h2>
			<p class="text-muted-foreground text-sm">Time to choose your kind of fun!</p>
		</div>
		<a href="/events" class="text-sm text-cyan-600 hover:underline">View All</a>
	</div>

	<div class="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
		{#each events as event}
			<div class="rounded-lg border bg-white shadow-sm transition hover:shadow-md">
				<img src={event.image} alt={event.title} class="h-40 w-full rounded-t-lg object-cover" />
				<div class="p-3">
					<h3 class="truncate text-sm font-semibold">{event.title}</h3>
					<p class="text-muted-foreground truncate text-xs">{event.location}</p>
					<p class="mt-1 text-sm font-semibold text-blue-600">{formatRupiah(event.startPrice)}</p>
				</div>
			</div>
		{/each}
	</div>
</section>

<!-- Testimonials / Trust Section -->
<section class="bg-muted px-4 py-16">
	<div class="mx-auto max-w-4xl text-center">
		<h2 class="text-muted-foreground mb-6 text-2xl font-bold">What Our Customers Say?</h2>
		<p class="text-muted-foreground mb-8">
			Thousands of people have enjoyed the convenience of buying their favorite event tickets!
		</p>
		<div class="flex flex-col justify-center gap-4 md:flex-row">
			<!-- Testimoni item -->

			<div
				class="bg-background flex flex-col items-center space-y-4 rounded-lg p-6 text-center shadow-md"
			>
				<Avatar class="h-9 w-9">
					<AvatarImage
						src="https://avatar.iran.liara.run/public/24"
						alt="https://avatar.iran.liara.run/public/24"
					/>
					<AvatarFallback>Amanda</AvatarFallback>
				</Avatar>
				<div>
					<p class="italic">"It was very easy, no more queue"</p>
					<span class="mt-2 block text-sm font-semibold text-zinc-700 dark:text-zinc-300"
						>– Amanda</span
					>
				</div>
			</div>
			<div
				class="bg-background flex flex-col items-center space-y-4 rounded-lg p-6 text-center shadow-md"
			>
				<Avatar class="h-9 w-9">
					<AvatarImage
						src="https://avatar.iran.liara.run/public/24"
						alt="https://avatar.iran.liara.run/public/24"
					/>
					<AvatarFallback>Danang</AvatarFallback>
				</Avatar>
				<div>
					<p class="italic">"Doesnt need too much time"</p>
					<span class="mt-2 block text-sm font-semibold text-zinc-700 dark:text-zinc-300"
						>– Danang</span
					>
				</div>
			</div>
			<div
				class="bg-background flex flex-col items-center space-y-4 rounded-lg p-6 text-center shadow-md"
			>
				<Avatar class="h-9 w-9">
					<AvatarImage
						src="https://avatar.iran.liara.run/public/24"
						alt="https://avatar.iran.liara.run/public/24"
					/>
					<AvatarFallback>Rina</AvatarFallback>
				</Avatar>
				<div>
					<p class="italic">"Really was a great experience"</p>
					<span class="mt-2 block text-sm font-semibold text-zinc-700 dark:text-zinc-300"
						>– Rina</span
					>
				</div>
			</div>
		</div>
	</div>
	<!-- Cara Beli Tiket -->
	<section class="px-4 pt-20">
		<div class="bg-background mx-auto max-w-6xl rounded-2xl px-8 py-10 text-center shadow-md">
			<h2 class="mb-2 text-2xl font-bold">How To Buy</h2>
			<p class="text-muted-foreground mb-10">Now buying tickets online is easy</p>

			<div class="grid grid-cols-1 gap-8 sm:grid-cols-2 md:grid-cols-4">
				<div class="flex flex-col items-center">
					<div class="relative mb-4">
						<div class="rounded-full bg-blue-100 p-4">
							<svg class="h-6 w-6 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
								<path d="M8 9a3 3 0 100-6 3 3 0 000 6zM2 13a6 6 0 1112 0H2z" />
							</svg>
						</div>
						<span
							class="absolute -top-2 -right-2 flex h-5 w-5 items-center justify-center rounded-full bg-orange-500 text-xs font-bold text-white"
							>1</span
						>
					</div>
					<h3 class="font-semibold">Sign New Account</h3>
					<p class="text-muted-foreground mt-1 text-sm">
						Register your account to purchase tickets.
					</p>
				</div>

				<div class="flex flex-col items-center">
					<div class="relative mb-4">
						<div class="rounded-full bg-blue-100 p-4">
							<svg class="h-6 w-6 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
								<path
									d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z"
								/>
							</svg>
						</div>
						<span
							class="absolute -top-2 -right-2 flex h-5 w-5 items-center justify-center rounded-full bg-orange-500 text-xs font-bold text-white"
							>2</span
						>
					</div>
					<h3 class="font-semibold">Select Event</h3>
					<p class="text-muted-foreground mt-1 text-sm">
						Choose the event and ticket type you want to buy.
					</p>
				</div>

				<div class="flex flex-col items-center">
					<div class="relative mb-4">
						<div class="rounded-full bg-blue-100 p-4">
							<svg class="h-6 w-6 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
								<path d="M4 4h12v2H4V4zm0 4h8v2H4V8zm0 4h12v2H4v-2z" />
							</svg>
						</div>
						<span
							class="absolute -top-2 -right-2 flex h-5 w-5 items-center justify-center rounded-full bg-orange-500 text-xs font-bold text-white"
							>3</span
						>
					</div>
					<h3 class="font-semibold">Payment</h3>
					<p class="text-muted-foreground mt-1 text-sm">Choose the payment method we provide.</p>
				</div>

				<div class="flex flex-col items-center">
					<div class="relative mb-4">
						<div class="rounded-full bg-blue-100 p-4">
							<svg class="h-6 w-6 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
								<path d="M3 4a1 1 0 011-1h12a1 1 0 011 1v11h-2v2H5v-2H3V4zm3 9h8V5H6v8z" />
							</svg>
						</div>
						<span
							class="absolute -top-2 -right-2 flex h-5 w-5 items-center justify-center rounded-full bg-orange-500 text-xs font-bold text-white"
							>4</span
						>
					</div>
					<h3 class="font-semibold">Payment Completed</h3>
					<p class="text-muted-foreground mt-1 text-sm">
						E-Ticket has been sent to your email & account.
					</p>
				</div>
			</div>
		</div>
	</section>

	<!-- Banner Promo -->
	<section class="px-4">
		<div class="mx-auto max-w-7xl">
			<img src="/bottom_banner.png" alt="Cara Beli Tiket" class="w-full rounded-xl object-cover" />
		</div>
	</section>
</section>
