import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch }) => {
	const id = params.id;

	try {
		// Fetch the activity data using the ID
		const response = await fetch(`/api/activities/${id}`);
		if (!response.ok) throw new Error('Failed to fetch activity data');
		const activity = await response.json();

		return {
			activity
		};
	} catch (error) {
		console.error(error);
		return {
			activity: null
		};
	}
};
