import type { LayoutServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

export const load: LayoutServerLoad = async ({ fetch, url }) => {
	const response = await fetch('/api/login/callback');
	if (response.ok) {
		const data = await response.json();

		const userRole = data.role;

		const expectedPathPrefix = `/${userRole}`;

		const publicPaths = ['/activities', '/home', '/signup'];

		const isPublicPath = publicPaths.some((path) => url.pathname.startsWith(path));

		if (!isPublicPath && !url.pathname.startsWith(expectedPathPrefix)) {
			throw redirect(302, '/home');
		}

		return {
			role: data.role,
			user: data.user
		};
	} else {
		return {
			role: null,
			user: null
		};
	}
};
