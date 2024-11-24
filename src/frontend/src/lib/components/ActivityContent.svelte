<script lang="ts">
	import type { Activity } from '$lib/types/activity';
	import { formatActivityDateTime } from '$lib/utils/activity';
	import calendar from '$lib/assets/calendar.png';
	import project from '$lib/assets/project.png';
	import workshop from '$lib/assets/workshop.png';
	import people from '$lib/assets/people.png';
	export let data: Activity | null;
</script>

{#if data != null}
	<div class="text-left">
		<h1 class="text-3xl">{data.title}</h1>
		<p class="mb-4 mt-2 text-sm italic text-gray-400">Submitted by {data.proposer}</p>
		<div>
			<h2 class="mb-2 text-xl">Activity Details</h2>
			<ul class="flex flex-col space-y-2">
				<li>
					<span class="inline-flex items-center align-middle"
						><img
							src={calendar}
							width="20px"
							alt="calendar"
							class="mr-[15px]"
						/>{formatActivityDateTime(data)}</span
					>
				</li>
				<li>
					<span class="inline-flex items-center text-center align-middle"
						><img
							src={data.format.toLowerCase() == 'project' ? project : workshop}
							alt={data.format.toLowerCase() == 'project' ? 'project' : 'workshop'}
							width="20px"
							class="mr-[15px]"
						/>{data.format}</span
					>
				</li>
				<li>
					<span class="inline-flex items-center text-center align-middle"
						><img
							src={people}
							width="20px"
							alt="participant"
							class="mr-[15px]"
						/>{data.maxParticipant}
					</span>
				</li>
			</ul>
		</div>
		<div class="my-4">
			<h2 class="mb-2 text-xl">Description</h2>
			<p>{data.description}</p>
		</div>
		<div class="my-2 flex flex-row items-center justify-between text-xl">Recruting Roles</div>
		<div
			class="align-center flex flex-wrap content-center"
			style="column-gap: 10px; row-gap: 10px;"
		>
			{#each data.activityRoles as role}
				<div class="rounded-lg bg-indigo-500 p-2 text-base text-white">
					{role}
				</div>
			{/each}
		</div>
	</div>
{:else}
	<div class="flex items-center justify-center text-base text-gray-300">Nothing to see here</div>
{/if}
