import Functions from '@/utils/functions';
// 定义路由配置
import {HOME_URL} from '@/config';

const funcs = new Functions();
const routers = [
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
    },
    // 可以继续添加其他路由配置
    {
        path: '/:catchAll(.*)',
        component: import('@/layouts/errPage/404.vue'),
        meta: {
            title: '404 Not Found'
        }
    },
];

export default routers;