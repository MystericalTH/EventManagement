export const load = async ({ fetch }) => {
	var verify;
	try {
		const res = await fetch(`/api/verify`);
		verify = await res.json();
		console.log('Verify: ' + verify);
	} catch (error) {
		console.log('Error: ' + error);
		if (import.meta.env.VITE_MSW_ENABLED === 'true') {
			console.log(true);
			verify = { role: 'unknown' };
		} else {
			console.log(false);
		}
	}
	return verify;
};
