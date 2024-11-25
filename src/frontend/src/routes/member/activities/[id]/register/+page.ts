import type { PageLoad } from './$types';
import { error } from '@sveltejs/kit';

export const load: PageLoad = async ({ params, fetch }) => {
	const id = params.id;

	// Fetch the activity data using the ID
	const response = await fetch(`/api/activities/${id}`);
	if (!response.ok) throw new Error('Failed to fetch activity data');
	const activity = await response.json();

	if (new Date(activity.startDate) < new Date()) {
		error(404, {
			message: 'Registration closed'
		});
	}

	return {
		activity
	};
};
