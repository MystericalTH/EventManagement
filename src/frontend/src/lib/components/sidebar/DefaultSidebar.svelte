<script lang="ts">
	import BaseSidebar from './BaseSidebar.svelte';
	import type { SidebarItem } from '$lib/types/sidebar';

	let items: SidebarItem[] = [
		{
			text: 'Login',
			subitems: [
				{ text: 'Member', href: '/api/login?role=member' },
				{ text: 'Admin', href: '/api/login?role=admin' }
			]
		}
	];
</script>

<BaseSidebar {items} hasLogout={false} />
