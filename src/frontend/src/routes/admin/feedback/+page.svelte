<script lang="ts">
    import { createPagination } from '$lib/utils/pagination.svelte';
	import type { FeedbackData } from '$lib/types/feedback';
    import type { PageData } from './$types';
	let { data }: { data: PageData } = $props();

    let selectedActivity = $state('');
    let feedback: FeedbackData[] = $state([]);
    let pagination = $state(createPagination([], 10));

    $effect(() => {
        if (selectedActivity) {
            fetchFeedback(selectedActivity);
        }
    });

    async function fetchFeedback(activityId: string) {
        try {
            const response = await fetch(`/api/activities/${activityId}/feedback`);
            if (response.ok) {
                feedback = await response.json();
                pagination = createPagination(feedback,10);
            } else if (response.status === 404) {
                console.error('Activity not found');
                feedback = [];
            } else {
                console.error('Failed to fetch feedback');
                feedback = [];
            }
        } catch (error) {
            console.error('Error fetching feedback:', error);
            feedback = [];
        }
    }
</script>

<h1 class="mb-4 text-2xl">Feedback</h1>
<!-- Dropdown to select an activity -->
<div class="mb-4">
	<label for="activity-select" class="block text-sm font-medium text-gray-700">Select Activity:</label>
	<select
		id="activity-select"
		bind:value={selectedActivity}
		class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
	>
        <option value="" disabled>Select an activity</option>
		{#each data.activities as activity}
			<option value={activity.id}>{activity.title}</option>
		{/each}
	</select>
</div>

<!-- Feedback Table -->
<div class="text-sm">
	<div class="flex max-w-[1200px] justify-center">
        <div class="overflow-auto border border-slate-200">
          {#if feedback.length > 0}
            <h2 class="mb-4 text-xl font-bold">Feedback for Activity {selectedActivity}</h2>
            <table class="min-w-[794px] divide-y divide-gray-200">
              <thead class="bg-gray-100">
                <tr>
                  <th class="w-36 px-3 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500">
                    User
                  </th>
                  <th class="w-64 px-3 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500">
                    Feedback
                  </th>
                  <th class="w-48 px-3 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500">
                    Date
                  </th>
                </tr>
              </thead>
              <!-- <tbody class="divide-y divide-gray-200">
                {#each feedback as item}
                  <tr>
                    <td class="h-12 w-36 whitespace-nowrap px-3 py-3 text-xs">{item.memberID || 'Anonymous'}</td>
                    <td class="h-12 w-64 overflow-scroll whitespace-nowrap px-3 py-3 text-xs">{item.feedbackMessage}</td>
                    <td class="h-12 w-48 px-3 py-3 text-xs">{new Date(item.feedbackDateTime).toLocaleDateString()}</td>
                  </tr>
                {/each}
              </tbody> -->
              <tbody class="divide-y divide-gray-200">
                {#each pagination.displayPage() as row}
                    {#key row.feedbackid}
                        <tr>
                            <td class="h-12 w-36 whitespace-nowrap px-3 py-3 text-xs"
                                >{row.fname + ' ' + row.lname}</td
                            >
                            <td class="h-12 w-48 overflow-scroll whitespace-nowrap px-3 py-3 text-xs"
                                >{row.feedbackdateTime}</td
                            >
                            <td
                                class="h-12 min-w-48 max-w-64 overflow-scroll whitespace-nowrap px-3 py-3 text-xs"
                                >{row.feedbackmessage}</td
                            >
                        </tr>
                    {/key}
                {/each}
                {#each { length: pagination.rowsPerPage - pagination.displayPage().length } as _}
                    <tr><td colSpan={6} class="h-12 bg-gray-200"></td></tr>
                {/each}
              </tbody>
            </table>
          {:else if selectedActivity}
            <p class="text-gray-500">No feedback available for this activity.</p>
          {/if}
        </div>
      </div>
      <div class="mt-4 flex flex-row justify-center">
		<button
			onclick={pagination.prevPage}
			class="mx-1 rounded-lg border-2 bg-white px-4 py-2 {pagination.hasPrevPage()
				? 'border-indigo-500 text-indigo-500 md:hover:bg-indigo-500 md:hover:text-white'
				: 'border-gray-100 text-gray-400'}"
		>
			Previous</button
		>
		<input
			type="text"
			bind:value={pagination.pageBuffer}
			onchange={pagination.setPage}
			class="w-8 rounded-lg border-2 border-slate-200 text-center"
		/>
		<button
			onclick={pagination.nextPage}
			class="mx-1 rounded-lg border-2 bg-white px-4 py-2 {pagination.hasNextPage()
				? 'border-indigo-500 text-indigo-500 md:hover:bg-indigo-500 md:hover:text-white'
				: 'border-gray-100 text-gray-400'}">Next</button
		>
	</div>
</div>