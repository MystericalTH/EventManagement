/**
 * Lazy-inject the MSW handler
 * so that no errors happen during
 * build/runtime due to invalid
 * imports from server/client.
 */
import { browser, dev } from '$app/environment';
export async function inject() {
	console.log('dev: ' + dev);
	console.log('browser: ' + browser);
	if (dev && browser) {
		console.log('BROWSER REFERRED');

		const { worker } = await import('./browser');
		// For live development, I disabled all warnings
		// for requests that are not mocked. Change how
		// you think it best fits your project.
		return worker.start({ onUnhandledRequest: 'bypass' }).catch(console.warn);
	}
}
