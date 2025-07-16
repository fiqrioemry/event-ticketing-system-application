<script lang="ts">
	export let buttonText: string = 'Sign in';
	export let description: string = 'Already have an account?';
	export let onClick: () => void = () => {};
</script>

<div class="mt-8 text-center">
	<p class="text-muted-foreground text-sm">
		{description}
		<button
			on:click={onClick}
			class="font-medium text-blue-500 transition-colors hover:text-blue-500/80"
		>
			{buttonText}
		</button>
	</p>
</div>
