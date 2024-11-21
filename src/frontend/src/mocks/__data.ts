import type { ActivityData } from '$lib/types/activity';
export const memberData = [
	{
		id: 1,
		fname: 'John',
		lname: 'Doe',
		phone: '123-456-7890',
		email: 'john.doe@example.com',
		interest: 'Reading' // Changed to singular reason: 'Looking to join a community of like-minded individuals.'
	},
	{
		id: 2,
		fname: 'Jane',
		lname: 'Smith',
		phone: '234-567-8901',
		email: 'jane.smith@example.com',
		interest: 'Cooking', // Changed to singular
		reason: 'Want to share my cooking experiences.'
	},
	{
		id: 3,
		fname: 'Alice',
		lname: 'Johnson',
		phone: '345-678-9012',
		email: 'alice.johnson@example.com',
		interest: 'Photography', // Changed to singular
		reason: 'Interested in learning more about photography.'
	},
	{
		id: 4,
		fname: 'Bob',
		lname: 'Brown',
		phone: '456-789-0123',
		email: 'bob.brown@example.com',
		interest: 'Cycling', // Changed to singular
		reason: 'Looking for a running group.'
	},
	{
		id: 5,
		fname: 'Charlie',
		lname: 'Davis',
		phone: '567-890-1234',
		email: 'charlie.davis@example.com',
		interest: 'Gaming', // Changed to singular
		reason: 'Want to meet other gamers.'
	},
	{
		id: 6,
		fname: 'Diana',
		lname: 'Miller',
		phone: '678-901-2345',
		email: 'diana.miller@example.com',
		interest: 'Yoga', // Changed to singular
		reason: 'Interested in mindfulness practices.'
	},
	{
		id: 7,
		fname: 'Eve',
		lname: 'Wilson',
		phone: '789-012-3456',
		email: 'eve.wilson@example.com',
		interest: 'Painting', // Changed to singular
		reason: 'Looking for art classes.'
	},
	{
		id: 8,
		fname: 'Frank',
		lname: 'Moore',
		phone: '890-123-4567',
		email: 'frank.moore@example.com',
		interest: 'Fishing', // Changed to singular
		reason: 'Want to join outdoor activities.'
	},
	{
		id: 9,
		fname: 'Grace',
		lname: 'Taylor',
		phone: '901-234-5678',
		email: 'grace.taylor@example.com',
		interest: 'Dancing', // Changed to singular
		reason: 'Looking for a dance group.'
	},
	{
		id: 10,
		fname: 'Henry',
		lname: 'Anderson',
		phone: '012-345-6789',
		email: 'henry.anderson@example.com',
		interest: 'Writing', // Changed to singular
		reason: 'Want to share my writing.'
	},
	{
		id: 11,
		fname: 'Ivy',
		lname: 'Thomas',
		phone: '123-456-7891',
		email: 'ivy.thomas@example.com',
		interest: 'Writing', // Changed to singular
		reason: 'Want to share my writing.'
	},
	{
		id: 12,
		fname: 'Jack',
		lname: 'White',
		phone: '234-567-8902',
		email: 'jack.white@example.com',
		interest: 'Music', // Changed to singular
		reason: 'Looking to join a music band.'
	},
	{
		id: 13,
		fname: 'Karen',
		lname: 'Green',
		phone: '345-678-9013',
		email: 'karen.green@example.com',
		interest: 'Cooking', // Changed to singular
		reason: 'Want to share recipes and book recommendations.'
	},
	{
		id: 14,
		fname: 'Leo',
		lname: 'King',
		phone: '456-789-0124',
		email: 'leo.king@example.com',
		interest: 'Running', // Changed to singular
		reason: 'Looking for a fitness group.'
	},
	{
		id: 15,
		fname: 'Mia',
		lname: 'Scott',
		phone: '567-890-1235',
		email: 'mia.scott@example.com',
		interest: 'Photography', // Changed to singular
		reason: 'Interested in travel photography.'
	}
];

export const activityData: ActivityData[] = [
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
	},
	{
		id: 4,
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
		id: 5,
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
		id: 6,
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
];
