<script lang="ts">
	let { data } = $props();

	import { createPagination } from '$lib/runes/pagination.svelte';
	let pagination = createPagination(data.requestList, 10);
	import { tableDataClass, tableHeaderClass } from '$lib/styles/table';
	import eye from '$lib/assets/eye.svg';
	import bin from '$lib/assets/bin.svg';
</script>

<h1 class="mb-4 text-2xl">Member Requests</h1>
<div class="text-sm">
	<div class="max-w-144 overflow-scroll">
		<table class="min-w-full table-fixed divide-y divide-gray-200 border border-slate-200">
			<thead class="bg-gray-100">
				<tr>
					<th class={tableHeaderClass + ' w-36'}>Name</th>
					<th class={tableHeaderClass}>Interest</th>
					<th class={tableHeaderClass + ' w-36'}>Email</th>
					<th class={tableHeaderClass + ' w-36'}>Phone Number</th>
					<th class={tableHeaderClass + ' w-8'}>Action</th>
				</tr>
			</thead>
			<tbody class="divide-y divide-gray-200">
				{#each pagination.displayPage() as row}
					<tr>
						<td class={tableDataClass}>{row.name}</td>
						<td class={tableDataClass}>{row.interest}</td>
						<td class={tableDataClass}>{row.email}</td>
						<td class={tableDataClass}>{row.phone}</td>
						<td>
							<div class="flex justify-center space-x-2">
								<button>
									<img src={eye} alt="See requests" width="20px" />
								</button>
								<button>
									<img src={bin} alt="Remove requests" width="20px" />
								</button>
							</div>
						</td>
					</tr>
				{/each}
				{#each { length: pagination.rowsPerPage - pagination.displayPage().length } as _, i}
					<tr><td colSpan={5} class={tableDataClass + ' bg-gray-200 text-gray-200'}>.</td></tr>
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
