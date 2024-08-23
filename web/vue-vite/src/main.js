import { createApp } from 'vue';
import App from '../App.vue';
import Functions from '@/utils/functions';
import {LOGIN_URL} from '@/config';
import TablePlus from 'icourt-table';
// reset style sheet
import "@/styles/reset.scss";
// CSS common style sheet
import "@/styles/common.scss";
// iconfont css
import "@/styles/assets/iconfont/iconfont.scss";
// font css
import "@/styles/assets/fonts/font.scss";
// element css
import 'element-plus/dist/index.css';
// custom element css
import "@/styles/element.scss";
// element plus
import ElementPlus from 'element-plus';
// element icons
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
// nprogress
import NProgress from '@/utils/nprogress';
// errorHandler
import errorHandler from '@/utils/errorHandler';

import routers from '@/routers/router';
import {createRouter, createWebHistory} from 'vue-router';

const funcs = new Functions();
const routes = createRouter({
    history: createWebHistory(),
    routes: routers
});

routes.beforeEach((to, from, next) => {
    NProgress.start();
    // 动态标题
    document.title = to.meta.title || funcs.lang('Manage Backend');

    // 设置语言 to = {...to, query: {...to.query, lang: funcs.getLang()}};
    to.fullPath = funcs.setUrlParam(to.fullPath, 'lang', funcs.getLang());
    from.fullPath = funcs.setUrlParam(from.fullPath, 'lang', funcs.getLang());

    // 判断是访问登陆页登录了就在当前页面,没有重置路由到登陆页
    if (to.path.toLocaleLowerCase() === LOGIN_URL) {
        if (funcs.checkLogin()) {
            return next(from.fullPath)
        } else {
            return next(to.fullPath)
        }
    }

    next();
});

/**
 * @description 路由跳转错误
 * */
routes.onError(error => {
    NProgress.done();
    console.warn('路由错误', error.message);
});

/**
 * @description 路由跳转结束
 * */
routes.afterEach(() => {
    NProgress.done();
});

const app = createApp(App);

app.config.errorHandler = errorHandler;

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

if (import.meta.env.VITE_V_CONSOLE === 'true' && /Mobi|Android/i.test(navigator.userAgent)) {
    import("vconsole").then((module) => {
        new module.default();
    });
}

app.use(routes);
app.use(ElementPlus);
app.use(TablePlus);
app.mount('#app');
