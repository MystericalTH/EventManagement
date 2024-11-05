<script lang="ts">
	import '../app.css';
	import Sidebar from '$lib/Sidebar.svelte';
	// Loaded from .env.local, guide covers this
	// step in a moment.
	const isMswEnabled = import.meta.env.VITE_MSW_ENABLED === 'true';
	// Flag to defer rendering of components
	// until certain criteria are met on dev,
	// e.g. MSW init.
	let isReady = !isMswEnabled;

	if (isMswEnabled) {
		import('$msw').then((res) => res.inject()).then(() => (isReady = true));
	}
</script>

<div class="layout">
	<Sidebar items={sidebarItems} />
	<main>
		{#if isReady}
			<slot />
		{/if}
	</main>
</div>

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