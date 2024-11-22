import type { PageLoad } from './engagements/$types';
import type { Activity } from '$lib/types';

export const load: PageLoad = async ({ fetch }) => {
	const response = await fetch('/api/member/activities');
	if (!response.ok) {
		console.log('err');
		throw new Error(`Failed to load activities: ${response.status} ${response.statusText}`);
	}
	let fetchActivities = await response.json();
	let activities: Activity[] = fetchActivities == null ? [] : [...fetchActivities];
	return { activities };
};
