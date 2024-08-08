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
        } 
    },
    {
        path: '/user',
        component: () => import('@/views/user/index.vue'),
        meta: {
            title: '用户列表'
        }
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