<script lang="ts">
	import Card from './Card.svelte';
	import type { Activity } from '$lib/types/activity';
	export let activities: Activity[];
	export let cardOnclick: (data: Activity) => void;
</script>

<div class="mb-4 flex h-56 flex-row items-center space-x-16 overflow-x-scroll bg-gray-100 px-4">
	{#if activities.length != 0}
		{#each activities as data}
			<Card {data} onclick={cardOnclick} />
		{/each}
	{:else}
		<div class="flex-1 text-center">
			<p class="text-gray-400">Nothing to see here</p>
		</div>
	{/if}
</div>
