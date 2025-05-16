import createService from '/@/utils/service.js';

/**
 * 登录
 *
 * @method login 用户登录
 */
export function loginApi() {
    return createService(
        {
            login: {
                name: '用户登录',
                url: '/login',
                method: 'post',
                params: {},
                token: {}
            },
        }
    );
}