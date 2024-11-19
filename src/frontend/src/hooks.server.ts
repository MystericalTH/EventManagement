import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	// Retrieve session ID from cookies
	const sessionId = event.cookies.get('session_id');

	// Store session ID in locals
	event.locals.sessionId = sessionId;

	return await resolve(event);
};
