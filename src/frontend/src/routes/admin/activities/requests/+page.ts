import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	const response = await fetch('/api/activities/requests');
	const activities = await response.json();

	return { activities };
};
