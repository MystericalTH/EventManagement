import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch }) => {
	const id = Number(params.id);

	// Function to fetch the next activity ID

	try {
		// Fetch the activity data using the ID
		const responseActivity = await fetch(`/api/activities/${id}`);
		if (!responseActivity.ok) throw new Error('Failed to fetch activity data');
		const activity = await responseActivity.json();

		// Fetch the registration status of the current user
		const registrationResponse = await fetch(`/api/activities/${id}/registration/status`);
		if (!registrationResponse.ok) throw new Error('Failed to fetch registration status');
		const isRegistered = await registrationResponse.json();

		// Fetch the feedback status of the current user
		const feedbackResponse = await fetch(`/api/activities/${id}/feedback/status`);
		if (!feedbackResponse.ok) throw new Error('Failed to fetch feedback status');
		const hasSubmittedFeedback = await feedbackResponse.json();
		// const response = await fetch(`/api/activities/${id + 1}`);
		// if (!response.ok) {
		// 	var nextActivityId = null; // No next activity
		// } else {
		// 	const result = await response.json();
		// 	var nextActivityId = id + 1;
		// }
		// Check if the event has already passed
		const isEventPast = new Date(activity.date) < new Date();

		// Fetch the next activity ID
		let data = {
			activity,
			isRegistered,
			isEventPast,
			hasSubmittedFeedback,
			// nextActivityId
		};
		console.log("from page.ts",data);
		return data;
	} catch (error) {
		console.error(error);
		return {
			activity: null,
			isRegistered: false,
			isEventPast: false,
			hasSubmittedFeedback: false,
			// nextActivityId: null
		};
	}
};
