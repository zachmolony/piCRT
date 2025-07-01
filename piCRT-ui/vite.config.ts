import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit(), tailwindcss()],
	server: {
		proxy: {
			'/categories': 'http://localhost:5000',
			'/videos': 'http://localhost:5000',
			'/play': 'http://localhost:5000',
			'/stop': 'http://localhost:5000',
			'/nowplaying': 'http://localhost:5000'
		}
	}
});
