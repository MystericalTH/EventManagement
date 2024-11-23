import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	const response = await fetch('/api/activities');
	let activities = await response.json();
	if (activities == null) activities = [];

	return { activities };
};
