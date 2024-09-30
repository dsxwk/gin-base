import {HOME_URL} from '@/config';
import Functions from '@/utils/functions';
import {menuJson} from '@/utils/data/menu';

// 引入 views 文件夹下所有 vue 文件
const modules = import.meta.glob('@/views/**/*.vue');

const funcs = new Functions();

// 递归获取动态路由
export const dynamicRouter = (menuJson = [], arr = []) => {
    for (const item of menuJson) {
        // 动态导入组件
        if (item.component && typeof item.component === "string") {
            item.component = modules['/views' + item.component + '.vue'];
        }
        if (item.path !== '') {
            // 添加当前路由项到数组
            arr.push(item);

            // 如果有子路由，则递归调用
            if (item.children && item.children.length > 0) {
                dynamicRouter(item.children, arr);
            }
        } else if (item.path === '' && item.children && item.children.length > 0) {
            dynamicRouter(item.children, arr);
        }
    }
    return arr;
};

// 处理parent layouts
export const getDynamicRouter = async function () {
    return [{
        path: '/layouts',
        redirect: HOME_URL,
        component: () => import('@/layouts/index.vue'),
        meta: {
            title: funcs.lang('Home')
        },
        children: dynamicRouter(menuJson)
    }];
}