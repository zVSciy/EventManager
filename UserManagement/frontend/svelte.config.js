import preprocess from "svelte-preprocess";
import adapter from "@sveltejs/adapter-static";

export default {
    preprocess: preprocess(),
    kit: {
        adapter: adapter({
            pages: "build",
            assets: "build",
            fallback: null,
        }),
    },
};
