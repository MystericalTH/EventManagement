<!-- title date format -->
<script>
	export let title;
	export let startDate;
	export let endDate;
	export let format;
	export let href;
</script>

<a
	{href}
	class="space-between bg-whiteshadow-md flex h-48 w-72 flex-none flex-col rounded-lg border border-gray-300 bg-white hover:border-indigo-500 hover:bg-indigo-500 hover:text-white"
>
	<div class="mx-4 flex-1 border-b border-gray-400 py-4 text-lg font-bold">{title}</div>
	<div
		class="flex w-full items-center justify-between rounded-b-lg bg-white p-4 text-sm text-gray-600"
	>
		<span>
			<span>{startDate}</span>
			{#if endDate != startDate}
				<span> - {endDate}</span>
			{/if}
		</span>

		<span>{format}</span>
	</div>
</a>
