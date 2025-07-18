// import adapter from '@sveltejs/adapter-static';

// const config = {
// 	kit: {
// 		adapter: adapter({
// 			pages: 'build',
// 			assets: 'build',
// 			fallback: 'app.html',    // Custom fallback
// 			precompress: false,
// 			strict: false
// 		})
// 	}
// };

// export default config;

import adapter from '@sveltejs/adapter-netlify';
import { vitePreprocess } from '@sveltejs/kit/vite';

const config = {
	preprocess: vitePreprocess(),
	kit: {
		adapter: adapter()
	}
};

export default config;
