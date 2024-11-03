<!-- +page.svelte -->
<!-- <script context="module" lang="ts">
    import type { PageLoad } from './$types';
  
    export const load: PageLoad = async ({ fetch }) => {
      const response = await fetch('/api/activities');
      const activities = await response.json();
      return { activities };
    };
</script> -->
  
<script lang="ts">
    import Card from '../../components/Card.svelte';
    // export let data: { activities: Array<{ title: string; startDate: string; endDate: string; format: string }> };

    // Mock data
    const data = {
        activities: [
        { title: 'Activity 1', startDate: '2025-11-01', endDate: '2025-11-05', format: 'Project', href: '/activity/1' },
        { title: 'Activity 2', startDate: '2024-11-01', endDate: '2024-11-30', format: 'Project', href: '/activity/2' },
        { title: 'Activity 3', startDate: '2023-09-01', endDate: '2023-09-05', format: 'Workshop', href: '/activity/3' }
        ]
    };

    const currentDate = new Date();

    const upcomingActivities = data.activities.filter(activity => {
      const startDate = new Date(activity.startDate);
      return startDate > currentDate;
    });

    const ongoingActivities = data.activities.filter(activity => {
      const startDate = new Date(activity.startDate);
      const endDate = new Date(activity.endDate);
      return startDate <= currentDate && currentDate <= endDate;
    });

    const completedActivities = data.activities.filter(activity => {
      const endDate = new Date(activity.endDate);
      return endDate < currentDate;
    });
</script>
  
<h1>Up Coming</h1>
{#each upcomingActivities as activity}
  <Card title={activity.title} startDate={activity.startDate} endDate={activity.endDate} format={activity.format} href={activity.href} />
{/each}
  
<h1>On Going</h1>
{#each ongoingActivities as activity}
  <Card title={activity.title} startDate={activity.startDate} endDate={activity.endDate} format={activity.format} href={activity.href} />
{/each}

<h1>Completed</h1>
{#each completedActivities as activity}
  <Card title={activity.title} startDate={activity.startDate} endDate={activity.endDate} format={activity.format} href={activity.href} />
{/each}