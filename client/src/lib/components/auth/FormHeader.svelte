<script lang="ts">
	export let title: string = 'Welcome to Our Tiketku';
	export let subtitle: string = 'Please sign in to continue';
</script>

<div class="mb-8 text-center lg:text-left">
	<h2
		class="mb-2 bg-gradient-to-r from-blue-500 to-indigo-500 bg-clip-text text-3xl font-bold text-transparent"
	>
		{title}
	</h2>
	<p class="text-muted-foreground">{subtitle}</p>
</div>
