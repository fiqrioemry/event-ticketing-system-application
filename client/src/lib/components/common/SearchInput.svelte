<!-- components/ui/SearchInput.svelte -->
<script lang="ts">
	import { Input } from '$lib/components/ui/input';
	export let search;
	export let setSearch: (val: string) => void;
	export let setPage: (val: number) => void;
	export let placeholder: string;

	function handleInput(e: Event) {
		const val = (e.target as HTMLInputElement).value;
		setPage(1);
		setSearch(val);
	}
</script>

<Input value={search} class="md:w-1/2" onfocus={handleInput} {placeholder} />
