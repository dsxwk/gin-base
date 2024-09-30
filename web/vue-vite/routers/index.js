import {createRouter, createWebHistory} from 'vue-router';
import Functions from '@/utils/functions';
// nprogress
import NProgress from '@/utils/nprogress';
import {staticRouters} from '@/routers/staticRouter.js';
import {LOGIN_URL} from '@/config';

const funcs = new Functions();

/**
 * @description 📚 路由参数配置简介
 * @param path ==> 路由菜单访问路径
 * @param name ==> 路由 name (对应页面组件 name, 可用作 KeepAlive 缓存标识 && 按钮权限筛选)
 * @param redirect ==> 路由重定向地址
 * @param component ==> 视图文件路径
 * @param meta ==> 路由菜单元信息
 * @param meta.icon ==> 菜单和面包屑对应的图标
 * @param meta.title ==> 路由标题 (用作 document.title || 菜单的名称)
 * @param meta.activeMenu ==> 当前路由为详情页时，需要高亮的菜单
 * @param meta.isLink ==> 路由外链时填写的访问地址
 * @param meta.isHide ==> 是否在菜单中隐藏 (通常列表详情页需要隐藏)
 * @param meta.isFull ==> 菜单是否全屏 (示例：数据大屏页面)
 * @param meta.isAffix ==> 菜单是否固定在标签页中 (首页通常是固定项)
 * @param meta.isKeepAlive ==> 当前路由是否缓存
 * */
const routers = createRouter({
    history: createWebHistory(),
    routes: [...staticRouters]
});

routers.beforeEach(async (to, from, next) => {
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
routers.onError(error => {
    NProgress.done();
    console.warn('路由错误', error.message);
});

/**
 * @description 路由跳转结束
 * */
routers.afterEach(() => {
    NProgress.done();
});

export default routers;