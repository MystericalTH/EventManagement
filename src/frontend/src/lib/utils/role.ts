export const Role = {
	ADMIN: 'admin',
	MEMBER: 'member',
	DEVELOPER: 'developer',
	DEFAULT: 'default'
};

export let getRole = function (roleString: string): string | null {
	if (roleString === 'admin') return Role.ADMIN;
	else if (roleString === 'member') return Role.MEMBER;
	else if (roleString === 'developer') return Role.DEVELOPER;
	else if (roleString === 'default') return Role.DEFAULT;
	else return null;
};
