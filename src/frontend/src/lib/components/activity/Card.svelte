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
	class="block h-48 w-72 flex-none rounded-lg border border-gray-300 p-4 shadow-md hover:bg-gray-100"
>
	<div class="mb-4 text-lg font-bold">{title}</div>
	<div class="flex w-full items-center justify-between text-sm text-gray-600">
		<span>
			<span>{startDate}</span>
			{#if endDate != startDate}
				<span> - {endDate}</span>
			{/if}
		</span>

		<span>{format}</span>
	</div>
</a>
