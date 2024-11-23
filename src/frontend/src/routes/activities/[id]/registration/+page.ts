import type { PageLoad } from './$types';
import type { MemberRegistration } from '$lib/types';

export const load: PageLoad = async ({ fetch, params }) => {
	const id = params.id;

	const response = await fetch(`/api/activities/${id}/registration`);
	const activityResponse = await fetch(`/api/activities/${id}`);
	if (!response.ok || !activityResponse.ok) {
		console.log('err');
		throw new Error(`Failed to load activities: ${response.status} ${response.statusText}`);
	}
	let json = await response.json();
	let activity = await activityResponse.json();
	let registrations: MemberRegistration[] = json == null ? [] : json;
	return { registrations, activity };
};
