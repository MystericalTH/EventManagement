<script>
	import '../app.css';
	import Sidebar from '$lib/components/sidebar/Sidebar.svelte';
	// Loaded from .env.local, guide covers this
	// step in a moment.
	let { data, children } = $props();
</script>

<main class="font-custom">
	<div class="flex h-screen">
		<Sidebar role={data.role} />
		<div class="m-4 flex flex-1 justify-center">
			<div class="flex w-5/6 max-w-[720px] flex-col">
				{@render children()}
			</div>
		</div>
	</div>
</main>
