// 导入需要的组件
import Index from '@views/layouts/index.vue';
import Login from '@views/login/login.vue';
import UserIndex from '@views/user/index.vue';

import NotFound from '@views/errPage/404.vue'

// 定义路由配置
const routers = [
    {
        path: '/login',
        component: Login,
        meta: {
            title: '登录'
        }
    },
    { 
        path: '/', 
        component: Index,
        meta: { 
            title: '首页' 
        } 
    },
    {
        path: '/user',
        component: UserIndex,
        meta: {
            title: '用户列表'
        }
    },

    // 可以继续添加其他路由配置
    { 
        path: '/:catchAll(.*)', 
        component: NotFound,
        meta: { 
            title: '404 Not Found' 
        } 
    },
];

export default routers;