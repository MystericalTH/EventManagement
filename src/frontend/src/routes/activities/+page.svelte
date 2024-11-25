<script lang="ts">
	import type { Activity } from '$lib/types';
	import ActivitySlider from '$lib/components/activity/ActivitySlider.svelte';
	import { goto } from '$app/navigation';

	let { data }: { data: { activities: Array<Activity> } } = $props();

	const currentDate = new Date();

	const upcomingActivities = data.activities.filter((activity: Activity) => {
		const startDate = new Date(activity.startDate);
		return startDate > currentDate;
	});

	const ongoingActivities = data.activities.filter((activity: Activity) => {
		const startDate = new Date(activity.startDate);
		const endDate = new Date(activity.endDate);
		return startDate <= currentDate && currentDate <= endDate;
	});

	const completedActivities = data.activities.filter((activity: Activity) => {
		const endDate = new Date(activity.endDate);
		return endDate < currentDate;
	});

	let selectedActivity: Activity | null = $state(null);

	function gotoActivity(activity: Activity) {
		selectedActivity = activity;
		goto(`/activities/${activity.id}`);
	}
</script>

<div class="overflow-y-scroll">
	<h1 class="mb-6 text-3xl font-bold">Activities</h1>
	<h2 class="mb-2 text-xl">Upcoming</h2>
	<ActivitySlider activities={upcomingActivities} cardOnclick={gotoActivity} />
	<h2 class="mb-2 text-xl">Ongoing</h2>
	<ActivitySlider activities={ongoingActivities} cardOnclick={gotoActivity} />
	<h2 class="mb-2 text-xl">Completed</h2>
	<ActivitySlider activities={completedActivities} cardOnclick={gotoActivity} />
</div>
