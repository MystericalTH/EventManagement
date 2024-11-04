import type { PageLoad } from './$types';
  
export const load: PageLoad = async ({ fetch }) => {
  const response = await fetch('/api/activities');
  const activities = await response.json();
  return { activities };
};