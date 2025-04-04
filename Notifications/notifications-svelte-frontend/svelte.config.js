import adapter from '@sveltejs/adapter-auto';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter(),
		paths: {	// set base path of SvelteKit application
		base: '/app_notification' // fixes any issues with redirect or routing (over the proxy)
		}
	}
};

export default config;
