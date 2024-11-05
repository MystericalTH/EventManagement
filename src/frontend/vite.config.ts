import { defineConfig } from 'vitest/config';
import { sveltekit } from '@sveltejs/kit/vite';
import * as path from 'path';

export default defineConfig({
	plugins: [sveltekit()],

	test: {
		include: ['src/**/*.{test,spec}.{js,ts}'],
		exclude: ['**/browser.ts', '**/handler.ts']
	},
	resolve: {
		alias: {
			$msw: path.resolve('./src/mocks'),
			$lib: path.resolve('./src/lib')
		}
	}
});
