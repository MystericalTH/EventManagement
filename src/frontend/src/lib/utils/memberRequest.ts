import type { MemberRequest } from '$lib/types/memberRequest';
import { adminState } from '$lib/states/states.svelte';
const removeCachedRequest = (id: number) => {
	adminState.memberRequestList.splice(
		adminState.memberRequestList.findIndex((e: MemberRequest) => e.id === id),
		1
	);
};

export const rejectRequest = (id: number) => {
	fetch(`/api/members/requests/${id}`, { method: 'DELETE' }).then((r) => {
		if (r.status == 204) {
			removeCachedRequest(id);
		} else {
			console.log('Not deleted');
			// TODO:  add handling
		}
	});
};
export const approveRequest = (id: number) => {
	fetch(`/api/members/requests/${id}/approve`, { method: 'POST' }).then((r) => {
		if (r.status == 204) {
			removeCachedRequest(id);
		} else {
			console.log('Not approved');
			// TODO:  add handling
		}
	});
};
