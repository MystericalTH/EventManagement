const createAdminState = () => {
	let memberRequestList: Array<any> = $state([]);
	let activityRequestList: Array<any> = $state([]);
	let activityList: Array<any> = $state([]);
	let memberList: Array<any> = $state([]);

	return {
		get memberRequestList() {
			return memberRequestList;
		},
		set memberRequestList(data) {
			memberRequestList = data;
		},
		get activityRequestList() {
			return activityRequestList;
		},
		set activityRequestList(data) {
			activityRequestList = data;
		},
		get activityList() {
			return activityList;
		},
		set activityList(data) {
			activityList = data;
		},
		get memberList() {
			return memberList;
		},
		set memberList(data) {
			memberList = data;
		}
	};
};

export let adminState = createAdminState();
