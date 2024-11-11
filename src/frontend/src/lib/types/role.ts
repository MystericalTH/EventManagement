export enum Role {
	Admin = 'admin',
	Member = 'member',
	Default = ''
}

export let getRole = function (roleString: string): Role {
	if (roleString === 'admin') return Role.Admin;
	else if (roleString === 'member') return Role.Member;
	else return Role.Default;
};
