import Functions from '@/utils/functions';
import {HOME_URL} from '@/config';

const funcs = new Functions();

/**
 * 静态路由
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
export const staticRouters = [
    {
        path: '/',
        redirect: HOME_URL
    },
    {
        path: '/login',
        component: () => import('@/views/login/login.vue'),
        meta: {
            title: funcs.lang('Login')
        }
    },
    {
        path: '/layouts',
        redirect: HOME_URL,
        component: () => import('@/layouts/index.vue'),
        meta: {
            title: funcs.lang('Home')
        },
        children: []
    },
    /*{
        path: '/layouts',
        redirect: HOME_URL,
        component: () => import('@/layouts/index.vue'),
        meta: {
            title: funcs.lang('Home')
        },
        children: [
            {
                path: '/dashboard',
                component: () => import('@/views/home/index.vue'),
            }
        ]
    },
    {
        path: '/layouts',
        redirect: '/article',
        component: () => import('@/layouts/index.vue'),
        meta: {
            title: funcs.lang('Article List')
        },
        children: [
            {
                path: '/article',
                component: () => import('@/views/article/index.vue'),
            }
        ]
    },
    {
        path: '/layouts',
        redirect: '/user',
        component: () => import('@/layouts/index.vue'),
        meta: {
            title: funcs.lang('User List')
        },
        children: [
            {
                path: '/user',
                component: () => import('@/views/user/index.vue'),
            }
        ]
    },
    {
        path: '/layouts',
        redirect: '/menu',
        component: () => import('@/layouts/index.vue'),
        meta: {
            title: funcs.lang('Menu List')
        },
        children: [
            {
                path: '/menu',
                component: () => import('@/views/system/menu/index.vue'),
            }
        ]
    },*/
    // 异常404路由
    {
        path: '/:catchAll(.*)',
        component: () => import('@/layouts/errPage/404.vue'),
        meta: {
            title: '404 Not Found'
        }
    }
];