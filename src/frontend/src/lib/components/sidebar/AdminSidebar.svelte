<script lang="ts">
	import type { SidebarItem } from '$lib/types/sidebar';
	import BaseSidebar from './BaseSidebar.svelte';
	const items: SidebarItem[] = [
		{
			text: 'Activity',
			subitems: [
				{ text: 'Manage Activities', href: '/admin/activities' },
				{ text: 'Activity Requests', href: '/admin/activities/requests' },
				{ text: 'See Feedback', href: '/admin/activities/feedback' }
			]
		},
		{
			text: 'Members',
			subitems: [
				{ text: 'Manage Members', href: '/admin/members' },
				{ text: 'Member Requests', href: '/admin/members/requests' }
			]
		}
	];
</script>

<BaseSidebar {items} />
