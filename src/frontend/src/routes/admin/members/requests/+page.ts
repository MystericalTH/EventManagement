import type { PageLoad } from './$types';
import { adminState } from '$lib/states/adminStates.svelte';
export const load: PageLoad = async ({ fetch }) => {
	const res = await fetch(`/api/members/requests`);
	let memberRequestList = await res.json();
	if (memberRequestList == null) memberRequestList = [];
	adminState.memberRequestList = memberRequestList;

	return { memberRequestList };
};
