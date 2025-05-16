import { createApp } from 'vue';
import App from './App.vue';
import router from '/@/router';
import errorHandler from '/@/utils/error/handle';
import pinia from '/@/stores/index';
import ElementPlus from 'element-plus';
import { i18n } from '/@/static/i18n';
import { directive } from '/@/directive';
import other from '/@/utils/other';
import '/@/theme/index.scss';
import * as ElementPlusIconsVue from '@element-plus/icons-vue';

const app = createApp(App);
directive(app);
other.elSvg(app);

// 全局注册所有图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component);
}
app.config.errorHandler = errorHandler;

if (import.meta.env.VITE_V_CONSOLE === 'true' && /Mobi|Android/i.test(navigator.userAgent)) {
    import("vconsole").then((module) => {
        new module.default();
    });
}

app.
use(pinia).
use(ElementPlus).
use(router).
use(i18n).
mount('#app');
