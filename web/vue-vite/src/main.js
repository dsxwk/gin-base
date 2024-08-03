import { createApp } from 'vue';
import App from '../App.vue';
// reset style sheet
import "@styles/reset.scss";
// CSS common style sheet
import "@styles/common.scss";
// iconfont css
import "@assets/iconfont/iconfont.scss";
// font css
import "@assets/fonts/font.scss";
// element css
import 'element-plus/dist/index.css';
// element dark css
import "element-plus/theme-chalk/dark/css-vars.css";
// element dark css
import "element-plus/theme-chalk/dark/css-vars.css";
// custom element dark css
import "@styles/element-dark.scss";
// custom element css
import "@styles/element.scss";
// element plus
import ElementPlus from 'element-plus';
// element icons
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
import NProgress from '@utils/nprogress';

import routers from '@routers/router';
import {createRouter, createWebHistory} from "vue-router";

const routes = createRouter({
    history: createWebHistory(),
    routes: routers
});

routes.beforeEach((to, from, next) => {
    NProgress.start();
    // 动态标题
    document.title = to.meta.title || '后台管理';
    next();
});

/**
 * @description 路由跳转错误
 * */
routes.onError(error => {
    NProgress.done();
    console.warn("路由错误", error.message);
});

/**
 * @description 路由跳转结束
 * */
routes.afterEach(() => {
    NProgress.done();
});

const app = createApp(App);

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.use(routes);
app.use(ElementPlus);
app.mount('#app');
