<script lang="ts">
	export let showOverlay: boolean;
</script>

{#if showOverlay}
	<div class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
		<div class="h-3/4 w-2/5 max-w-lg rounded-lg bg-white p-6 shadow-lg"><slot /></div>
	</div>
{/if}
