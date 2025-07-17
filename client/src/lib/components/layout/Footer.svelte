<script>
	import {
		Mail,
		PhoneCall,
		Calendar,
		MapPin,
		Twitter,
		Instagram,
		Facebook,
		Linkedin
	} from '@lucide/svelte';
</script>

<footer class="bg-gray-900 text-white">
	<!-- Main Footer Content -->
	<div class="mx-auto max-w-7xl px-4 py-16">
		<div class="grid grid-cols-1 gap-8 md:grid-cols-2 lg:grid-cols-4">
			<!-- Logo & Description -->
			<div class="lg:col-span-1">
				<!-- Logo -->
				<div class="mb-6 flex items-center gap-2">
					<div
						class="flex h-10 w-10 items-center justify-center rounded-lg bg-gradient-to-r from-blue-600 to-indigo-600 text-white shadow-lg"
					>
						<Calendar class="h-5 w-5" />
					</div>
					<span
						class="bg-gradient-to-r from-blue-400 to-indigo-400 bg-clip-text text-xl font-bold text-transparent"
					>
						TiketKu
					</span>
				</div>

				<!-- Description -->
				<p class="mb-6 leading-relaxed text-gray-300">
					Discover amazing events, connect with people, and create unforgettable experiences. Your
					gateway to the best events happening around you.
				</p>

				<!-- Social Media -->
				<div class="flex gap-3">
					<span
						class="flex h-9 w-9 items-center justify-center rounded-lg bg-gray-800 text-gray-400 transition-all duration-200 hover:bg-gradient-to-r hover:from-blue-600 hover:to-indigo-600 hover:text-white hover:shadow-lg"
						aria-label="Twitter"
					>
						<Twitter class="h-4 w-4" />
					</span>
					<span
						class="flex h-9 w-9 items-center justify-center rounded-lg bg-gray-800 text-gray-400 transition-all duration-200 hover:bg-gradient-to-r hover:from-blue-600 hover:to-indigo-600 hover:text-white hover:shadow-lg"
						aria-label="Instagram"
					>
						<Instagram class="h-4 w-4" />
					</span>
					<span
						class="flex h-9 w-9 items-center justify-center rounded-lg bg-gray-800 text-gray-400 transition-all duration-200 hover:bg-gradient-to-r hover:from-blue-600 hover:to-indigo-600 hover:text-white hover:shadow-lg"
						aria-label="Facebook"
					>
						<Facebook class="h-4 w-4" />
					</span>
					<span
						class="flex h-9 w-9 items-center justify-center rounded-lg bg-gray-800 text-gray-400 transition-all duration-200 hover:bg-gradient-to-r hover:from-blue-600 hover:to-indigo-600 hover:text-white hover:shadow-lg"
						aria-label="LinkedIn"
					>
						<Linkedin class="h-4 w-4" />
					</span>
				</div>
			</div>

			<!-- Quick Links -->
			<div>
				<h3 class="mb-6 text-lg font-semibold text-white">Quick Links</h3>
				<ul class="space-y-3">
					<li>
						<a
							href="/"
							class="text-gray-300 decoration-blue-400 underline-offset-4 transition-colors duration-200 hover:text-blue-400 hover:underline"
						>
							Home
						</a>
					</li>
					<li>
						<a
							href="/events"
							class="text-gray-300 decoration-blue-400 underline-offset-4 transition-colors duration-200 hover:text-blue-400 hover:underline"
						>
							Browse Events
						</a>
					</li>
					<li>
						<a
							href="/about"
							class="text-gray-300 decoration-blue-400 underline-offset-4 transition-colors duration-200 hover:text-blue-400 hover:underline"
						>
							About Us
						</a>
					</li>
					<li>
						<a
							href="/signin"
							class="text-gray-300 decoration-blue-400 underline-offset-4 transition-colors duration-200 hover:text-blue-400 hover:underline"
						>
							Sign In
						</a>
					</li>
				</ul>
			</div>

			<!-- Support -->
			<div>
				<h3 class="mb-6 text-lg font-semibold text-white">Support</h3>
				<ul class="space-y-3">
					<li>
						<span
							class="text-gray-300 decoration-blue-400 underline-offset-4 transition-colors duration-200 hover:text-blue-400 hover:underline"
						>
							Help Center
						</span>
					</li>
					<li>
						<span
							class="text-gray-300 decoration-blue-400 underline-offset-4 transition-colors duration-200 hover:text-blue-400 hover:underline"
						>
							FAQ
						</span>
					</li>
					<li>
						<span
							class="text-gray-300 decoration-blue-400 underline-offset-4 transition-colors duration-200 hover:text-blue-400 hover:underline"
						>
							Privacy Policy
						</span>
					</li>
					<li>
						<span
							class="text-gray-300 decoration-blue-400 underline-offset-4 transition-colors duration-200 hover:text-blue-400 hover:underline"
						>
							Terms of Service
						</span>
					</li>
					<li>
						<span
							class="text-gray-300 decoration-blue-400 underline-offset-4 transition-colors duration-200 hover:text-blue-400 hover:underline"
						>
							Community Guidelines
						</span>
					</li>
				</ul>
			</div>

			<!-- Contact -->
			<div>
				<h3 class="mb-6 text-lg font-semibold text-white">Get in Touch</h3>
				<ul class="space-y-4">
					<li>
						<a
							href="mailto:support@TiketKu.com"
							class="group flex items-center gap-3 text-gray-300 transition-colors duration-200 hover:text-blue-400"
						>
							<div
								class="flex h-8 w-8 items-center justify-center rounded-lg bg-gray-800 transition-colors duration-200 group-hover:bg-blue-600"
							>
								<Mail class="h-4 w-4" />
							</div>
							<div>
								<span class="text-sm font-medium">Email</span>
								<p class="text-xs text-gray-400">support@TiketKu.com</p>
							</div>
						</a>
					</li>
					<li>
						<a
							href="tel:+6281234567890"
							class="group flex items-center gap-3 text-gray-300 transition-colors duration-200 hover:text-blue-400"
						>
							<div
								class="flex h-8 w-8 items-center justify-center rounded-lg bg-gray-800 transition-colors duration-200 group-hover:bg-blue-600"
							>
								<PhoneCall class="h-4 w-4" />
							</div>
							<div>
								<span class="text-sm font-medium">Phone</span>
								<p class="text-xs text-gray-400">+62 812-3456-7890</p>
							</div>
						</a>
					</li>
					<li>
						<div class="flex items-center gap-3 text-gray-300">
							<div class="flex h-8 w-8 items-center justify-center rounded-lg bg-gray-800">
								<MapPin class="h-4 w-4" />
							</div>
							<div>
								<span class="text-sm font-medium">Address</span>
								<p class="text-xs text-gray-400">Jakarta, Indonesia</p>
							</div>
						</div>
					</li>
				</ul>
			</div>
		</div>
	</div>

	<!-- Newsletter Section -->
	<div class="border-t border-gray-800">
		<div class="mx-auto max-w-7xl px-4 py-8">
			<div class="flex flex-col items-center text-center">
				<h3 class="mb-2 text-lg font-semibold text-white">Stay Updated</h3>
				<p class="mb-6 max-w-md text-gray-300">
					Get notified about new events and exclusive offers directly to your inbox.
				</p>
				<div class="flex w-full max-w-md gap-2">
					<input
						type="email"
						placeholder="Enter your email"
						class="flex-1 rounded-lg border border-gray-700 bg-gray-800 px-4 py-2 text-white placeholder-gray-400 focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 focus:outline-none"
					/>
					<button
						class="rounded-lg bg-gradient-to-r from-blue-600 to-indigo-600 px-6 py-2 font-medium text-white transition-all duration-200 hover:from-blue-700 hover:to-indigo-700 hover:shadow-lg focus:ring-2 focus:ring-blue-500/20 focus:outline-none"
					>
						Subscribe
					</button>
				</div>
			</div>
		</div>
	</div>

	<!-- Bottom Footer -->
	<div class="border-t border-gray-800 bg-gray-950">
		<div class="mx-auto max-w-7xl px-4 py-6">
			<div class="flex flex-col items-center justify-between gap-4 md:flex-row">
				<p class="text-sm text-gray-400">
					&copy; {new Date().getFullYear()} TiketKu. All rights reserved.
				</p>
				<div class="flex items-center gap-1 text-sm text-gray-400">
					<span>Made with</span>
					<span class="text-indigo-400">♥</span>
					<span>by</span>
					<a
						href="https://ahmadfiqrioemry.com"
						class="font-medium text-blue-400 transition-colors duration-200 hover:text-blue-300"
						target="_blank"
						rel="noopener noreferrer"
					>
						Ahmad Fiqri Oemry
					</a>
				</div>
			</div>
		</div>
	</div>
</footer>
