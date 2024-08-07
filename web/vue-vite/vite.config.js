// import {fileURLToPath, URL} from 'node:url';
import { resolve } from "path";
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

import dotenv from 'dotenv';

// https://vitejs.dev/config/
export default defineConfig(({mode}) => {
  dotenv.config({path: `.env.${mode}`}) // 加载对应环境的配置文件

  return {
    plugins: [
      vue(),
    ],
    resolve: {
      alias: {
        /*'@': fileURLToPath(new URL('./src', import.meta.url)),
        // 设置路径别名
        '@app': '/app',
        '@views': '/views',
        '@routers': '/routers',
        '@modules': '/app/modules',
        '@services': '/app/services',
        '@components': '/components',
        '@layouts': '/views/layouts',
        '@utils': '/utils',
        '@config': '/config',
        '@assets': '/src/assets',
        '@styles': '/src/styles',*/
        "@": resolve(__dirname, "./"),
      }
    },
  }
});
