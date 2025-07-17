<script>
	import { formatDate } from '$lib/utils/formatter';
	export let event;
</script>

<section class="relative mt-8">
	<div
		class="relative mx-4 h-[32rem] overflow-hidden rounded-2xl bg-gradient-to-r from-blue-600 to-indigo-600 shadow-2xl lg:mx-8 lg:h-[36rem]"
	>
		<!-- Background Image with better overlay -->
		<img src={event.image} alt={event.title} class="absolute inset-0 h-full w-full object-cover" />

		<!-- Improved gradient overlay for better text readability -->
		<div class="absolute inset-0 bg-gradient-to-t from-black/70 via-black/30 to-transparent"></div>
		<div class="absolute inset-0 bg-gradient-to-r from-blue-900/40 to-indigo-900/40"></div>

		<!-- Content Container with better positioning -->
		<div class="absolute inset-0 flex flex-col justify-end">
			<div class="p-8 lg:p-12">
				<!-- Title with improved typography -->
				<h1
					class="mb-4 text-3xl leading-tight font-bold text-white drop-shadow-lg sm:text-4xl lg:text-5xl xl:text-6xl"
				>
					{event.title}
				</h1>

				<!-- Event details with improved layout -->
				<div class="flex flex-col gap-4 text-white/90 sm:flex-row sm:items-center sm:gap-6">
					<!-- Location -->
					<div class="group flex items-center gap-2">
						<div
							class="flex h-8 w-8 items-center justify-center rounded-full bg-white/20 backdrop-blur-sm transition-all duration-200 group-hover:bg-white/30"
						>
							<svg class="h-4 w-4" fill="currentColor" viewBox="0 0 20 20">
								<path
									fill-rule="evenodd"
									d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z"
									clip-rule="evenodd"
								/>
							</svg>
						</div>
						<span class="text-sm font-medium sm:text-base">{event.location}</span>
					</div>

					<!-- Date -->
					<div class="group flex items-center gap-2">
						<div
							class="flex h-8 w-8 items-center justify-center rounded-full bg-white/20 backdrop-blur-sm transition-all duration-200 group-hover:bg-white/30"
						>
							<svg class="h-4 w-4" fill="currentColor" viewBox="0 0 20 20">
								<path
									fill-rule="evenodd"
									d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z"
									clip-rule="evenodd"
								/>
							</svg>
						</div>
						<span class="text-sm font-medium sm:text-base">{formatDate(event.date)}</span>
					</div>

					<!-- Time -->
					<div class="group flex items-center gap-2">
						<div
							class="flex h-8 w-8 items-center justify-center rounded-full bg-white/20 backdrop-blur-sm transition-all duration-200 group-hover:bg-white/30"
						>
							<svg class="h-4 w-4" fill="currentColor" viewBox="0 0 20 20">
								<path
									fill-rule="evenodd"
									d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z"
									clip-rule="evenodd"
								/>
							</svg>
						</div>
						<span class="text-sm font-medium sm:text-base"
							>{event.startTime}:00 - {event.endTime}:00</span
						>
					</div>
				</div>
			</div>
		</div>

		<!-- Decorative elements -->
		<div class="absolute top-8 right-8 h-32 w-32 rounded-full bg-white/10 blur-3xl"></div>
		<div class="absolute right-16 bottom-16 h-24 w-24 rounded-full bg-blue-400/20 blur-2xl"></div>
	</div>
</section>
