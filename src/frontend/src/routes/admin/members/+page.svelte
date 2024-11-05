<script lang="ts">
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
	let users = data.item;
	console.log(data);
</script>

<table>
	<thead>
		<tr>
			<th>Name</th>
			<th>Email</th>
		</tr>
	</thead>
	<tbody>
		{#each users as user}
			<tr>
				<td>{user.name}</td>
				<td>{user.email}</td>
			</tr>
		{/each}
	</tbody>
</table>
