import {createRouter, createWebHashHistory, createWebHistory} from 'vue-router';
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import pinia from '/@/stores/index';
import {staticRoutes, notFoundAndNoPower} from '/@/router/route';
import {Session} from '/@/utils/storage';
import {useRoutesList} from '/@/stores/routesList';
import {useKeepALiveNames} from '/@/stores/keepAliveNames';
import {storeToRefs} from 'pinia';
import {initBackEndControlRoutes} from '/@/router/backEnd';
import {i18n} from '/@/static/i18n/index';
import {useThemeConfig} from '/@/stores/themeConfig';

// 读取 `/@/stores/themeConfig.js` 是否开启后端控制路由配置
const storesThemeConfig = useThemeConfig(pinia);
const {themeConfig} = storeToRefs(storesThemeConfig);
const {isRequestRoutes} = themeConfig.value;

/**
 * 创建一个可以被 Vue 应用程序使用的路由实例
 * @method createRouter(options: RouterOptions): Router
 * @link 参考：https://next.router.vuejs.org/zh/api/#createrouter
 */
export const router = createRouter({
    history: import.meta.env.VITE_ROUTER_MODE === 'history' ? createWebHistory() : createWebHashHistory(),
    /**
     * 说明：
     * 1、notFoundAndNoPower 默认添加 404、401 界面，防止一直提示 No match found for location with path 'xxx'
     * 2、backEnd.ts(后端控制路由)、frontEnd.ts(前端控制路由) 中也需要加 notFoundAndNoPower 404、401 界面。
     *    防止 404、401 不在 layout 布局中，不设置的话，404、401 界面将全屏显示
     */
    routes: [...notFoundAndNoPower, ...staticRoutes],
});

/**
 * 路由多级嵌套数组处理成一维数组
 * @param arr 传入路由菜单数据数组
 * @returns 返回处理后的一维路由菜单数组
 */
export function formatFlatteningRoutes(arr) {
    if (arr.length <= 0) return false;
    for (let i = 0; i < arr.length; i++) {
        if (arr[i].children) {
            arr = arr.slice(0, i + 1).concat(arr[i].children, arr.slice(i + 1));
        }
    }
    return arr;
}

/**
 * 一维数组处理成多级嵌套数组（只保留二级：也就是二级以上全部处理成只有二级，keep-alive 支持二级缓存）
 * @description isKeepAlive 处理 `name` 值，进行缓存。顶级关闭，全部不缓存
 * @link 参考：https://v3.cn.vuejs.org/api/built-in-components.html#keep-alive
 * @param arr 处理后的一维路由菜单数组
 * @returns 返回将一维数组重新处理成 `定义动态路由（dynamicRoutes）` 的格式
 */
export function formatTwoStageRoutes(arr) {
    if (arr.length <= 0) return false;
    const newArr = [];
    const cacheList = [];
    arr.forEach((v) => {
        if (v.path === '/') {
            newArr.push({
                component: v.component,
                name: v.name,
                path: v.path,
                redirect: v.redirect,
                meta: v.meta,
                children: []
            });
        } else {
            // 判断是否是动态路由（xx/:id/:name），用于 tagsView 等中使用
            // 修复：https://gitee.com/lyt-top/vue-next-admin/issues/I3YX6G
            if (v.path.indexOf('/:') > -1) {
                v.meta['isDynamic'] = true;
                v.meta['isDynamicPath'] = v.path;
            }
            newArr[0].children.push({...v});
            // 存 name 值，keep-alive 中 include 使用，实现路由的缓存
            // 路径：/@/layouts/routerView/parent.vue
            if (newArr[0].meta.isKeepAlive && v.meta.isKeepAlive) {
                cacheList.push(v.name);
                const stores = useKeepALiveNames(pinia);
                stores.setCacheKeepAlive(cacheList);
            }
        }
    });
    return newArr;
}

// 路由加载前
router.beforeEach(async (to, from, next) => {
    NProgress.configure({showSpinner: false});
    if (to.meta.title) NProgress.start();
    document.title = i18n.global.t(to.meta.title) || 'Gin Base 后台管理';
    const token = Session.get('token');
    if (to.path === '/login' && !token) {
        next();
        NProgress.done();
    } else {
        if (!token) {
            next(`/login?redirect=${to.path}&params=${JSON.stringify(to.query ? to.query : to.params)}`);
            Session.clear();
            NProgress.done();
        } else if (token && to.path === '/login') {
            next('/home');
            NProgress.done();
        } else {
            const storesRoutesList = useRoutesList(pinia);
            const {routesList} = storeToRefs(storesRoutesList);
            if (routesList.value.length === 0) {
                if (isRequestRoutes) {
                    // 后端控制路由：路由数据初始化，防止刷新时丢失
                    await initBackEndControlRoutes();
                    // 解决刷新时，一直跳 404 页面问题，关联问题 No match found for location with path 'xxx'
                    // to.query 防止页面刷新时，普通路由带参数时，参数丢失。动态路由（xxx/:id/:name"）isDynamic 无需处理
                    next({path: to.path, query: to.query});
                }
            } else {
                next();
            }
        }
    }
});

// 路由错误
router.onError(error => {
    NProgress.done();
    console.warn('路由错误', error.message);
});

// 路由加载后
router.afterEach(() => {
    NProgress.done();
});

// 导出路由
export default router;