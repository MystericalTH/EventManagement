<script>
	import '../app.css';
	import Sidebar from '$lib/components/sidebar/Sidebar.svelte';
	// Loaded from .env.local, guide covers this
	// step in a moment.
	let { data, children } = $props();
	console.log(data);
	const isMswEnabled = import.meta.env.VITE_MSW_ENABLED === 'true';
	// Flag to defer rendering of components

	// until certain criteria are met on dev,
	// e.g. MSW init.
	let isReady = $state(!isMswEnabled);
	if (isMswEnabled) {
		console.log(isMswEnabled);
		import('$msw')
			.then((res) => res.inject())
			.then(() => {
				isReady = true;
			});
	}
</script>

<main>
	{#if isReady}
		<div class="flex h-screen">
			<Sidebar role={data.role} />
			<div class="flex-1">
				{@render children()}
			</div>
		</div>
	{/if}
</main>

<style>
	.layout {
		display: flex;
		height: 100vh;
	}
	main {
		flex: 1;
		padding: 1rem;
		background-color: #ecf0f1;
	}
</style>