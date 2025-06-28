<script>
	import { Mail, PhoneCall } from '@lucide/svelte';
	import AppLogo from '$lib/components/common/AppLogo.svelte';
</script>

<footer class="border-t border-slate-800 bg-slate-900 py-10 text-slate-300">
	<div class="mx-auto grid max-w-7xl grid-cols-1 gap-10 px-4 md:grid-cols-4">
		<!-- Logo & Description -->
		<div>
			<AppLogo />
			<p class="mt-3 text-sm text-slate-400">
				Simplify your work with intelligent automation flows powered by AI.
			</p>
		</div>

		<!-- Quick Links -->
		<div>
			<h3 class="mb-3 font-semibold text-white">Quick Links</h3>
			<ul class="space-y-2 text-sm">
				<li><a href="/" class="hover:underline">Home</a></li>
				<li><a href="/signin" class="hover:underline">Login</a></li>
			</ul>
		</div>

		<!-- Product -->
		<div>
			<h3 class="mb-3 font-semibold text-white">Product</h3>
			<ul class="space-y-2 text-sm">
				<li><a href="/events" class="hover:underline">Event</a></li>
				<li><a href="/about" class="hover:underline">About</a></li>
			</ul>
		</div>

		<!-- Contact -->
		<div>
			<h3 class="mb-3 font-semibold text-white">Contact</h3>
			<ul class="space-y-2 text-sm text-slate-400">
				<li class="flex items-center gap-2">
					<Mail class="h-4 w-4" />
					<a href="mailto:support@autoflowai.com" class="hover:underline">support@tiketku.com</a>
				</li>
				<li class="flex items-center gap-2">
					<PhoneCall class="h-4 w-4" />
					<a href="tel:+6281234567890" class="hover:underline">+62 812-3456-7890</a>
				</li>
			</ul>
		</div>
	</div>

	<!-- Bottom Footer -->
	<div class="mt-10 border-t border-slate-800 px-4 pt-6 text-center text-xs text-slate-500">
		&copy; {new Date().getFullYear()} PT. Kreasi Anak Bangsa — All rights reserved.
	</div>
</footer>
