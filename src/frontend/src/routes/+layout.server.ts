import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ fetch }) => {
	const response = await fetch('/api/login/callback');
	const res = await response.json();
	return {
		role: res.role
	};
};
