export interface ActivityData {
	id: number;
	title: string;
	proposer: string;
	startDate: string;
	endDate: string;
	maxNumber: number;
	format: string;
	description: string;
	proposeDateTime: string;
	acceptAdmin?: string;
	acceptDateTime?: string;
	applicationStatus?: string;
}
