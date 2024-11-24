<script lang="ts">
	export let showOverlay: boolean;
	export let width: string = '';
	export let height: string = '';
</script>

{#if showOverlay}
	<div class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
		<div
			class={`h-${height} w-${width} max-w-lg rounded-lg bg-white p-6 shadow-lg ${
				width != '' ? `w-${width} ` : 'w-2/5 '
			} ${height != '' ? `h-${height}` : 'h-3/4'}`}
		>
			<slot></slot>
		</div>
	</div>{/if}
