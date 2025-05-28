import createService from '/@/utils/service.js';

/**
 * 角色
 *
 * @method list 角色列表
 */
export function roleApi() {
    return createService(
        {
            list: {
                name: '角色列表',
                url: '/role',
                method: 'get',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            create: {
                name: '创建角色',
                url: '/role',
                method: 'post',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            update: {
                name: '更新角色',
                url: '/role/:id',
                method: 'put',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            delete: {
                name: '删除角色',
                url: '/role/:id',
                method: 'delete',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
        }
    );
}