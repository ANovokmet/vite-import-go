import { defineConfig } from 'vite';
import { viteImportGoPlugin } from './lib/vite-import-go-plugin';

export default defineConfig({
	plugins: [viteImportGoPlugin()],
});
