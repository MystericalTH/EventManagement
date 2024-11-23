import type { Activity, Feedback MemberRegistration, Member } from '$lib/types';

export const memberData: Member[] = [
	{
		id: 1,
		fname: 'John',
		lname: 'Doe',
		phone: '123-456-7890',
		email: 'john.doe@example.com',
		interest: 'Reading', // Changed to singular
		reason: 'Looking to join a community of like-minded individuals.'
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

export const activityData: Activity[] = [
	{
		id: 1,
		title: 'Morning Yoga',
		proposer: 'Alice',
		startDate: '2023-11-01',
		startTime: '06:00',
		endDate: '2023-11-01',
		endTime: '07:00',
		maxParticipant: 20,
		format: 'Project',
		description: 'A relaxing morning yoga session.',
		proposeDateTime: '2023-10-20T08:00:00Z',
		acceptAdmin: 'Admin1',
		acceptDateTime: '2023-10-21T09:00:00Z',
		roles: ['guide', 'participant'],
		applicationStatus: 'Pending'
	},
	{
		id: 2,
		title: 'Tech Talk',
		proposer: 'Bob',
		startDate: '2023-11-02',
		startTime: '14:00',
		endDate: '2023-11-02',
		endTime: '16:00',
		maxParticipant: 50,
		format: 'Online',
		description: 'A talk on the latest in tech.',
		proposeDateTime: '2023-10-22T10:00:00Z',
		acceptAdmin: 'Admin2',
		acceptDateTime: '2023-10-23T11:00:00Z',
		roles: ['speaker', 'listener'],
		applicationStatus: 'Approved'
	},
	{
		id: 3,
		title: 'Cooking Workshop',
		proposer: 'Charlie',
		startDate: '2023-11-04',
		startTime: null,
		endDate: '2023-11-05',
		endTime: null,
		maxParticipant: 15,
		format: 'Project',
		description: 'Learn to cook delicious meals.',
		proposeDateTime: '2023-10-23T12:00:00Z',
		acceptAdmin: 'Admin3',
		acceptDateTime: '2023-10-24T13:00:00Z',
		roles: ['guide', 'participant'],
		applicationStatus: 'Pending'
	},
	{
		id: 4,
		title: 'Art Exhibition',
		proposer: 'Dave',
		startDate: '2023-11-04',
		startTime: null,
		endDate: '2023-11-04',
		endTime: null,
		maxParticipant: 100,
		format: 'Project',
		description: 'An exhibition of modern art.',
		proposeDateTime: '2023-10-24T14:00:00Z',
		acceptAdmin: 'Admin4',
		acceptDateTime: '2023-10-25T15:00:00Z',
		roles: ['guide', 'participant'],
		applicationStatus: 'Approved'
	},
	{
		id: 5,
		title: 'Music Concert',
		proposer: 'Eve',
		startDate: '2023-11-05',
		startTime: '18:00',
		endDate: '2023-11-06',
		endTime: '21:00',
		maxParticipant: 200,
		format: 'Project',
		description: 'A live music concert.',
		proposeDateTime: '2023-10-25T16:00:00Z',
		acceptAdmin: 'Admin5',
		acceptDateTime: '2023-10-26T17:00:00Z',
		roles: ['guide', 'participant'],
		applicationStatus: 'Approved'
	}
];
export const mockFeedbackData: Feedback[] = [
	{
		feedbackid: 1,
		activityid: 101,
		memberid: 1001,
		fname: 'John',
		lname: 'Doe',
		feedbackmessage: 'Great event!',
		feedbackdatetime: '2023-10-01T10:00:00Z'
	},
	{
		feedbackid: 2,
		activityid: 101,
		memberid: 1002,
		fname: 'Jane',
		lname: 'Smith',
		feedbackmessage: 'Had a wonderful time.',
		feedbackdatetime: '2023-10-01T11:00:00Z'
	},
	{
		feedbackid: 3,
		activityid: 101,
		memberid: 1003,
		fname: 'Alice',
		lname: 'Johnson',
		feedbackmessage: 'Very well organized.',
		feedbackdatetime: '2023-10-01T12:00:00Z'
	}
];
export const memberRegistrationData: MemberRegistration[] = [
	{
		memberid: 1,
		fname: 'John',
		lname: 'Doe',
		phone: '123-456-7890',
		email: 'john.doe@example.com',
		role: 'Organizer',
		expectation: 'To manage events efficiently'
	},
	{
		memberid: 2,
		fname: 'Jane',
		lname: 'Smith',
		phone: '987-654-3210',
		email: 'jane.smith@example.com',
		role: 'Volunteer',
		expectation: 'To help with event logistics'
	},
	{
		memberid: 3,
		fname: 'Alice',
		lname: 'Johnson',
		phone: '555-123-4567',
		email: 'alice.johnson@example.com',
		role: 'Speaker',
		expectation: 'To deliver insightful talks'
	},
	{
		memberid: 4,
		fname: 'Bob',
		lname: 'Brown',
		phone: '444-555-6666',
		email: 'bob.brown@example.com',
		role: 'Attendee',
		expectation: 'To network and learn'
	},
	{
		memberid: 5,
		fname: 'Charlie',
		lname: 'Davis',
		phone: '333-444-5555',
		email: 'charlie.davis@example.com',
		role: 'Sponsor',
		expectation: 'To promote products and services'
	}
];
