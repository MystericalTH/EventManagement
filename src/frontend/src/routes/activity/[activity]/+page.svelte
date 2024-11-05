<script lang="ts">
	import type { ActivityData } from '../../../lib/types/activity.ts';

	export let data: { activity: ActivityData };
</script>

<div class="items-center text-center mx-auto max-w-md">
	<h1 class="text-2xl font-semibold my-10">{data.activity.title}</h1>
	<div class="flex items-center space-x-72">
		<span class="material-icons text-gray-600 font-medium leading-none align-middle">{data.activity.startDate}</span>
		<span class="font-medium leading-none align-middle">{data.activity.format}</span>
	</div>
	<div class="mt-10 text-lg">
		<p>{data.activity.description}</p>
	</div>
</div>