import type { MemberRequest } from '$lib/types/memberRequest';
import { adminState } from '$lib/states/adminStates.svelte';
const removeCachedRequest = (id: number, pagination: any) => {
	adminState.memberRequestList.splice(
		adminState.memberRequestList.findIndex((e: MemberRequest) => e.id === id),
		1
	);
	if (pagination.maxPage < pagination.pageBuffer) {
		pagination.setPage(pagination.maxPage);
	}
};

export const rejectRequest = (id: number, pagination) => {
	fetch(`/api/members/requests/${id}`, { method: 'DELETE' }).then((r) => {
		if (r.status == 204) {
			removeCachedRequest(id, pagination);
		} else {
			console.log('Not deleted');
			// TODO:  add handling
		}
	});
};
export const approveRequest = (id: number, pagination) => {
	fetch(`/api/members/requests/${id}/approve`, { method: 'POST' }).then((r) => {
		if (r.status == 204) {
			removeCachedRequest(id, pagination);
		} else {
			console.log('Not approved');
			// TODO:  add handling
		}
	});
};
