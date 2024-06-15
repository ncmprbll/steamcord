import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import fs from 'fs';

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		port: 443,
        https: {
            key: fs.readFileSync('/etc/letsencrypt/live/project.pyroman.io/privkey.pem'),
            cert: fs.readFileSync('/etc/letsencrypt/live/project.pyroman.io/fullchain.pem')
        },
        proxy: {}
	}
});
