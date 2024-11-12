<script lang="ts">
	import BaseSidebar from './BaseSidebar.svelte';
	const items = [
		{
			text: 'Activity',
			subitems: [
				{ text: 'Your Activities', href: '/member/activities' },
				{ text: 'Your Proposals', href: '/member/activities/proposals' }
			]
		},
		{
			text: 'Propose New Activity',
			href: '/member/activities/propose'
		}
	];
</script>

<BaseSidebar {items} />
