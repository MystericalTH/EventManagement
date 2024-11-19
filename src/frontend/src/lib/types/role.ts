export const Role = {
	ADMIN: 'admin',
	MEMBER: 'member',
	DEFAULT: 'default'
};

export let getRole = function (roleString: string): string {
	if (roleString === 'admin') return Role.ADMIN;
	else if (roleString === 'member') return Role.MEMBER;
	else return Role.DEFAULT;
};
