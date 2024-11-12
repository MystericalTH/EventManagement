<script lang="ts">
	import { Role, getRole } from '$lib/types/role';
	import BaseSidebar from './BaseSidebar.svelte';
	import { adminItems, memberItems, defaultItems } from './__items';
	export let role: string;
	let enumRole = getRole(role);
	console.log(role);
</script>

{#if enumRole === Role.Admin}
	<BaseSidebar items={adminItems} />
{:else if enumRole === Role.Member}
	<BaseSidebar items={memberItems} />
{:else}
	<BaseSidebar items={defaultItems} />
{/if}
