<script lang="ts">
	import type { ActivityData } from '$lib/types/activity.ts';
	import ActivitySlider from '$lib/components/activity/ActivitySlider.svelte';

	let { data }: { data: { activities: Array<Activity> } } = $props();

	const currentDate = new Date();

	const upcomingActivities = data.activities.filter((activity) => {
		const startDate = new Date(activity.startDate);
		return startDate > currentDate;
	});

	const ongoingActivities = data.activities.filter((activity) => {
		const startDate = new Date(activity.startDate);
		const endDate = new Date(activity.endDate);
		return startDate <= currentDate && currentDate <= endDate;
	});

	const completedActivities = data.activities.filter((activity) => {
		const endDate = new Date(activity.endDate);
		return endDate < currentDate;
	});
</script>

<div class="overflow-y-scroll">
	<h1 class="mb-6 text-3xl font-bold">Activities</h1>
	<h2 class="mb-2 text-xl">Upcoming</h2>
	<ActivitySlider activities={upcomingActivities} />
	<h2 class="mb-2 text-xl">Ongoing</h2>
	<ActivitySlider activities={ongoingActivities} />
	<h2 class="mb-2 text-xl">Completed</h2>
	<ActivitySlider activities={completedActivities} />
</div>
