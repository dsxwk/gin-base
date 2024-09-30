import { createApp } from 'vue';
import App from '../App.vue';
import {ref} from 'vue';
import TablePlus from 'icourt-table';
import {getDarkColor, getLightColor} from '@/utils/color';
const DEFAULT_PRIMARY = '#009688';
const isDark = ref(false);
// global scss
document.documentElement.style.setProperty('--el-color-primary', DEFAULT_PRIMARY);
for (let i = 1; i <= 9; i++) {
    const primaryColor = isDark.value ? `${getDarkColor(DEFAULT_PRIMARY, i / 10)}` : `${getLightColor(DEFAULT_PRIMARY, i / 10)}`;
    document.documentElement.style.setProperty(`--el-color-primary-light-${i}`, primaryColor);
}
// reset style sheet
import '@/styles/reset.scss';
// CSS common style sheet
import "@/styles/common.scss";
// iconfont css
import '@/styles/assets/iconfont/iconfont.scss';
// font css
import '@/styles/assets/fonts/font.scss';
// element css
import 'element-plus/dist/index.css';
// custom element css
import '@/styles/element.scss';
// element plus
import ElementPlus from 'element-plus';
// element icons
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
// errorHandler
import errorHandler from '@/utils/errorHandler';

import routers from '@/routers';

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

app.use(routers);
app.use(ElementPlus);
app.use(TablePlus);
app.mount('#app');
