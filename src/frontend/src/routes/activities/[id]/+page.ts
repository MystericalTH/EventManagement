import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch }) => {
	const id = Number(params.id);

	// Function to fetch the next activity ID
	var isAuthorized = true;
	// Fetch the activity data using the ID
	const responseActivity = await fetch(`/api/activities/${id}`);
	if (!responseActivity.ok) throw new Error('Failed to fetch activity data');
	const activity = await responseActivity.json();

	// Fetch the registration status of the current user
	const registrationResponse = await fetch(`/api/activities/${id}/registration/status`);
	if (registrationResponse.status == 401) isAuthorized = false;
	if (isAuthorized && !registrationResponse.ok)
		throw new Error('Failed to fetch registration status');
	const registrationData = await registrationResponse.json();
	const isRegistered = registrationData.isRegistered;

	// Fetch the feedback status of the current user
	const feedbackResponse = await fetch(`/api/activities/${id}/feedback/status`);
	if (isAuthorized && !feedbackResponse.ok) throw new Error('Failed to fetch feedback status');
	const feedbackData = await feedbackResponse.json();
	const hasSubmittedFeedback = feedbackData.hasSubmittedFeedback;

	const response = await fetch(`/api/activities/${id + 1}`);
	let nextActivityId: number | null;
	if (!response.ok) {
		nextActivityId = null; // No next activity
	} else {
		// const result = await response.json();
		nextActivityId = id + 1;
	}
	// Check if the event has already passed
	const isEventPast = new Date(activity.startDate) < new Date();

	// Fetch the next activity ID
	let data = {
		isAuthorized,
		activity,
		isRegistered,
		isEventPast,
		hasSubmittedFeedback,
		nextActivityId
	};
	return data;
};
