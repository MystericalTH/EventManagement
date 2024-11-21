<script lang="ts">
	import moment from 'moment';
	import Card from './Card.svelte';
	import type { ActivityData } from '$lib/types/activity';
	export let activities: ActivityData[];
</script>

<div class="mb-4 flex h-56 flex-row items-center space-x-16 overflow-x-scroll bg-gray-100 px-4">
	{#if activities.length != 0}
		{#each activities as activity}
			<Card
				title={activity.title}
				startDate={moment(activity.startDate).format('D/M/YYYY')}
				endDate={moment(activity.endDate).format('D/M/YYYY')}
				format={activity.format}
				href={`activities/${activity.id.toString()}`}
			/>
		{/each}
	{:else}
		<div class="flex-1 text-center">
			<p class="text-gray-400">Nothing to see here</p>
		</div>
	{/if}
</div>
