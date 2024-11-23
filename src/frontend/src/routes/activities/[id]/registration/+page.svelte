<script lang="ts">
	import { createPagination } from '$lib/utils/pagination.svelte';
	import type { Activity, MemberRegistration, Pagination } from '$lib/types';
	let { data }: { data: { activity: Activity; registrations: MemberRegistration[] } } = $props();
	let pagination: Pagination<MemberRegistration> = createPagination<MemberRegistration>(
		data.registrations,
		10
	);
</script>

<h1 class="text-2xl">Activity Registration</h1>
<p class="mb-4 mt-2 text-base text-gray-700">{data.activity.title}</p>
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
						class="w-24 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
						>Role</th
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
						class="w-96 px-3 py-3 text-left text-sm font-medium uppercase tracking-wider text-gray-500"
						>Expectation</th
					>
				</tr>
			</thead>
			<tbody class="divide-y divide-gray-200">
				{#each pagination.displayPage() as row}
					{#key row.memberid}
						<tr>
							<td class="h-12 w-36 whitespace-nowrap px-3 py-3 text-sm"
								>{row.fname + ' ' + row.lname}</td
							>
							<td class="h-12 w-24 whitespace-nowrap px-3 py-3 text-sm">{row.role}</td>
							<td class="h-12 w-48 overflow-scroll whitespace-nowrap px-3 py-3 text-sm"
								>{row.email}</td
							>
							<td class="h-12 w-40 px-3 py-3 text-sm">{row.phone}</td>
							<td
								class="h-12 w-[32rem] min-w-96 max-w-64 overflow-scroll whitespace-nowrap px-3 py-3 text-sm"
								>{row.expectation}</td
							>
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
