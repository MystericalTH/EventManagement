<script lang="ts">
	import { createPagination } from '$lib/utils/pagination.svelte';
	import ActionButton from '$lib/components/ActionButton.svelte';
	import ActivityContent from '$lib/components/ActivityContent.svelte';
	import type { Pagination, Activity } from '$lib/types';
	let { data }: { data: { activities: Activity[] } } = $props();
	import { trash, checkCircle, view, cross } from '$lib/assets/action-button-icons';
	import { rejectActivityRequest, approveActivityRequest } from '$lib/utils/adminActions.js';
	import { formatActivityDateTime } from '$lib/utils/activity';
	import Overlay from '$lib/components/Overlay.svelte';
	let pagination: Pagination<Activity> = createPagination<Activity>(data.activities, 10);

	let showOverlay = $state(false);
	let selectedActivity: Activity | null = $state(null);

	function viewActivity(activity: Activity) {
		selectedActivity = activity;
		showOverlay = true;
	}
	const closeOverlay = () => {
		selectedActivity = null;
		showOverlay = false;
	};
</script>

<h1 class="mb-4 font-custom text-2xl">Activity Requests</h1>
<div class="text-base">
	<div class="flex justify-center">
		<div class="overflow-auto border border-slate-200">
			<table class="min-w-[794px] divide-y divide-gray-200">
				<thead class="bg-gray-100">
					<tr>
						<th
							class="w-36 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
							>Title</th
						>
						<th
							class="w-24 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
							>Format</th
						>
						<th
							class="w-56 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
							>Date</th
						>
						<th
							class="w-20 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
							>Estimated Participants</th
						>
						<th
							class="w-8 px-3 py-3 text-center text-sm font-medium uppercase tracking-wider text-gray-500"
							>Actions</th
						>
					</tr>
				</thead>
				<tbody class="divide-y divide-gray-200">
					{#each pagination.displayPage() as row}
						{#key row.id}
							<tr>
								<td class="h-12 w-36 whitespace-nowrap px-3 py-3 text-sm">{row.title}</td>
								<td class="h-12 w-24 whitespace-nowrap px-3 py-3 text-sm">{row.format}</td>
								<td class="h-12 w-56 px-3 py-3 text-sm">{formatActivityDateTime(row)}</td>
								<td class="h-12 w-20 whitespace-nowrap px-3 py-3 text-sm">{row.maxParticipant}</td>
								<td>
									<div class="flex justify-center space-x-2">
										<ActionButton
											imgsrc={view}
											action={() => viewActivity(row)}
											alt={`View request ${row.id}`}
											width={'20px'}
										></ActionButton>
										<ActionButton
											imgsrc={checkCircle}
											action={() => approveActivityRequest(row.id, pagination)}
											alt={`Approve request ${row.id}`}
											width={'20px'}
										></ActionButton>
										<ActionButton
											imgsrc={trash}
											action={() => rejectActivityRequest(row.id, pagination)}
											alt={`Reject request ${row.id}`}
											width={'20px'}
										></ActionButton>
									</div>
								</td>
							</tr>
						{/key}
					{/each}
					{#each { length: pagination.rowsPerPage - pagination.displayPage().length } as _}
						<tr><td colSpan={6} class="h-12 bg-gray-200"></td></tr>
					{/each}
				</tbody>
			</table>
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
<Overlay {showOverlay}>
	<div class="flex justify-end">
		<ActionButton imgsrc={cross} action={closeOverlay} width="20px" alt="Close" />
	</div>
	<div class="overflow-y-scroll">
		<ActivityContent data={selectedActivity} />
	</div>
</Overlay>
