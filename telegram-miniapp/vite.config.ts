import tailwindcss from "@tailwindcss/vite"
import react from '@vitejs/plugin-react';
import {defineConfig} from 'vite';
import tsconfigPaths from 'vite-tsconfig-paths';
import mkcert from 'vite-plugin-mkcert';

// https://vitejs.dev/config/
export default defineConfig({
    css: {
        preprocessorOptions: {
            scss: {
                api: 'modern',
            },
        },
    },
    plugins: [
        // Allows using React dev server along with building a React application with Vite.
        // https://npmjs.com/package/@vitejs/plugin-react-swc
        react(),
        tailwindcss(),
        // Allows using the compilerOptions.paths property in tsconfig.json.
        // https://www.npmjs.com/package/vite-tsconfig-paths
        tsconfigPaths(),
        process.env.HTTPS && mkcert(),
    ],
    build: {
        target: 'esnext',
        minify: 'terser'
    },
    publicDir: './public',
    server: {
        // Exposes your dev server and makes it accessible for the devices in the same network.
        host: true,
    },
});
