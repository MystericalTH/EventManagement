// +page.server.ts
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals, params, request }) => {
	const sessionId = locals.sessionId;
	const activityId = Number(params.id);

	const isRegistered = sessionId ? await isUserRegistered(sessionId, activityId, request) : false;

	const activity = await getActivityData(activityId);
	const nextActivityId = await getNextActivityId(activityId);

	return {
		activity,
		isRegistered,
		nextActivityId
	};
};

// Function to check if the user is registered for the activity
async function isUserRegistered(
	sessionId: string,
	activityId: number,
	request: Request
): Promise<boolean> {
	try {
		const response = await fetch(`/api/activities/${activityId}/registration`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Cookie: request.headers.get('cookie') || ''
			},
			body: JSON.stringify({ activity_id: activityId })
		});

		if (!response.ok) {
			throw new Error('Failed to check registration status');
		}

		const result = await response.json();
		return result.is_registered;
	} catch (error) {
		console.error('Error checking registration status:', error);
		return false;
	}
}

// Function to fetch activity data
async function getActivityData(activityId: number) {
	// const response = await fetch(`/api/activities/${activityId}`);
	// if (!response.ok) {
	//   throw new Error('Failed to fetch activity data');
	// }
	// return await response.json();
	return {
		id: activityId,
		title: 'Sample Activity',
		startDate: '2023-10-01',
		format: 'Online',
		description: 'This is a sample activity description.'
	};
}

// Function to fetch the next activity ID
async function getNextActivityId(currentActivityId: number): Promise<number | null> {
	// const response = await fetch(`api/activities/${currentActivityId+1}`);
	// if (!response.ok) {
	//   return null; // No next activity
	// }
	// const result = await response.json();
	// return result.nextActivityId || null;
	if (currentActivityId >= 3) {
		return null;
	}
	return currentActivityId + 1;
}
