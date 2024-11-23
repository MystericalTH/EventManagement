import type { Activity } from '$lib/types/activity';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	const response = await fetch('/api/activities');
	if (!response.ok) {
		console.log('err');
		throw new Error(`Failed to load activities: ${response.status} ${response.statusText}`);
	}
	const activities: Activity[] = await response.json();
	return { activities };
};
