<script>
	import Sidebar from '$lib/Sidebar.svelte';

	const sidebarItems = [
		{
			text: 'Activity',
			href: null,
			subitems: [
				{ text: 'Your Activities', href: '/member/activities', subitems: null },
				{ text: 'Your Proposals', href: '/member/activities/proposals', subitems: null }
			]
		},
		{
			text: 'Propose New Activity',
			href: '/member/activities/propose',
			subitems: null
		}
	];
</script>

<div class="layout">
	<Sidebar items={sidebarItems} />
	<main>
		<slot></slot>
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
