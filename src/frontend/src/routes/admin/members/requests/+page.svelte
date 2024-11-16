<script lang="ts">
	import { createPagination } from '$lib/utils/pagination.svelte';
	import { rejectRequest } from '$lib/utils/memberRequest';
	import eye from '$lib/assets/eye.svg';
	import bin from '$lib/assets/bin.svg';
	import { adminState } from '$lib/states/states.svelte';
	let pagination = createPagination(adminState.memberRequestList, 10);
</script>

<h1 class="mb-4 text-2xl">Member Requests</h1>
<div class="text-sm">
	<div class="max-w-144 overflow-scroll border border-slate-200">
		<table class="table-fixed divide-y divide-gray-200 border border-slate-200">
			<thead class="bg-gray-100">
				<tr>
					<th
						class="w-36 px-3 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
						>Name</th
					>
					<th
						class="w-20 px-3 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
						>Interest</th
					>
					<th
						class="w-36 px-3 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
						>Email</th
					>
					<th
						class="w-44 min-w-36 px-3 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
						>Phone Number</th
					>
					<th
						class="min-w-56 px-3 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
						>Reason to join</th
					>
					<th
						class="w-8 px-3 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
						>Action</th
					>
				</tr>
			</thead>
			<tbody class="divide-y divide-gray-200">
				{#each pagination.displayPage() as row}
					<tr>
						<td class="h-12 whitespace-nowrap px-3 py-3 text-xs"
							>{row.firstName + ' ' + row.lastName}</td
						>
						<td class="flex h-12 flex-row space-x-2 px-3 py-3 text-xs">
							{#each row.interests as interest}
								<div class="rounded-md bg-indigo-300 p-1 text-white">{interest}</div>
							{/each}
						</td>
						<td class="h-12 whitespace-nowrap px-3 py-3 text-xs">{row.email}</td>
						<td class="h-12 whitespace-nowrap px-3 py-3 text-xs">{row.phoneNumber}</td>
						<td class="h-12 whitespace-pre-wrap px-3 py-3 text-xs">{row.reason}</td>
						<td>
							<div class="flex justify-center space-x-2">
								<button>
									<img src={eye} alt="See requests" width="20px" />
								</button>
								<button onclick={() => rejectRequest(row.id)}>
									<img src={bin} alt="Remove requests" width="20px" />
								</button>
							</div>
						</td>
					</tr>
				{/each}
				{#each { length: pagination.rowsPerPage - pagination.displayPage().length } as _}
					<tr
						><td
							colSpan={5}
							class="h-12 whitespace-nowrap bg-gray-200 px-3 py-3 text-xs text-gray-200">.</td
						></tr
					>
				{/each}
			</tbody>
		</table>
	</div>

	<div class="mt-4 flex flex-row justify-center">
		<button onclick={pagination.prevPage} class="mx-1 rounded-lg bg-indigo-500 px-4 py-2 text-white"
			>Previous</button
		>
		<input
			type="text"
			bind:value={pagination.pageBuffer}
			onchange={() => {
				pagination.currentPage = pagination.pageBuffer;
			}}
			class="w-8 rounded-lg border-2 border-slate-200 text-center"
		/>
		<button onclick={pagination.nextPage} class="mx-1 rounded-lg bg-indigo-500 px-4 py-2 text-white"
			>Next</button
		>
	</div>
</div>
