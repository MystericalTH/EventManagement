export type Member = {
	id: number;
	fname: string;
	lname: string;
	phone: string;
	email: string;
	interest: string;
	githuburl?: string;
	reason: string;
};

export type MemberRegistration = {
	memberid: number;
	fname: string;
	lname: string;
	phone: string;
	email: string;
	role: string;
	expectation: string;
};
