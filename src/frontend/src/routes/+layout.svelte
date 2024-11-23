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
		<div class="flex w-[calc(100vw-12rem)] flex-none justify-center p-4">
			<div class="flex w-full flex-col overflow-x-clip">
				{@render children()}
			</div>
		</div>
	</div>
</main>
