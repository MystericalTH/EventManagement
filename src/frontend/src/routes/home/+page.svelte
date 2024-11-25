<script lang="ts">
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import type { Activity } from '$lib/types';
    import Overlay from '$lib/components/Overlay.svelte';
    import ActivityContent from '$lib/components/ActivityContent.svelte';
    import ActionButton from '$lib/components/ActionButton.svelte';
    import Card from '$lib/components/activity/Card.svelte';
    import { cross } from '$lib/assets/action-button-icons';

    let activities: Activity[] = [];
    let upcomingActivities: Activity[] = [];
    let showOverlay = false;
    let selectedActivity: Activity | null = null;

    onMount(async () => {
        try {
            const res = await fetch('/api/activities');
            if (res.ok) {
                activities = await res.json();
                const currentDate = new Date();
                upcomingActivities = activities.filter((activity: Activity) => {
                    const startDate = new Date(activity.startDate);
                    return startDate > currentDate;
                }).slice(0, 3); // Show only the first 3 upcoming activities
            } else {
                console.error('Failed to fetch activities');
            }
        } catch (error) {
            console.error('Error fetching activities:', error);
        }
    });

    function viewActivity(activity: Activity) {
        selectedActivity = activity;
        showOverlay = true;
    }

    function cardOnclick(activity: Activity) {
        viewActivity(activity);
    }
    
    const closeOverlay = () => {
        selectedActivity = null;
        showOverlay = false;
    };
</script>

<div class="home-introduction text-center my-8 mx-auto max-w-3xl">
    <h1 class="text-4xl font-bold mb-4">Welcome to Our Club</h1>
    <p class="text-lg mb-6 leading-relaxed text-justify px-4">
        We are a community passionate about exploring new ideas and sharing knowledge.
        Our club offers a range of activities, workshops, and projects that members can participate in.
        Whether you're interested in learning new skills, collaborating on projects, or meeting like-minded individuals,
        there's something here for everyone.
    </p>
    <button on:click={() => goto('/activities')} class="btn btn-primary bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-700">
        Explore Activities
    </button>
</div>

<div class="featured-activities my-8 mx-auto w-10/12">
    <h2 class="text-2xl font-semibold mb-4 text-left">Some Upcoming Activities</h2>
    {#if upcomingActivities.length > 0}
        <div class="flex flex-row space-x-4 overflow-x-auto">
            {#each upcomingActivities as activity}
                <!-- <li class="activity-item border border-gray-300 p-4 mb-4 rounded">
                    <h3 class="text-xl font-semibold">{activity.title}</h3>
                    <p class="text-gray-700 mb-2">{activity.description}</p>
                    <button on:click={() => viewActivity(activity)} class="btn btn-secondary bg-gray-500 text-white py-2 px-4 rounded hover:bg-gray-700">
                        View Details
                    </button>
                </li> -->
                <Card data={activity} onclick={cardOnclick} />
            {/each}
        </div>
    {:else}
        <p class="text-gray-700">No upcoming activities at the moment. Please check back later.</p>
    {/if}
</div>

<Overlay {showOverlay}>
    <div class="flex justify-end">
        <ActionButton imgsrc={cross} action={closeOverlay} width="20px" alt="Close" />
    </div>
    <div class="overflow-y-scroll">
        <ActivityContent data={selectedActivity} />
    </div>
</Overlay>