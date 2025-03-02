const DEV_BASE_API_URL = 'http://localhost:5000';
const PROD_BASE_API_URL = '';

export const BASE_API_URL =
	process.env.NODE_ENV === 'production' ? PROD_BASE_API_URL : DEV_BASE_API_URL;

async function playCategory(category: string) {
	await fetch(`${BASE_API_URL}/play/${category}`, { method: 'POST' });
}

async function stopPlayback() {
	await fetch(`${BASE_API_URL}/stop`, { method: 'POST' });
}

export { playCategory, stopPlayback };

export const prerender = true;
