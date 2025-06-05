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
            captcha: {
                name: '获取验证码',
                url: '/captcha',
                method: 'get',
                params: {},
                token: {}
            },
            verify: {
                name: '验证验证码',
                url: '/captcha',
                method: 'post',
                params: {},
                token: {}
            }
        }
    );
}