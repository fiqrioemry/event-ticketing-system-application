<script lang="ts">
	import { Calendar, MapPin, Users } from '@lucide/svelte';
</script>

<div class="space-y-6">
	<!-- Loading Header -->
	<div class="text-center">
		<div
			class="mx-auto h-8 w-48 animate-pulse rounded-lg bg-gradient-to-r from-blue-200 to-indigo-200"
		></div>
		<div class="mx-auto mt-2 h-4 w-64 animate-pulse rounded-lg bg-gray-200"></div>
	</div>

	<!-- Loading Cards Grid -->
	<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
		{#each Array(6) as _, index}
			<div class="overflow-hidden rounded-xl border border-gray-200 bg-white shadow-sm">
				<!-- Loading Image -->
				<div class="aspect-video animate-pulse bg-gradient-to-r from-blue-100 to-indigo-100">
					<div class="flex h-full items-center justify-center">
						<Calendar class="h-12 w-12 text-blue-300" />
					</div>
				</div>

				<div class="p-6">
					<!-- Loading Title -->
					<div class="mb-3 space-y-2">
						<div class="h-6 w-3/4 animate-pulse rounded-lg bg-gray-200"></div>
						<div class="h-6 w-1/2 animate-pulse rounded-lg bg-gray-200"></div>
					</div>

					<!-- Loading Description -->
					<div class="mb-4 space-y-2">
						<div class="bg-gray-150 h-4 w-full animate-pulse rounded-lg"></div>
						<div class="bg-gray-150 h-4 w-5/6 animate-pulse rounded-lg"></div>
						<div class="bg-gray-150 h-4 w-2/3 animate-pulse rounded-lg"></div>
					</div>

					<!-- Loading Details -->
					<div class="mb-4 space-y-2">
						<div class="flex items-center">
							<Calendar class="mr-2 h-4 w-4 text-blue-300" />
							<div class="bg-gray-150 h-4 w-32 animate-pulse rounded-lg"></div>
						</div>
						<div class="flex items-center">
							<MapPin class="mr-2 h-4 w-4 text-indigo-300" />
							<div class="bg-gray-150 h-4 w-40 animate-pulse rounded-lg"></div>
						</div>
						<div class="flex items-center">
							<Users class="mr-2 h-4 w-4 text-blue-300" />
							<div class="bg-gray-150 h-4 w-24 animate-pulse rounded-lg"></div>
						</div>
					</div>

					<!-- Loading Status -->
					<div class="mb-4">
						<div class="bg-gray-150 h-6 w-20 animate-pulse rounded-full"></div>
					</div>

					<!-- Loading Button -->
					<div class="flex items-center justify-between">
						<div
							class="h-10 flex-1 animate-pulse rounded-lg bg-gradient-to-r from-blue-200 to-indigo-200"
						></div>
						<div class="bg-gray-150 ml-3 h-6 w-12 animate-pulse rounded-lg"></div>
					</div>
				</div>
			</div>
		{/each}
	</div>

	<!-- Loading Pagination -->
	<div class="rounded-xl border border-gray-200 bg-white shadow-sm">
		<div class="flex items-center justify-between px-6 py-4">
			<div class="bg-gray-150 h-5 w-32 animate-pulse rounded-lg"></div>
			<div class="flex items-center space-x-2">
				<div class="bg-gray-150 h-8 w-8 animate-pulse rounded-lg"></div>
				<div class="bg-gray-150 h-5 w-16 animate-pulse rounded-lg"></div>
				<div class="bg-gray-150 h-8 w-8 animate-pulse rounded-lg"></div>
			</div>
		</div>
	</div>
</div>

<style>
	@keyframes pulse {
		0%,
		100% {
			opacity: 1;
		}
		50% {
			opacity: 0.5;
		}
	}

	.animate-pulse {
		animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
	}

	.bg-gray-150 {
		background-color: #f9fafb;
	}
</style>
