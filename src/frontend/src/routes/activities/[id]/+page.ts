import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch }) => {
  const id = Number(params.id);
  
  // Function to fetch the next activity ID
  const getNextActivityId = async (currentActivityId: number): Promise<number | null> => {
    const response = await fetch(`/api/activities/${currentActivityId + 1}`);
    if (!response.ok) {
      return null; // No next activity
    }
    const result = await response.json();
    return result.nextActivityId || null;
  };

  try{
    // Fetch the activity data using the ID
    const response = await fetch(`/api/activities/${id}`);
    if (!response.ok) throw new Error('Failed to fetch activity data');
    const activity = await response.json();

    // Fetch the registration status of the current user
    const registrationResponse = await fetch(`/api/activities/${id}/registration/status`);
    if (!registrationResponse.ok) throw new Error('Failed to fetch registration status');
    const isRegistered = await registrationResponse.json();

    // Fetch the feedback status of the current user
    const feedbackResponse = await fetch(`/api/activities/${id}/feedback/status`);
    if (!feedbackResponse.ok) throw new Error('Failed to fetch feedback status');
    const hasSubmittedFeedback = await feedbackResponse.json();

    // Check if the event has already passed
    const isEventPast = new Date(activity.date) < new Date();

    // Fetch the next activity ID
    const nextActivityId = await getNextActivityId(id);

    return {
        activity,
        isRegistered,
        isEventPast,
        hasSubmittedFeedback,
        nextActivityId
      };
  } catch (error) {
    console.error(error);
    return {
      activity: null,
      isRegistered: false,
      isEventPast: false,
      hasSubmittedFeedback: false,
      nextActivityId: null
    };
  }
};
