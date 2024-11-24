<script lang="ts">
	import type { PageData } from './$types';
	import type { Member, Pagination } from '$lib/types';

	let { data }: { data: PageData } = $props();

	import ActionButton from '$lib/components/ActionButton.svelte';
	import { createPagination } from '$lib/utils/pagination.svelte';
	let pagination: Pagination<Member> = createPagination<Member>(data.memberList, 10);
	import { removeMember } from '$lib/utils/adminActions';
	import { edit, trash } from '$lib/assets/action-button-icons';
</script>

<h1 class="mb-4 font-custom text-2xl">Manage Members</h1>
<div class="flex justify-center">
	<div class="overflow-auto border border-slate-200">
		<table class="min-w-[794px] divide-y divide-gray-200 font-custom">
			<thead class="bg-gray-50">
				<tr>
					<th
						class="w-36 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
						>Name</th
					>
					<th
						class="w-36 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
						>Interest</th
					>
					<th
						class="w-48 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
						>Email</th
					>

					<th
						class="w-36 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
						>Phone Number</th
					>

					<th
						class="w-36 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
						>Github URL</th
					>
					<th
						class="w-20 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
						>Actions</th
					>
				</tr>
			</thead>
			<tbody class="divide-y divide-gray-200 bg-white">
				{#each pagination.displayPage() as row}
					<tr>
						<td class="h-12 w-36 whitespace-nowrap px-3 py-3 text-sm"
							>{row.fname + ' ' + row.lname}</td
						>
						<td class="h-12 w-36 whitespace-nowrap px-3 py-3 text-sm">{row.interest}</td>
						<td class="h-12 w-48 whitespace-nowrap px-3 py-3 text-sm">{row.email}</td>
						<td class="h-12 min-w-36 whitespace-nowrap px-3 py-3 text-sm">{row.phone}</td>
						<td class="h-12 min-w-36 whitespace-nowrap px-3 py-3 text-sm">{row.githuburl}</td>
						<td>
							<div class="flex justify-center space-x-2">
								<ActionButton
									imgsrc={edit}
									action={() => {}}
									alt={`Edit member ${row.memberid}`}
									width={'20px'}
								></ActionButton>
								<ActionButton
									imgsrc={trash}
									action={() => removeMember(row.memberid, pagination)}
									alt={`Remove member ${row.memberid}`}
									width={'20px'}
								></ActionButton>
							</div></td
						>
					</tr>
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
