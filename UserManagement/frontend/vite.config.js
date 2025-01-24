import { sveltekit } from '@sveltejs/kit/vite';

export default {
    plugins: [sveltekit()],
    server: {
        proxy: {
            '/event': {
                target: 'http://localhost:8000',
                changeOrigin: true,
                secure: false
            }
        }
    }
};