<script lang="ts">
	import * as Card from '$lib/components/ui/card/index.js';
	import Button from '$lib/components/ui/button/button.svelte';
	import FormInput from '$lib/components/forms/FormInput.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
</script>

<svelte:head>
	<title>Sign Up | MyBrand</title>
	<meta
		name="description"
		content="Create a new account with MyBrand to manage your events and tickets."
	/>
</svelte:head>

<section class="bg-accent flex h-[70vh] items-center justify-center">
	<Card.Root class="w-full max-w-sm shadow-none ">
		<Card.Header class="text-center">
			<Card.Title class="text-2xl">Welcome Back</Card.Title>
			<Card.Description>Sign in into your account to start</Card.Description>
		</Card.Header>
		<Card.Content>
			<form method="POST" action="?login" class="space-y-4">
				<Input type="email" name="email" placeholder="johndoe@example.com" />
				<Input type="password" name="password" placeholder="**********" />
				<Button type="submit" class="w-full">Signin</Button>
			</form>

			<div class="my-4 flex items-center gap-4">
				<div class="flex-1 border-t border-gray-300"></div>
				<span class="text-sm text-gray-500">or continue with</span>
				<div class="flex-1 border-t border-gray-300"></div>
			</div>
			<div class="flex gap-2">
				<Button
					variant="outline"
					href="/signin/google"
					class="flex w-1/2 items-center justify-center gap-2"
				>
					<!-- Google Icon -->
					<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" class="h-5 w-5">
						<path
							fill="#fbc02d"
							d="M43.6 20.5h-1.9V20H24v8h11.3C33.9 32.5 29.4 35 24 35c-6.1 0-11.3-5-11.3-11S17.9 13 24 13c2.9 0 5.6 1.1 7.6 2.9l5.7-5.7C33.1 6.3 28.8 4.5 24 4.5 13.8 4.5 5.5 12.9 5.5 23S13.8 41.5 24 41.5c9.9 0 18.1-7.1 18.1-18 0-1.4-.2-2.3-.5-3z"
						/>
						<path
							fill="#e53935"
							d="M6.3 14.7l6.6 4.8C14.3 16.1 18.7 13 24 13c2.9 0 5.6 1.1 7.6 2.9l5.7-5.7C33.1 6.3 28.8 4.5 24 4.5c-7.2 0-13.4 4-16.7 9.9z"
						/>
						<path
							fill="#4caf50"
							d="M24 41.5c5.2 0 10.1-1.9 13.8-5.1l-6.4-5.4c-2.2 1.5-5 2.4-7.4 2.4-5.5 0-10.2-3.6-11.8-8.4l-6.6 5.1C10.6 37.8 16.9 41.5 24 41.5z"
						/>
						<path
							fill="#1565c0"
							d="M43.6 20.5h-1.9V20H24v8h11.3c-1.3 3.8-5.5 7-11.3 7-5.5 0-10.2-3.6-11.8-8.4l-6.6 5.1C10.6 37.8 16.9 41.5 24 41.5c9.9 0 18.1-7.1 18.1-18 0-1.4-.2-2.3-.5-3z"
						/>
					</svg>
					Google
				</Button>

				<Button
					variant="outline"
					href="/signin/github"
					class="flex w-1/2 items-center justify-center gap-2"
				>
					<!-- GitHub Icon -->
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="h-5 w-5"
						fill="currentColor"
						viewBox="0 0 24 24"
					>
						<path
							d="M12 0C5.37 0 0 5.373 0 12c0 5.303 3.438 9.8 8.207 11.387.6.113.793-.258.793-.577v-2.257c-3.338.724-4.033-1.416-4.033-1.416-.546-1.385-1.333-1.754-1.333-1.754-1.089-.745.083-.729.083-.729 1.205.084 1.84 1.238 1.84 1.238 1.07 1.835 2.807 1.305 3.492.997.107-.775.418-1.305.76-1.605-2.665-.304-5.466-1.336-5.466-5.931 0-1.31.468-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.3 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.803 5.624-5.475 5.921.43.372.823 1.102.823 2.222v3.293c0 .322.192.694.8.576C20.565 21.796 24 17.3 24 12c0-6.627-5.373-12-12-12z"
						/>
					</svg>
					GitHub
				</Button>
			</div>
		</Card.Content>
		<Card.Footer>
			<p>
				Dont have an account ? <a href="/signup" class="text-blue-600 hover:underline">Sign up</a>
			</p>
		</Card.Footer>
	</Card.Root>
</section>
