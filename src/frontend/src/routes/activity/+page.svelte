<script lang="ts">
    export let data: { activities: Array<{ title: string; startDate: string; endDate: string; format: string; href: string }> };

    import Card from '../../components/Card.svelte';

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