<script lang="ts">
	import type { Activity } from '$lib/types/activity.ts';
	import ActivitySlider from '$lib/components/activity/ActivitySlider.svelte';
	import ActionButton from '$lib/components/ActionButton.svelte';
	import ActivityContent from '$lib/components/ActivityContent.svelte';
	import Overlay from '$lib/components/Overlay.svelte';
	import { cross } from '$lib/assets/action-button-icons';
	let { data }: { data: { activities: Array<Activity> } } = $props();
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

	const currentDate = new Date();

	const upcomingActivities = data.activities.filter((activity) => {
		const startDate = new Date(activity.startDate);
		console.log(startDate);
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
	<h1 class="mb-6 text-3xl font-bold">Your Engagements</h1>
	<h2 class="mb-2 text-xl">Upcoming</h2>
	<ActivitySlider activities={upcomingActivities} cardOnclick={viewActivity} />
	<h2 class="mb-2 text-xl">Ongoing</h2>
	<ActivitySlider activities={ongoingActivities} cardOnclick={viewActivity} />
	<h2 class="mb-2 text-xl">Completed</h2>
	<ActivitySlider activities={completedActivities} cardOnclick={viewActivity} />
</div>

<Overlay {showOverlay}>
	<div class="flex h-full flex-col">
		<div class="flex flex-none justify-end">
			<ActionButton imgsrc={cross} action={closeOverlay} width="20px" alt="Close" />
		</div>
		<div class="flex flex-1 flex-col overflow-y-auto">
			<ActivityContent data={selectedActivity} />
			{#if selectedActivity != null}
				<div class="mt-auto flex flex-row items-center justify-between p-2">
					<div>
						<a class="hover:text-indigo-400" href={`/activities/${selectedActivity.id}`}
							>Go to page</a
						>
					</div>
				</div>
			{/if}
		</div>
	</div>
</Overlay>
