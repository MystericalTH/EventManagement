import type { MemberRequest } from '$lib/types/memberRequest';
const removeItem = (id: number, pagination: any) => {
	pagination.data.splice(
		pagination.data.findIndex((e: MemberRequest) => e.id === id),
		1
	);
	if (pagination.maxPage < pagination.pageBuffer) {
		pagination.setPage(pagination.maxPage);
	}
};

export const rejectRequest = (id: number, pagination) => {
	fetch(`/api/members/requests/${id}`, { method: 'DELETE' }).then((r) => {
		if (r.status == 204) {
			removeItem(id, pagination);
		} else {
			console.log('Not deleted');
			// TODO:  add handling
		}
	});
};

export const approveRequest = (id: number, pagination) => {
	fetch(`/api/members/requests/${id}/approve`, { method: 'POST' }).then((r) => {
		if (r.status == 204) {
			removeItem(id, pagination);
		} else {
			console.log('Not approved');
			// TODO:  add handling
		}
	});
};

export const removeMember = (id: number, pagination) => {
	fetch(`/api/members/requests/${id}`, { method: 'DELETE' }).then((r) => {
		if (r.status == 204) {
			removeItem(id, pagination);
		} else {
			console.log('Not deleted');
			// TODO:  add handling
		}
	});
};
