export interface Activity {
	id: number;
	title: string;
	proposer: string;
	startDate: string;
	startTime: string;
	endDate: string;
	endTime: string;
	maxNumbert: number;
	format: string;
	description: string;
	proposeDateTime: string;
	advisor: string;
	acceptAdmin?: string;
	acceptDateTime?: string;
	applicationStatus?: string;
	activityRole: string[];
}
