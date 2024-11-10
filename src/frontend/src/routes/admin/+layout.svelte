<script lang="ts">
	import Sidebar from '$lib/components/Sidebar.svelte';

	const sidebarItems = [
		{
			text: 'Activities',
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

<div class="flex h-screen">
	<Sidebar items={sidebarItems} />
	<main class="flex-1 bg-gray-100 p-4">
		<slot></slot>
	</main>
</div>
