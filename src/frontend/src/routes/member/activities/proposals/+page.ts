import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	try {
		const response = await fetch('/api/member/activities/proposals');
		let activities = await response.json();
		if (activities == null) activities = [];

		return { activities };
	} catch {
		return { activities: [] };
	}
};
