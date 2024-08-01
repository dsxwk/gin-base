import { createApp } from 'vue';
import App from '../App.vue';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
// 全局引入icon
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
// reset style sheet
import "@styles/reset.scss";
// CSS common style sheet
import "@styles/common.scss";
// iconfont css
import "@assets/iconfont/iconfont.scss";
// font css
import "@assets/fonts/font.scss";
// element dark css
import "element-plus/theme-chalk/dark/css-vars.css";
// custom element dark css
import "@styles/element-dark.scss";
// custom element css
import "@styles/element.scss";

import routers from '@routers/router';
import {createRouter, createWebHistory} from "vue-router";

const routes = createRouter({
    history: createWebHistory(),
    routes: routers
});

const app = createApp(App);

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.use(routes);
app.use(ElementPlus);
app.mount('#app');
