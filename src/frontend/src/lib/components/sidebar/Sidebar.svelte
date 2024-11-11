<script lang="ts">
	import { Role, getRole } from '$lib/types/role';
	import AdminSidebar from './AdminSidebar.svelte';
	import DefaultSidebar from './DefaultSidebar.svelte';
	import MemberSidebar from './MemberSidebar.svelte';
	export let role: string;
	let enumRole = getRole(role);
	console.log(role);
</script>

{#if enumRole === Role.Admin}
	<AdminSidebar />
{:else if enumRole === Role.Member}
	<MemberSidebar />
{:else}
	<DefaultSidebar />
{/if}
