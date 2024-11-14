// +page.server.ts
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals, params, request }) => {
	const sessionId = locals.sessionId;
	const activityId = Number(params.id);

	const isRegistered = sessionId ? await isUserRegistered(sessionId, activityId, request) : false;

	const activity = await getActivityData(activityId);
	const nextActivityId = await getNextActivityId(activityId);

  // Check if the event is past its end date
  const isEventPast = new Date() > new Date(activity.endDate);

  // Check if the user has submitted feedback
  const hasSubmittedFeedback = sessionId
    ? await userHasSubmittedFeedback(sessionId, activityId, request)
    : false;

  return {
    activity,
    isRegistered,
    isEventPast,
    hasSubmittedFeedback,
    nextActivityId
  };

// Function to check if the user has submitted feedback
async function userHasSubmittedFeedback(sessionId: string, activityId: number, request: Request): Promise<boolean> {
  try {
    const response = await fetch(`/api/activities/${activityId}/feedback/status`, {
      headers: {
        'Cookie': request.headers.get('cookie') || ''
      }
    });

    if (!response.ok) {
      throw new Error('Failed to check feedback status');
    }

    const result = await response.json();
    return result.hasSubmittedFeedback;
  } catch (error) {
    console.error('Error checking feedback status:', error);
    return false;
  }
}

// Function to check if the user is registered for the activity
async function isUserRegistered(sessionId: string, activityId: number, request: Request): Promise<boolean> {
  try {
    const response = await fetch(`/api/activities/${activityId}/registration/status`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Cookie': request.headers.get('cookie') || ''
      },
      body: JSON.stringify({ activity_id: activityId })
    });
// =======
// async function isUserRegistered(
// 	sessionId: string,
// 	activityId: number,
// 	request: Request
// ): Promise<boolean> {
// 	try {
// 		const response = await fetch(`/api/activities/${activityId}/registration`, {
// 			method: 'POST',
// 			headers: {
// 				'Content-Type': 'application/json',
// 				Cookie: request.headers.get('cookie') || ''
// 			},
// 			body: JSON.stringify({ activity_id: activityId })
// 		});

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
    endDate: '2023-10-31',
    format: 'Online',
    description: 'This is a sample activity description.'
  };

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
