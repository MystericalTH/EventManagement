import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch }) => {
	const id = params.id;

	try {
		// Fetch the activity data using the ID
		const response = await fetch(`/api/activities/${id}`);
		if (!response.ok) throw new Error('Failed to fetch activity data');
		const activity = await response.json();

		// Fetch the activity roles related to this activity
		const rolesResponse = await fetch(`/api/activities/${id}/roles`);
		if (!rolesResponse.ok) throw new Error('Failed to fetch activity roles');
		const activityRoles = await rolesResponse.json();

		return {
			activity,
			activityRoles
		};
	} catch (error) {
		console.error(error);
		return {
			activity: null,
			activityRoles: []
		};
	}
};
