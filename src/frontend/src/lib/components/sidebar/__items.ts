import type { SidebarItem } from "$lib/types/sidebar";

export const defaultItems: SidebarItem[] = [
  {
    text: 'Signup',
    href: '/api/signup'
  },
  {
    text: 'Login',
    subitems: [
      { text: 'Member', href: '/api/login?role=member' },
      { text: 'Admin', href: '/api/login?role=admin' }
    ]
  }
]

export const adminItems: SidebarItem[] = [
  {
    text: 'Activity',
    subitems: [
      { text: 'Manage Activities', href: '/admin/activities' },
      { text: 'Activity Requests', href: '/admin/activities/requests' },
      { text: 'See Feedback', href: '/admin/activities/feedback' }
    ]
  },
  {
    text: 'Members',
    subitems: [
      { text: 'Manage Members', href: '/admin/members' },
      { text: 'Member Requests', href: '/admin/members/requests' }
    ]
  }
];

export const memberItems: SidebarItem[] = [
  {
    text: 'Activity',
    subitems: [
      { text: 'Your Activities', href: '/member/activities' },
      { text: 'Your Proposals', href: '/member/activities/proposals' }
    ]
  },
  {
    text: 'Propose New Activity',
    href: '/member/activities/propose'
  }
];

