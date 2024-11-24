<script lang="ts">
	export let showOverlay: boolean;
	export let width: string = '2/5';
	export let height: string = '3/4';
</script>

{#if showOverlay}
	<div class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
		<div class={`h-${height} w-${width} max-w-lg rounded-lg bg-white p-6 shadow-lg`}><slot /></div>
	</div>
{/if}
