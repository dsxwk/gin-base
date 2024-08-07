import { resolve } from 'path';
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
// import { createProxy } from './utils/proxy';

import dotenv from 'dotenv';

// https://vitejs.dev/config/
export default defineConfig(({mode}) => {
  // 加载对应环境的配置文件
  const { parsed: env } = dotenv.config({ path: `.env.${mode}` }) || {};

  const port = env?.VITE_PORT ? parseInt(env.VITE_PORT, 10) : 3000;
  const open = env?.VITE_OPEN === 'true';
  // const proxy = env?.VITE_PROXY ? createProxy(env.VITE_PROXY) : {};

  return {
    plugins: [
      vue(),
    ],
    resolve: {
      alias: {
        '@': resolve(__dirname, './'),
      }
    },
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: `@import '@/styles/var.scss';`
        }
      }
    },
    server: {
      host: '0.0.0.0',
      port: port,
      open: open,
      // 启用cros
      cors: true,
      // 配置代理
      // proxy: proxy
    },
    // plus
    optimizeDeps: {
      include: []
    }
  }
});
