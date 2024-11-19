import { adminState } from '$lib/states/adminStates.svelte';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	const res = await fetch(`/api/members`);
	const memberList = await res.json();
	adminState.memberList = memberList;
	return { memberList };
};
