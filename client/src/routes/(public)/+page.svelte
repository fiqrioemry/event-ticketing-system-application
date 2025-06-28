<script lang="ts">
	import { ArrowRight, BookText } from '@lucide/svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { Avatar, AvatarFallback, AvatarImage } from '$lib/components/ui/avatar';

	export let data;
	let session = data || {};

	export const popularEvents = [
		{
			id: '6e890af2-2f0f-41d2-bcc3-d9450fcd73f1',
			title: 'Konser Musik Senja',
			image: 'https://placehold.co/400x400?text=Konser+Musik+Senja',
			description: 'Nikmati malam dengan alunan musik akustik dari musisi ternama Indonesia.',
			location: 'Jakarta, Indonesia',
			date: '2025-07-12T19:00:00Z',
			status: 'active'
		},
		{
			id: 'b7de28d4-1b3d-4a8d-bf69-03b1715c23a4',
			title: 'Tech Conference 2025',
			image: 'https://placehold.co/400x400?text=Tech+Conference',
			description:
				'Konferensi teknologi terbesar tahun ini dengan pembicara dari Google, AWS, dan Microsoft.',
			location: 'Bandung Convention Center',
			date: '2025-08-02T09:00:00Z',
			status: 'active'
		},
		{
			id: '3f79e3bb-2b98-429e-b1e0-b861f9fcdbbe',
			title: 'Workshop UI/UX Design',
			image: 'https://placehold.co/400x400?text=UI%2FUX+Workshop',
			description:
				'Pelatihan intensif untuk belajar membuat desain digital yang menarik dan efisien.',
			location: 'Yogyakarta, Indonesia',
			date: '2025-07-28T13:00:00Z',
			status: 'active'
		},
		{
			id: 'df347ec2-58a1-4b8a-b16d-4829ef7ffaaa',
			title: 'Festival Kuliner Nusantara',
			image: 'https://placehold.co/400x400?text=Festival+Kuliner',
			description: 'Rasakan cita rasa dari Sabang sampai Merauke dalam satu tempat!',
			location: 'Surabaya Food Park',
			date: '2025-09-10T10:00:00Z',
			status: 'active'
		},
		{
			id: 'fa125d23-dccc-4e70-b00c-3b0c7255a3a2',
			title: 'Seminar Finansial Milenial',
			image: 'https://placehold.co/400x400?text=Seminar+Finansial',
			description: 'Pelajari cara mengatur keuangan dan berinvestasi di usia muda.',
			location: 'Online (Zoom Webinar)',
			date: '2025-07-20T18:00:00Z',
			status: 'active'
		}
	];
</script>

<svelte:head>
	<title>Home | MyBrand</title>
	<meta name="description" content="Automate your work effortlessly with AI-powered workflows." />
</svelte:head>

<!-- Hero Section -->
<section class="bg-muted min-h-[40vh]">
	<div class="mx-auto flex max-w-4xl flex-col items-center justify-center py-24 text-center">
		<h1 class="text-4xl font-bold text-zinc-900 md:text-5xl dark:text-white">
			Temukan & Beli <span class="text-primary">Tiket Event</span> Favoritmu
		</h1>
		<p class="text-muted-foreground mt-4 text-lg md:text-xl">
			Booking tiket konser, seminar, workshop, dan acara favorit lainnya hanya dalam hitungan detik.
		</p>
		<div class="mt-6 flex flex-wrap justify-center gap-4">
			<Button href="/events" variant="default" size="lg">
				Lihat Event <ArrowRight class="ml-2 h-4 w-4" />
			</Button>
		</div>
	</div>
</section>

<!-- Popular Events -->
<section class="px-4 py-16">
	<div class="mx-auto max-w-7xl">
		<h2 class="mb-8 text-center text-2xl font-bold text-zinc-800 dark:text-white">Event Populer</h2>
		<div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
			<!-- Card Event -->
			{#each popularEvents as event}
				<div
					class="rounded-xl border p-4 shadow-sm transition hover:shadow-md dark:border-zinc-800"
				>
					<img src={event.image} alt={event.title} class="h-48 w-full rounded-md object-cover" />
					<h3 class="mt-4 text-lg font-semibold text-zinc-800 dark:text-white">{event.title}</h3>
					<p class="text-muted-foreground text-sm">{event.location} — {event.date}</p>
					<div class="mt-2">
						<Button variant="secondary" size="sm">Beli Tiket</Button>
					</div>
				</div>
			{/each}
		</div>
	</div>
</section>

<!-- Testimonials / Trust Section -->
<section class="bg-slate-50 px-4 py-16 dark:bg-zinc-800">
	<div class="mx-auto max-w-4xl text-center">
		<h2 class="mb-6 text-2xl font-bold text-zinc-800 dark:text-white">Apa Kata Pengguna?</h2>
		<p class="text-muted-foreground mb-8">
			Ribuan orang sudah menikmati kemudahan beli tiket event favorit mereka!
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
					<p class="italic">"Sangat mudah dan cepat. Tidak perlu antre!"</p>
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
					<AvatarFallback>Amanda</AvatarFallback>
				</Avatar>
				<div>
					<p class="italic">"Sangat mudah dan cepat. Tidak perlu antre!"</p>
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
					<AvatarFallback>Amanda</AvatarFallback>
				</Avatar>
				<div>
					<p class="italic">"Sangat mudah dan cepat. Tidak perlu antre!"</p>
					<span class="mt-2 block text-sm font-semibold text-zinc-700 dark:text-zinc-300"
						>– Amanda</span
					>
				</div>
			</div>
		</div>
	</div>
</section>
