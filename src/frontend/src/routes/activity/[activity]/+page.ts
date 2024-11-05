import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, params }) => {
  const response = await fetch(`/api/activities/${params}`);
//   const activity = await response.json();
  const activity = { id: '1', title: 'Activity 1', startDate: '2025-11-01', format: 'Project', description: 'This is description' };
  return { activity };
};