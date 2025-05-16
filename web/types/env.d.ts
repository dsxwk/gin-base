/// <reference types="vite/client" />

interface ImportMetaEnv {
    readonly VITE_API_URL: string; // 定义环境变量类型
    readonly VITE_ROUTER_MODE: string; // 定义环境变量类型
    readonly VITE_TOKEN_NAME: string; // token键名
    readonly VITE_V_CONSOLE: string; // 控制台
}

interface ImportMeta {
    readonly env: ImportMetaEnv;
}