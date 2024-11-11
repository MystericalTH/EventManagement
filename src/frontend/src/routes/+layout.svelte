<script lang="ts">
	import '../app.css';
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
	<main>
		{#if isReady}
			<slot />
		{/if}
	</main>
</div>
