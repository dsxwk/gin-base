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
        }
    );
}