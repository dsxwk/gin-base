import createService from '/@/utils/service.js';

/**
 * 字典
 *
 * @method list 字典列表
 */
export function dictApi() {
    return createService(
        {
            list: {
                name: '字典列表',
                url: '/dict',
                method: 'get',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            create: {
                name: '创建字典',
                url: '/dict',
                method: 'post',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            update: {
                name: '更新字典',
                url: '/dict/:id',
                method: 'put',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            detail: {
                name: '字典详情',
                url: '/dict/:id',
                method: 'get',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            delete: {
                name: '删除字典',
                url: '/dict/:id',
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