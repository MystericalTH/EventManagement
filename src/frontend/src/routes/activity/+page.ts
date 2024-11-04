import type { PageLoad } from './$types';
  
export const load: PageLoad = async ({ fetch }) => {
  const response = await fetch('/api/activities');
//   const activities = await response.json();
  const activities = [
    { title: 'Activity 1', startDate: '2025-11-01', endDate: '2025-11-05', format: 'Project', href: '/activity/1' },
    { title: 'Activity 2', startDate: '2024-11-01', endDate: '2024-11-30', format: 'Project', href: '/activity/2' },
    { title: 'Activity 3', startDate: '2023-09-01', endDate: '2023-09-05', format: 'Workshop', href: '/activity/3' }
    ];
  return { activities };
};