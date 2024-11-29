import { sveltekit } from "@sveltejs/kit/vite";

export default {
    plugins: [sveltekit()],
    server: {
        port: 3000,
        strictPort: true,
    },
};
