<script lang="ts">
	import type { Activity } from '$lib/types/activity.ts';
	import ActivitySlider from '$lib/components/activity/ActivitySlider.svelte';
	import { cross } from '$lib/assets/action-button-icons';
	import Overlay from '$lib/components/Overlay.svelte';
	import ActivityContent from '$lib/components/ActivityContent.svelte';
	import ActionButton from '$lib/components/ActionButton.svelte';

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
	let showOverlay = $state(false);
	let selectedActivity: Activity | null = $state(null);

	function viewActivity(activity: Activity) {
		selectedActivity = activity;
		showOverlay = true;
	}
	const closeOverlay = () => {
		selectedActivity = null;
		showOverlay = false;
	};
</script>

<div class="overflow-y-scroll">
	<h1 class="mb-6 text-3xl font-bold">Activities</h1>
	<h2 class="mb-2 text-xl">Upcoming</h2>
	<ActivitySlider activities={upcomingActivities} cardOnclick={viewActivity} />
	<h2 class="mb-2 text-xl">Ongoing</h2>
	<ActivitySlider activities={ongoingActivities} cardOnclick={viewActivity} />
	<h2 class="mb-2 text-xl">Completed</h2>
	<ActivitySlider activities={completedActivities} cardOnclick={viewActivity} />
</div>

<Overlay {showOverlay}>
	<div class="flex justify-end">
		<ActionButton imgsrc={cross} action={closeOverlay} width="20px" alt="Close" />
	</div>
	<div class="overflow-y-scroll">
		<ActivityContent data={selectedActivity} />
	</div>
</Overlay>
