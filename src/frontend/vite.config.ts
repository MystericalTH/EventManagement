import { defineConfig } from 'vitest/config';
import { loadEnv } from 'vite';
import { sveltekit } from '@sveltejs/kit/vite';
import * as path from 'path';
import { handlers } from './src/mocks/handlers';
import msw from '@iodigital/vite-plugin-msw';

export default defineConfig(({ command, mode }) => {
	// Load env file based on `mode` in the current working directory.
	// Set the third parameter to '' to load all env regardless of the `VITE_` prefix.
	const env = loadEnv(mode, process.cwd(), '');
	return {
		plugins: [sveltekit(), msw({ handlers })],

		test: {
			include: ['src/**/*.{test,spec}.{js,ts}'],
			exclude: ['**/browser.ts', '**/handler.ts']
		},
		resolve: {
			alias: {
				$msw: path.resolve('./src/mocks'),
				$lib: path.resolve('./src/lib')
			}
		},
		server: {
			proxy: {
				'/api': `${env.VITE_API_DOMAIN}`
			}
		}
	};
});
