// 定义路由配置
const routers = [
    {
        path: '/login',
        component: () => import('@/views/login/login.vue'),
        meta: {
            title: '登录'
        }
    },
    {
        path: '/',
        component: () => import('@/layouts/index.vue'),
        meta: {
            title: '首页'
        },
        children: [
            {
                path: '/home',
                component: () => import('@/views/home/index.vue'),
            }
        ]
    },
    {
        path: '',
        component: () => import('@/layouts/index.vue'),
        meta: {
            title: '用户列表'
        },
        children: [
            {
                path: '/user',
                component: () => import('@/views/user/index.vue'),
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