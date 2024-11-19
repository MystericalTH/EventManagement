import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async ({ fetch }) => {
	const response = await fetch('/api/login/callback');
	const res = await response.json();
	return {
		role: res.role
	};
};
