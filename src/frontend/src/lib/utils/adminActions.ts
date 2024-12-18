import type { Member, Pagination } from '$lib/types';
const removeItem = (id: number, pagination: Pagination<any>) => {
	let idx: number = pagination.data.findIndex((e: any) => {
		if (e.memberid != null) e.id = e.memberid;
		return e.id === id;
	});

	pagination.data.splice(idx, 1);
	if (pagination.maxPage < pagination.pageBuffer) {
		pagination.pageBuffer = pagination.maxPage;
		pagination.setPage();
	}
};

export const rejectMemberRequest = (id: number, pagination: Pagination<any>) => {
	fetch(`/api/members/${id}`, { method: 'DELETE' }).then((r) => {
		if (r.status == 204) {
			removeItem(id, pagination);
		} else {
			console.log('Not deleted');
			// TODO:  add handling
		}
	});
};

export const approveMemberRequest = (id: number, pagination: Pagination<any>) => {
	fetch(`/api/members/${id}/approve`, { method: 'PUT' }).then((r) => {
		if (r.status == 204) {
			removeItem(id, pagination);
		} else {
			console.log('Not approved');
			// TODO:  add handling
		}
	});
};

export const removeMember = (id: number, pagination: Pagination<any>) => {
	fetch(`/api/members/${id}`, { method: 'DELETE' }).then((r) => {
		if (r.status == 204) {
			removeItem(id, pagination);
		} else {
			console.log('Not deleted');
			// TODO:  add handling
		}
	});
};

export const approveActivityRequest = (id: number, pagination: Pagination<any>) => {
	fetch(`/api/activities/${id}/approve`, { method: 'PUT' }).then((r) => {
		if (r.status == 204) {
			removeItem(id, pagination);
		} else {
			console.log('Not approved');
			// TODO:  add handling
		}
	});
};

export const removeActivity = (id: number, pagination: Pagination<any>) => {
	fetch(`/api/activities/${id}`, { method: 'DELETE' }).then((r) => {
		if (r.status == 204) {
			removeItem(id, pagination);
		} else {
			console.log('Not deleted');
			// TODO:  add handling
		}
	});
};

export const rejectActivityRequest = (id: number, pagination: Pagination<any>) => {
	fetch(`/api/activities/${id}`, { method: 'DELETE' }).then((r) => {
		if (r.status == 204) {
			removeItem(id, pagination);
		} else {
			console.log('Not deleted');
			// TODO:  add handling
		}
	});
};
