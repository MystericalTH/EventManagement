import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	try {
		const response = await fetch('/api/activities/requests');
		let activities = await response.json();
		if (activities == null) activities = [];

		return { activities };
	} catch {
		return { activities: [] };
	}
};
