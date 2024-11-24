<script lang="ts">
	import type { PageData } from './$types';
	let { data }: { data: PageData } = $props();

	import ActivityContent from '$lib/components/ActivityContent.svelte';

	console.log(data);
	
	const openRegisterPage = () => {
		window.location.href = `/activities/${data.activity.id}/register`;
	};

	const openFeedbackPage = () => {
		window.location.href = `/activities/${data.activity.id}/feedback`;
	};

	// const navigateToActivity = (offset: number) => {
	// 	const currentId = data.activity.id;
	// 	const newId = currentId + offset;
	// 	window.location.href = `/activities/${newId}`;
	// };
</script>

<div class="flex justify-center">
	{#if data.activity}
		<div class="w-5/6 items-center text-center">
			<ActivityContent data={data.activity} />
			{#if data.isEventPast && data.isRegistered.isRegistered}
				{#if data.hasSubmittedFeedback}
					<button class="rounded bg-gray-500 px-4 py-2 text-white" disabled
						>Feedback Submitted</button
					>
				{:else}
					<button class="rounded bg-green-500 px-4 py-2 text-white" onclick={openFeedbackPage}
						>Submit Feedback</button
					>
				{/if}
			{:else if data.isRegistered.isRegistered}
				<button class="rounded bg-gray-500 px-4 py-2 text-white" disabled>Registered</button>
			{:else}
				<button class="rounded bg-blue-500 px-4 py-2 text-white" onclick={openRegisterPage}
					>Register</button
				>
			{/if}

			<div class="mt-10 flex justify-between">
				<a
					href={`/activities/${data.activity.id - 1}`}
					class="rounded bg-gray-300 px-4 py-2 text-black"
					style="visibility: {data.activity.id > 1 ? 'visible' : 'hidden'}"
				>
					Previous
				</a>
				<a
					href={`/activities/${data.activity.id + 1}`}
					class="rounded bg-gray-300 px-4 py-2 text-black"
					style="visibility: {data.nextActivityId !== null ? 'visible' : 'hidden'}"
				>
					Next
				</a>
			</div>
		</div>
	{/if}
</div>
