import type { SidebarItem } from '$lib/types';

export const defaultItems = (currentUrl: string): SidebarItem[] => {
	return [
		{
			text: 'Signup',
			href: '/signup'
		},
		{
			text: 'Login',
			subitems: [
				{ text: 'Member', href: `/api/login?role=member\&redirect_uri=${currentUrl}` },
				{ text: 'Admin', href: `/api/login?role=admin\&redirect_uri=${currentUrl}` },
				{ text: 'Developer', href: `/api/login?role=developer\&redirect_uri=${currentUrl}` }
			]
		}
	];
};

export const adminItems: SidebarItem[] = [
	{
		text: 'Activities',
		subitems: [
			{ text: 'Manage Activities', href: '/admin/activities' },
			{ text: 'Activity Requests', href: '/admin/activities/requests' },
			{ text: 'See Feedback', href: '/admin/feedback' }
		]
	},
	{
		text: 'Members',
		subitems: [
			{ text: 'Manage Members', href: '/admin/members' },
			{ text: 'Member Requests', href: '/admin/members/requests' }
		]
	},
	{ text: 'Chat with Developer', href: '/admin/chat/dev' }
];

export const memberItems: SidebarItem[] = [
	{ text: 'Home', href: '/home' },
	{ text: 'Activities', href: '/activities' },
	{
		text: 'Personal List',
		subitems: [
			{ text: 'Your Engagements', href: '/member/activities' },
			{ text: 'Your Proposals', href: '/member/activities/proposals' },
			{
				text: 'Propose Activity',
				href: '/member/activities/propose'
			}
		]
	}
];

export const developerItems: SidebarItem[] = [
	{ text: 'Dev Documentation', href: '/developer/dev-docs' },
	{ text: 'Logging Dashboard', href: '/developer/logging' },
	{ text: 'Chat with Admin', href: '/developer/chat/admin' }
];
