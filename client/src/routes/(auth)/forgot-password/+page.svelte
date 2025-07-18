<script>
	import ForgotForm from '$lib/components/auth/ForgotForm.svelte';
	import LeftSideContent from '$lib/components/auth/LeftSideContent.svelte';
	import NavigationHeader from '$lib/components/auth/NavigationHeader.svelte';
</script>

<svelte:head>
	<title>Forgot Password | TiketKu</title>
	<meta name="description" content="Reset your password on TiketKu." />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<meta property="og:title" content="Tiketku - Book Amazing Events" />
	<meta
		property="og:description"
		content="Book concerts, festivals, workshops & sports events in Indonesia"
	/>
	<meta property="og:image" content="/og-tiketku.jpg" />
</svelte:head>

<!-- Left Side Content -->
<LeftSideContent variant="forgot-password" />

<!-- Right Side Content -->
<div class="bg-background relative flex w-full items-center justify-center lg:w-1/2">
	<NavigationHeader />
	<ForgotForm />
</div>
