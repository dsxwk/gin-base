import createService from '/@/utils/service.js';
/**
 * 菜单
 *
 * @method list 菜单列表
 */
export function menuApi() {
    return createService(
        {
            list: {
                name: '菜单列表',
                url: '/menu',
                method: 'get',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            create: {
                name: '创建菜单',
                url: '/menu',
                method: 'post',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            update: {
                name: '更新菜单',
                url: '/menu/:id',
                method: 'put',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            delete: {
                name: '删除菜单',
                url: '/menu/:id',
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
            actionList: {
                name: '功能列表',
                url: '/menu/:menuId/action',
                method: 'get',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            createAction: {
                name: '新增功能',
                url: '/menu/:menuId/action',
                method: 'post',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            updateAction: {
                name: '修改功能',
                url: '/menu/:menuId/action/:id',
                method: 'put',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            deleteAction: {
                name: '删除功能',
                url: '/menu/:menuId/action/:id',
                method: 'delete',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            actionDetail: {
                name: '功能详情',
                url: '/menu/:menuId/action/:id',
                method: 'get',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            }
        }
    );
}