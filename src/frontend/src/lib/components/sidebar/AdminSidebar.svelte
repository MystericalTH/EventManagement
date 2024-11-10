<script>
	import Sidebar from '../Sidebar.svelte';
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

<Sidebar items={sidebarItems} />
