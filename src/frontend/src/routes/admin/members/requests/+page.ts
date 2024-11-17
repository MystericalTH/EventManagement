import type { PageLoad } from './$types';
import { adminState } from '$lib/states/adminStates.svelte';
export const load: PageLoad = async ({ fetch }) => {
	const res = await fetch(`/api/members/requests`);
	adminState.memberRequestList = await res.json();

	return { adminState: adminState };
};
