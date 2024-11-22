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
	acceptAdmin?: string;
	acceptDateTime?: string;
	applicationStatus?: string;
	roles: string[];
}
