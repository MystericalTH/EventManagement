<script>
	import '../app.css';
	import Sidebar from '$lib/components/sidebar/Sidebar.svelte';
	// Loaded from .env.local, guide covers this
	// step in a moment.
	let { data, children } = $props();
</script>

<main>
		<div class="flex h-screen">
			<Sidebar role={data.role} />
			<div class="flex-1">
				{@render children()}
			</div>
		</div>
</main>
