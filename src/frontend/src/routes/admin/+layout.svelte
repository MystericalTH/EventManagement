<script>
	import Sidebar from '$lib/Sidebar.svelte';

	const sidebarItems = [
		{
			text: 'Activity',
			href: null,
			subitems: [
				{ text: 'Manage Activities', href: '/admin/activities', subitems: null },
				{ text: 'Activity Requests', href: '/admin/activities/requests', subitems: null },
				{ text: 'See Feedback', href: '/admin/activities/feedback', subitems: null }
			]
		},
		{
			text: 'Members',
			href: null,
			subitems: [
				{ text: 'Manage Members', href: '/admin/members', subitems: null },
				{ text: 'Member Requests', href: '/admin/members/requests', subitems: null }
			]
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
