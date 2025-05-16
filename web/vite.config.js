import {resolve} from 'path';
import vue from '@vitejs/plugin-vue';
import vueJsx from '@vitejs/plugin-vue-jsx';
import {defineConfig, loadEnv} from 'vite';
import Inspector from 'vite-plugin-vue-inspector';
import vueSetupExtend from 'vite-plugin-vue-setup-extend-plus';
import viteCompression from 'vite-plugin-compression';
import {buildConfig} from '/@/utils/build';
const pathResolve = (dir) => {
    return resolve(__dirname, '.', dir);
};

// https://vite.dev/config/
export default defineConfig((mode) => {
    const env = loadEnv(mode.mode || '', process.cwd());
    const port = env?.VITE_PORT ? parseInt(env.VITE_PORT, 10) : 3000;

    return {
        plugins: [
            vue(),
            vueJsx(),
            Inspector(),
            vueSetupExtend,
            viteCompression,
            JSON.parse(env.VITE_OPEN_CDN) ? buildConfig.cdn() : null
        ],
        root: process.cwd(),
        resolve: {
            alias: {
                '/@': pathResolve('./'),
            },
        },
        build: {
            sourcemap: true, // 启用 Source Maps
            outDir: 'dist',
            chunkSizeWarningLimit: 1500,
            rollupOptions: {
                output: {
                    chunkFileNames: 'assets/js/[name]-[hash].js',
                    entryFileNames: 'assets/js/[name]-[hash].js',
                    assetFileNames: 'assets/[ext]/[name]-[hash].[ext]',
                    manualChunks(id) {
                        if (id.includes('node_modules')) {
                            return id.toString().match(/\/node_modules\/(?!.pnpm)([^\/]*)\//)?.[1] ?? 'vender';
                        }
                    },
                },
                ...(JSON.parse(env.VITE_OPEN_CDN) ? { external: buildConfig.external } : {}),
            },
        },
        base: mode.command === 'serve' ? './' : env.VITE_PUBLIC_PATH,
        server: {
            host: '0.0.0.0',
            port: port,
            open: JSON.parse(env.VITE_OPEN),
            cors: true,
            sourcemap: true, // 开发环境启用 Source Maps
            proxy: {}, // 代理
        },
        optimizeDeps: {
            include: [],
            exclude: ['vue-demi']
        },
        css: {preprocessorOptions: {css: {charset: false}}},
        define: {
            __NEXT_VERSION__: JSON.stringify(process.env.npm_package_version),
            __NEXT_NAME__: JSON.stringify(process.env.npm_package_name),
        },
    };
});