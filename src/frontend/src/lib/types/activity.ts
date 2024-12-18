export interface Activity {
	id: number;
	title: string;
	proposer: string;
	startDate: string;
	startTime: string;
	endDate: string;
	endTime: string;
	maxParticipant: number;
	format: string;
	description: string;
	proposeDateTime: string;
	advisor: string;
	acceptAdmin?: string;
	acceptDateTime?: string;
	applicationStatus: string;
	activityRoles: string[];
}
