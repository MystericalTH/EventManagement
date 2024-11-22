<script lang="ts">
	import { createPagination } from '$lib/utils/pagination.svelte';
	import { rejectMemberRequest, approveMemberRequest } from '$lib/utils/adminActions';
	import type { Member, Pagination } from '$lib/types';
	let { data } = $props();

	import ActionButton from '$lib/components/ActionButton.svelte';
	import { checkCircle, trash } from '$lib/assets/action-button-icons';
	let pagination: Pagination<Member> = createPagination<Member>(data.memberRequestList, 10);
</script>

<h1 class="mb-4 text-2xl">Member Requests</h1>
<div class="mb-4 content-center">
	Requests remaining <div
		class="inline-block h-7 w-7 content-center rounded-lg {pagination.count <= 0
			? 'bg-gray-200 text-gray-400'
			: 'bg-indigo-500 text-white'} text-center align-middle"
	>
		{pagination.count}
	</div>
</div>
<div class="flex justify-center">
	<div class="overflow-auto border border-slate-200">
		<table class="min-w-[794px] divide-y divide-gray-200">
			<thead class="bg-gray-100">
				<tr>
					<th
						class="w-36 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
						>Name</th
					>
					<th
						class="w-48 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
						>Interest</th
					>
					<th
						class="w-48 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
						>Email</th
					>
					<th
						class="w-40 min-w-36 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
						>Phone Number</th
					>
					<th
						class="w-64 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
						>Reason to join</th
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
							<td class="h-12 w-36 whitespace-nowrap px-3 py-3 text-sm"
								>{row.fname + ' ' + row.lname}</td
							>
							<!--<td class="h-12 max-w-48 px-3 text-sm">
									<div class="flex content-center space-x-2">
										{#each row.interests as interest}
											<div class="rounded-md bg-indigo-300 p-1 text-white">{interest}</div>
										{/each}
									</div>
                </td>-->
							<td class="h-12 w-36 whitespace-nowrap px-3 py-3 text-sm">{row.interest}</td>
							<td class="h-12 w-48 overflow-scroll whitespace-nowrap px-3 py-3 text-sm"
								>{row.email}</td
							>
							<td class="h-12 w-40 px-3 py-3 text-sm">{row.phone}</td>
							<td class="h-12 min-w-48 max-w-64 overflow-scroll whitespace-nowrap px-3 py-3 text-sm"
								>{row.reason}</td
							>
							<td>
								<div class="flex justify-center space-x-2">
									<ActionButton
										imgsrc={checkCircle}
										action={() => approveMemberRequest(row.id, pagination)}
										alt={`Approve request ${row.id}`}
										width={'20px'}
									></ActionButton>
									<ActionButton
										imgsrc={trash}
										action={() => rejectMemberRequest(row.id, pagination)}
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
