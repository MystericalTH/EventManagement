import type { HandleFetch } from '@sveltejs/kit';
import { browser, dev } from '$app/environment';

if (dev && browser && import.meta.env.VITE_MSW_ENABLED == 'true') {
	console.log('BROWSER REFERRED');

	const { worker } = await import('$msw/browser');
	// For live development, I disabled all warnings
	// for requests that are not mocked. Change how
	// you think it best fits your project.
	worker.start({ onUnhandledRequest: 'bypass' }).catch(console.warn);
}
