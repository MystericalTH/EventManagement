<script lang="ts">
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
	let users = data.item;
	let headerClass: string =
		'px-3 py-3 text-left text-xs  font-medium uppercase tracking-wider text-gray-500';
	let itemClass: string = 'whitespace-nowrap px-3 py-3 text-xs';
</script>

<h1 class="mb-4 font-custom text-2xl">Manage Members</h1>
<table class="min-w-full divide-y divide-gray-200 font-custom">
	<thead class="bg-gray-50">
		<tr>
			<th class={headerClass}>Name</th>
			<th class={headerClass}>Email</th>
			<th class={headerClass}>Test</th>
		</tr>
	</thead>
	<tbody class="divide-y divide-gray-200 bg-white">
		{#each users as user}
			<tr>
				<td class={itemClass}>{user.name}</td>
				<td class={itemClass}>{user.email}</td>
				<td class={itemClass}>Data 3</td>
			</tr>
		{/each}
	</tbody>
</table>
