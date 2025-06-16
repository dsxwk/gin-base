import createService from '/@/utils/service.js';

/**
 * 文章
 *
 * @method list 文章列表
 */
export function articleApi() {
    return createService(
        {
            list: {
                name: '文章列表',
                url: '/article',
                method: 'get',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            create: {
                name: '创建文章',
                url: '/article',
                method: 'post',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            update: {
                name: '更新文章',
                url: '/article/:id',
                method: 'put',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            detail: {
                name: '文章详情',
                url: '/article/:id',
                method: 'get',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            },
            delete: {
                name: '删除文章',
                url: '/article/:id',
                method: 'delete',
                params: {},
                token: {
                    name: 'token',
                    value: true
                }
            }
        }
    );
}