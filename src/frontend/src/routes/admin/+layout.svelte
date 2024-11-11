<script>
	import AdminSidebar from '$lib/components/sidebar/AdminSidebar.svelte';
</script>

<div class="layout">
	<AdminSidebar></AdminSidebar>
	<main>
		<slot></slot>
	</main>
</div>
