import type { ActivityData } from '$lib/types/activity';

export const activityData = [
	{
		id: 1,
		title: 'Activity 1',
		proposer: 'John Doe',
		startDate: '2021-10-01',
		endDate: '2021-10-02',
		maxNumber: 10,
		format: 'Project',
		description: 'This is activity 1',
		proposeDateTime: '2021-09-01',
		acceptAdmin: 'Admin',
		acceptDateTime: '2021-09-02',
		applicationStatus: 'Approved'
	},
	{
		id: 2,
		title: 'Activity 2',
		proposer: 'John Doe',
		startDate: '2021-10-01',
		endDate: '2021-10-02',
		maxNumber: 10,
		format: 'Project',
		description: 'This is activity 2',
		proposeDateTime: '2021-09-01',
		acceptAdmin: 'Admin',
		acceptDateTime: '2021-09-02',
		applicationStatus: 'Approved'
	},
	{
		id: 3,
		title: 'Activity 3',
		proposer: 'John Doe',
		startDate: '2021-10-01',
		endDate: '2021-10-02',
		maxNumber: 10,
		format: 'Workshop',
		description: 'This is activity 3',
		proposeDateTime: '2021-09-01',
		acceptAdmin: 'Admin',
		acceptDateTime: '2021-09-02',
		applicationStatus: 'Approved'
	}
] as ActivityData[];
