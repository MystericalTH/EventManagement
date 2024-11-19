import { adminState } from '$lib/states/adminStates.svelte';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	const res = await fetch(`/api/members`);
	let memberList = await res.json();
	if (memberList == null) memberList = [];
	adminState.memberList = memberList;
	return { memberList };
};
