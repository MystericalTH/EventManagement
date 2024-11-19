<script lang="ts">
	import type { ActivityData } from '$lib/types/activity.ts';

	export let data: {
		activity: ActivityData;
		isRegistered: boolean;
		isEventPast: boolean;
		hasSubmittedFeedback: boolean;	
		nextActivityId: number | null;
	};

	const openRegisterPage = () => {
		window.location.href = `/activity/${data.activity.id}/register`;
	};

	const openFeedbackPage = () => {
		window.location.href = `/activity/${data.activity.id}/feedback`;
	};

	const navigateToActivity = (offset: number) => {
		const currentId = data.activity.id;
		const newId = currentId + offset;
		window.location.href = `/activity/${newId}`;
	};
</script>

<div class="mx-auto max-w-md items-center text-center">
	<h1 class="my-10 text-2xl font-semibold">{data.activity.title}</h1>
	<div class="flex items-center space-x-72">
		<span class="material-icons align-middle font-medium leading-none text-gray-600"
			>{data.activity.startDate}</span
		>
		<span class="align-middle font-medium leading-none">{data.activity.format}</span>
	</div>
	<div class="my-10 text-lg">
		<p>{data.activity.description}</p>
	</div>	
	
	{#if data.isEventPast && data.isRegistered}
		{#if data.hasSubmittedFeedback}
			<button class="bg-gray-500 text-white py-2 px-4 rounded" disabled>Feedback Submitted</button>
		{:else}
			<button class="bg-green-500 text-white py-2 px-4 rounded" on:click={openFeedbackPage}>Submit Feedback</button>
		{/if}
	{:else}
		{#if data.isRegistered}
			<button class="bg-gray-500 text-white py-2 px-4 rounded" disabled>Registered</button>
		{:else}
			<button class="bg-blue-500 text-white py-2 px-4 rounded" on:click={openRegisterPage}>Register</button>
		{/if}
	{/if}
    
	<div class="mt-10 flex justify-between">
		<button
			class="rounded bg-gray-300 px-4 py-2 text-black"
			on:click={() => navigateToActivity(-1)}
			style="visibility: {data.activity.id > 1 ? 'visible' : 'hidden'}"
		>
			Previous
		</button>
		<button
			class="rounded bg-gray-300 px-4 py-2 text-black"
			on:click={() => navigateToActivity(1)}
			style="visibility: {data.nextActivityId !== null ? 'visible' : 'hidden'}"
		>
			Next
		</button>
	</div>
</div>
