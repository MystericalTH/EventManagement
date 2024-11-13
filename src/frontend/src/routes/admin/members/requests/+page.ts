import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, params }) => {
	const res = await fetch(`/api/members/requests`);
	const item = await res.json();
	return { requestList: item };
};
