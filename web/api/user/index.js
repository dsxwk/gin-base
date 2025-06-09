import createService from '/@/utils/service.js';

/**
 * 用户
 *
 * @method list 用户列表
 */
export function userApi() {
    return createService(
        {
            list: {
                name: '用户列表',
                url: '/user',
                method: 'get',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            create: {
                name: '创建用户',
                url: '/user',
                method: 'post',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            update: {
                name: '更新用户',
                url: '/user/:id',
                method: 'put',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            detail: {
                name: '用户详情',
                url: '/user/:id',
                method: 'get',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            delete: {
                name: '删除用户',
                url: '/user/:id',
                method: 'delete',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            roleList: {
                name: '角色列表',
                url: '/role',
                method: 'get',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
        }
    );
}