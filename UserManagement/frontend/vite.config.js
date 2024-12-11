import { defineConfig } from 'vite';
import { sveltekit } from '@sveltejs/kit/vite';

export default defineConfig({
  plugins: [sveltekit()],
  server: {
    host: '0.0.0.0',
    watch: {
      usePolling: true
    }
  }
});

// import { sveltekit } from "@sveltejs/kit/vite";

// export default {
//     plugins: [sveltekit()],
//     server: {
//         port: 3000,
//         strictPort: true,
//     },
// };
