<script lang="ts">
	import RegistrationForm from '$lib/components/auth/RegistrationForm.svelte';
	import NavigationHeader from '$lib/components/auth/NavigationHeader.svelte';
	import LeftSideContent from '$lib/components/auth/LeftSideContent.svelte';
</script>

<svelte:head>
	<title>Sign Up | Create Account for free on tiketku</title>
	<meta name="description" content="Create your account and join TiketKu." />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<meta property="og:title" content="Tiketku - Book Amazing Events" />
	<meta
		property="og:description"
		content="Book concerts, festivals, workshops & sports events in Indonesia"
	/>
	<meta property="og:image" content="/og-tiketku.jpg" />
</svelte:head>

<!-- Left Side Content -->
<LeftSideContent variant="signup" />

<!-- Right Side Content -->
<div class="bg-background relative flex w-full items-center justify-center lg:w-1/2">
	<NavigationHeader />
	<RegistrationForm />
</div>
