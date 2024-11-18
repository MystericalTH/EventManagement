<script lang="ts">
	import { Role, getRole } from '$lib/types/role';
	import BaseSidebar from './BaseSidebar.svelte';
	import { adminItems, memberItems, defaultItems } from './__items';
	import { page } from '$app/stores';

	let currentUrl: string = '';
	page.subscribe(($page) => {
		const url = new URL($page.url.href);
		currentUrl = url.pathname;
	});

	export let role: string;
	let enumRole: string = getRole(role);
</script>

{#if enumRole === Role.ADMIN}
	<BaseSidebar items={adminItems} />
{:else if enumRole === Role.MEMBER}
	<BaseSidebar items={memberItems} />
{:else}
	<BaseSidebar items={defaultItems(currentUrl)} hasLogout={false} />
{/if}
