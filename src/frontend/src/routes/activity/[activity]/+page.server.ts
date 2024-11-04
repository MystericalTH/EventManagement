import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals }) => {
  const sessionId = locals.sessionId;

  // Use the session ID to check registration status
  const isRegistered = sessionId ? await isUserRegistered(sessionId) : false;

  return { isRegistered };
};

// Function to check if the user is registered
async function isUserRegistered(sessionId: string): Promise<boolean> {
  // Replace with your logic to verify registration
  // For example, query your database with the sessionId
  return true; // or false based on registration status
}