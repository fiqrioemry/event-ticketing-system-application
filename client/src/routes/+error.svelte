<script lang="ts">
	import { onMount } from 'svelte';
	import { Home } from '@lucide/svelte';
	import { browser } from '$app/environment';

	let mounted = false;
	let ticketAnimation = false;

	onMount(() => {
		mounted = true;
		setTimeout(() => {
			ticketAnimation = true;
		}, 500);
	});

	function goHome() {
		if (browser) {
			window.location.href = '/';
		}
	}
</script>

<svelte:head>
	<title>404 - Page Not Found</title>
</svelte:head>

<div class="not-found-container">
	<!-- Background Elements -->
	<div class="bg-elements">
		<div class="cloud cloud-1"></div>
		<div class="cloud cloud-2"></div>
		<div class="cloud cloud-3"></div>
		<div class="plane"></div>
	</div>

	<!-- Main Content -->
	<div class="content" class:mounted>
		<!-- 404 Ticket Design -->
		<div class="ticket-container" class:animate={ticketAnimation}>
			<div class="ticket">
				<div class="ticket-left">
					<div class="ticket-hole"></div>
					<div class="ticket-hole"></div>
					<div class="ticket-hole"></div>
					<div class="ticket-hole"></div>
					<div class="ticket-hole"></div>
				</div>

				<div class="ticket-main">
					<div class="ticket-header">
						<div class="airline-logo">
							<svg width="40" height="40" viewBox="0 0 40 40">
								<circle cx="20" cy="20" r="18" fill="#3B82F6" />
								<path d="M12 20l8-8v5h8v6h-8v5l-8-8z" fill="white" />
							</svg>
						</div>
						<div class="flight-info">
							<span class="flight-number">ERROR 404</span>
							<span class="flight-status">TIDAK DITEMUKAN</span>
						</div>
					</div>

					<div class="ticket-body">
						<div class="route">
							<div class="departure">
								<span class="code">YOU</span>
								<span class="city">Lokasi Anda</span>
							</div>
							<div class="route-line">
								<div class="plane-icon">✈️</div>
							</div>
							<div class="arrival">
								<span class="code">404</span>
								<span class="city">Halaman Hilang</span>
							</div>
						</div>

						<div class="ticket-details">
							<div class="detail-item">
								<span class="label">GATE</span>
								<span class="value">404</span>
							</div>
							<div class="detail-item">
								<span class="label">SEAT</span>
								<span class="value">N/A</span>
							</div>
							<div class="detail-item">
								<span class="label">STATUS</span>
								<span class="value error">LOST</span>
							</div>
						</div>
					</div>
				</div>

				<div class="ticket-right">
					<div class="ticket-hole"></div>
					<div class="ticket-hole"></div>
					<div class="ticket-hole"></div>
					<div class="ticket-hole"></div>
					<div class="ticket-hole"></div>
				</div>
			</div>
		</div>

		<!-- Error Message -->
		<div class="error-message">
			<h1>Oops! Halaman Tidak Ditemukan</h1>
			<p>
				Sepertinya tiket yang Anda cari telah terbang ke tempat lain. Mungkin sudah kedaluwarsa atau
				tidak pernah ada.
			</p>
		</div>

		<!-- Action Buttons -->
		<div class="action-buttons">
			<button class="btn btn-primary" on:click={goHome}>
				<Home />
				Kembali ke Beranda
			</button>
		</div>
	</div>
</div>

<style>
	.not-found-container {
		min-height: 100vh;
		background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 2rem;
		position: relative;
		overflow: hidden;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
	}

	/* Background Elements */
	.bg-elements {
		position: absolute;
		width: 100%;
		height: 100%;
		top: 0;
		left: 0;
		pointer-events: none;
	}

	.cloud {
		position: absolute;
		background: rgba(255, 255, 255, 0.1);
		border-radius: 50px;
		animation: float 20s infinite ease-in-out;
	}

	.cloud::before,
	.cloud::after {
		content: '';
		position: absolute;
		background: rgba(255, 255, 255, 0.1);
		border-radius: 50px;
	}

	.cloud-1 {
		width: 100px;
		height: 40px;
		top: 20%;
		left: 10%;
		animation-delay: 0s;
	}

	.cloud-1::before {
		width: 50px;
		height: 50px;
		top: -25px;
		left: 10px;
	}

	.cloud-1::after {
		width: 60px;
		height: 40px;
		top: -15px;
		right: 10px;
	}

	.cloud-2 {
		width: 80px;
		height: 30px;
		top: 30%;
		right: 15%;
		animation-delay: -5s;
	}

	.cloud-2::before {
		width: 40px;
		height: 40px;
		top: -20px;
		left: 15px;
	}

	.cloud-2::after {
		width: 50px;
		height: 30px;
		top: -10px;
		right: 15px;
	}

	.cloud-3 {
		width: 120px;
		height: 50px;
		bottom: 20%;
		left: 20%;
		animation-delay: -10s;
	}

	.cloud-3::before {
		width: 60px;
		height: 60px;
		top: -30px;
		left: 20px;
	}

	.cloud-3::after {
		width: 70px;
		height: 50px;
		top: -20px;
		right: 20px;
	}

	.plane {
		position: absolute;
		width: 60px;
		height: 20px;
		background: rgba(255, 255, 255, 0.2);
		border-radius: 0 10px 10px 0;
		top: 15%;
		right: -100px;
		animation: fly 25s infinite linear;
	}

	.plane::before {
		content: '✈️';
		position: absolute;
		font-size: 24px;
		top: -2px;
		left: -5px;
	}

	@keyframes float {
		0%,
		100% {
			transform: translateX(0) translateY(0);
		}
		25% {
			transform: translateX(20px) translateY(-10px);
		}
		50% {
			transform: translateX(0) translateY(-20px);
		}
		75% {
			transform: translateX(-20px) translateY(-10px);
		}
	}

	@keyframes fly {
		0% {
			right: -100px;
		}
		100% {
			right: 100vw;
		}
	}

	/* Main Content */
	.content {
		text-align: center;
		max-width: 800px;
		width: 100%;
		opacity: 0;
		transform: translateY(50px);
		transition: all 0.8s ease-out;
	}

	.content.mounted {
		opacity: 1;
		transform: translateY(0);
	}

	/* Ticket Design */
	.ticket-container {
		margin-bottom: 3rem;
		perspective: 1000px;
	}

	.ticket {
		display: flex;
		background: white;
		border-radius: 12px;
		box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
		overflow: hidden;
		max-width: 600px;
		margin: 0 auto;
		transform: rotateX(10deg) rotateY(-5deg);
		transition: all 0.6s ease-out;
	}

	.ticket-container.animate .ticket {
		transform: rotateX(0deg) rotateY(0deg);
	}

	.ticket-left,
	.ticket-right {
		width: 30px;
		background: #f8f9fa;
		display: flex;
		flex-direction: column;
		justify-content: space-around;
		align-items: center;
		padding: 1rem 0;
	}

	.ticket-hole {
		width: 12px;
		height: 12px;
		border-radius: 50%;
		background: #667eea;
	}

	.ticket-main {
		flex: 1;
		padding: 2rem;
	}

	.ticket-header {
		display: flex;
		align-items: center;
		gap: 1rem;
		margin-bottom: 2rem;
	}

	.airline-logo {
		flex-shrink: 0;
	}

	.flight-info {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
	}

	.flight-number {
		font-size: 1.5rem;
		font-weight: bold;
		color: #1f2937;
	}

	.flight-status {
		font-size: 0.9rem;
		color: #ef4444;
		font-weight: 600;
	}

	.route {
		display: flex;
		align-items: center;
		justify-content: space-between;
		margin-bottom: 2rem;
	}

	.departure,
	.arrival {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.code {
		font-size: 2rem;
		font-weight: bold;
		color: #1f2937;
	}

	.city {
		font-size: 0.9rem;
		color: #6b7280;
		margin-top: 0.25rem;
	}

	.route-line {
		flex: 1;
		height: 2px;
		background: linear-gradient(to right, #3b82f6, #ef4444);
		margin: 0 2rem;
		position: relative;
		border-radius: 1px;
	}

	.plane-icon {
		position: absolute;
		top: -10px;
		left: 50%;
		transform: translateX(-50%);
		font-size: 1.2rem;
		animation: bounce 2s infinite;
	}

	.ticket-details {
		display: flex;
		justify-content: space-between;
		gap: 1rem;
	}

	.detail-item {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.label {
		font-size: 0.7rem;
		color: #9ca3af;
		text-transform: uppercase;
		font-weight: 600;
		margin-bottom: 0.25rem;
	}

	.value {
		font-size: 1rem;
		font-weight: bold;
		color: #1f2937;
	}

	.value.error {
		color: #ef4444;
	}

	/* Error Message */
	.error-message {
		margin-bottom: 3rem;
		color: white;
	}

	.error-message h1 {
		font-size: 2.5rem;
		margin-bottom: 1rem;
		font-weight: 700;
	}

	.error-message p {
		font-size: 1.2rem;
		opacity: 0.9;
		line-height: 1.6;
		max-width: 600px;
		margin: 0 auto;
	}

	/* Action Buttons */
	.action-buttons {
		display: flex;
		gap: 1rem;
		justify-content: center;
		margin-bottom: 3rem;
		flex-wrap: wrap;
	}

	.btn {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		padding: 1rem 2rem;
		border: none;
		border-radius: 8px;
		font-size: 1rem;
		font-weight: 600;
		cursor: pointer;
		transition: all 0.3s ease;
		text-decoration: none;
	}

	.btn-primary {
		background: #3b82f6;
		color: white;
	}

	.btn-primary:hover {
		background: #2563eb;
		transform: translateY(-2px);
		box-shadow: 0 10px 20px rgba(59, 130, 246, 0.3);
	}

	@keyframes bounce {
		0%,
		100% {
			transform: translateX(-50%) translateY(0);
		}
		50% {
			transform: translateX(-50%) translateY(-5px);
		}
	}

	/* Responsive Design */
	@media (max-width: 768px) {
		.not-found-container {
			padding: 1rem;
		}

		.ticket-main {
			padding: 1.5rem;
		}

		.error-message h1 {
			font-size: 2rem;
		}

		.error-message p {
			font-size: 1rem;
		}

		.route {
			flex-direction: column;
			gap: 1rem;
		}

		.route-line {
			width: 2px;
			height: 50px;
			margin: 0;
		}

		.plane-icon {
			left: -10px;
			top: 50%;
			transform: translateY(-50%) rotate(90deg);
		}

		.ticket-details {
			flex-direction: column;
			gap: 1rem;
		}

		.action-buttons {
			flex-direction: column;
			align-items: center;
		}

		.btn {
			width: 100%;
			max-width: 300px;
		}
	}
</style>
