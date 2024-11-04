import type { PageLoad } from './$types';
  
export const load: PageLoad = async ({ fetch }) => {
  const response = await fetch('/api/activities');
//   const activities = await response.json();
  const activities = [
    { id: 1, title: 'Activity 1', startDate: '2025-11-01', endDate: '2025-11-05', format: 'Project' },
    { id: 2, title: 'Activity 2', startDate: '2024-11-01', endDate: '2024-11-30', format: 'Project' },
    { id: 3, title: 'Activity 3', startDate: '2023-09-01', endDate: '2023-09-05', format: 'Workshop' }
    ];
  return { activities };
};