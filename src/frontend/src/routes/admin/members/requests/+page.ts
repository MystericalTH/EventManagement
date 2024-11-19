import type { PageLoad } from './$types';
import { adminState } from '$lib/states/adminStates.svelte';
export const load: PageLoad = async ({ fetch }) => {
	const res = await fetch(`/api/members/requests`);
	let memberRequestList = await res.json();
	adminState.memberRequestList = memberRequestList;

	return { memberRequestList };
};
