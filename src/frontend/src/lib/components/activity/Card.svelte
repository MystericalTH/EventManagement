<!-- title date format -->
<script lang="ts">
	import type { Activity } from '$lib/types';
	import { formatActivityDateTime } from '$lib/utils/activity';

	export let data: Activity;
	export let onclick: (data: Activity) => void;
</script>

<button
	onclick={() => onclick(data)}
	class="space-between bg-whiteshadow-md flex h-48 w-72 flex-none flex-col rounded-lg border border-gray-300 bg-white hover:border-indigo-500 hover:bg-indigo-500 hover:text-white hover:shadow-lg"
>
	<div class="mx-4 w-64 flex-1 border-b border-gray-400 py-4 text-left text-lg font-bold">
		{data.title}
	</div>
	<div
		class="flex w-full items-center justify-between rounded-b-lg bg-white p-4 text-sm text-gray-600"
	>
		<span class="border-r-200 ... w-44 truncate whitespace-nowrap text-left"
			>{formatActivityDateTime(data)}</span
		>
		<span>{data.format}</span>
	</div>
</button>
