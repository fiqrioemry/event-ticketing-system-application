<!-- src/lib/components/ui/ProfileSkeleton.svelte -->
<script>
	import { Skeleton } from '$lib/components/ui/skeleton';
	export let className = '';
</script>

<div class="overflow-hidden rounded-lg border border-gray-200 bg-white shadow-sm {className}">
	<!-- Profile Header Skeleton -->
	<div class="bg-gradient-to-r from-blue-600 to-purple-600 px-6 py-8">
		<div class="flex items-center space-x-6">
			<!-- Avatar Skeleton -->
			<div class="relative flex-shrink-0">
				<Skeleton class="h-20 w-20 rounded-full border-4 border-white bg-white/20" />

				<!-- Avatar Edit Button Skeleton -->
				<Skeleton class="absolute -right-1 -bottom-1 h-8 w-8 rounded-full bg-white/30" />
			</div>

			<!-- User Info Skeleton -->
			<div class="flex-1 space-y-2">
				<!-- Name Skeleton -->
				<Skeleton class="h-8 w-48 bg-white/20" />

				<!-- Email Skeleton -->
				<Skeleton class="h-4 w-64 bg-white/15" />

				<!-- Member Since Skeleton -->
				<div class="mt-2 flex items-center space-x-2">
					<Skeleton class="h-4 w-4 bg-white/15" />
					<Skeleton class="h-4 w-36 bg-white/15" />
				</div>
			</div>
		</div>
	</div>

	<!-- Profile Details Skeleton -->
	<div class="px-6 py-6">
		<div class="grid grid-cols-1 gap-6 md:grid-cols-2">
			<!-- Account Information Section -->
			<div>
				<!-- Section Title Skeleton -->
				<Skeleton class="mb-4 h-6 w-48" />

				<div class="space-y-4">
					<!-- Fullname Field Skeleton -->
					<div>
						<Skeleton class="mb-1 h-4 w-16" />
						<div class="flex items-center justify-between rounded-md bg-gray-100 px-3 py-2">
							<Skeleton class="h-4 w-32" />
							<Skeleton class="h-4 w-4" />
						</div>
					</div>

					<!-- Email Field Skeleton -->
					<div>
						<Skeleton class="mb-1 h-4 w-24" />
						<div class="flex items-center justify-between rounded-md bg-gray-100 px-3 py-2">
							<Skeleton class="h-4 w-40" />
							<Skeleton class="h-3 w-20" />
						</div>
					</div>
				</div>
			</div>

			<!-- Profile Statistics Section -->
			<div>
				<!-- Section Title Skeleton -->
				<Skeleton class="mb-4 h-6 w-40" />

				<div class="space-y-4">
					<!-- Stat Card Skeletons -->
					{#each Array(3) as _, i}
						<div class="rounded-lg border bg-gray-50 p-4">
							<div class="flex items-center">
								<div class="flex-shrink-0">
									<Skeleton class="h-6 w-6" />
								</div>
								<div class="ml-3 flex-1 space-y-2">
									<Skeleton class="h-4 w-24" />
									<Skeleton class="h-4 w-16" />
								</div>
							</div>
						</div>
					{/each}
				</div>
			</div>
		</div>
	</div>

	<!-- Action Buttons Skeleton -->
	<div class="bg-muted border-muted border-t p-6">
		<div class="flex flex-wrap gap-3">
			<!-- Button Skeletons -->
			<Skeleton class="h-10 w-full md:w-60" />
			<Skeleton class="h-10 w-full md:w-60" />
		</div>
	</div>
</div>
