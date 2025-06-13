// import {ElNotification} from 'element-plus';
import pnotify from '/@/utils/pnotify/alert.js';
import pnotifyConfirm from '/@/utils/pnotify/confirm.js';
import {Session} from '/@/utils/storage';
import {ElMessage} from 'element-plus';

const recentErrors = new Map();

/**
 * 全局代码错误捕捉
 *
 * @param error
 * @returns {boolean}
 */
const errorHandler = (error) => {
    console.log(error);

    // 过滤 HTTP 请求错误
    if (error.status || error.status === 0) {
        return false;
    }

    let errorMap = {
        InternalError: 'Javascript引擎内部错误',
        ReferenceError: '未找到对象',
        TypeError: '使用了错误的类型或对象',
        RangeError: '使用内置对象时，参数超范围',
        SyntaxError: '语法错误',
        EvalError: '错误的使用了Eval',
        URIError: 'URI错误',
        400: '请求错误',
        401: '请求未授权',
        403: '禁止访问',
        404: '请求未找到',
        500: 'System Error',
    };

    const code = error?.code || error?.name || 'unknown';
    const msg = typeof error === 'string' ? error : error?.msg || '未知错误';
    const errorKey = `${code}-${msg}`;

    // 短时间内重复错误，直接忽略
    if (recentErrors.has(errorKey)) return false;

    recentErrors.set(errorKey, true);
    // 允许再次弹出同类错误
    setTimeout(() => recentErrors.delete(errorKey), 1);

    const errorName = errorMap[error.name] || errorMap[error.code] || '未知错误';

    pnotify.error(
        typeof error === 'string' ? error : error?.msg ? error?.msg : '未知错误',
        errorName
    );

    if (error?.code === 401) {
        Session.clear();
        pnotifyConfirm.notice(
            typeof error === 'string' ? error : error?.msg ? error?.msg : '未知错误',
            '登录已过期, 点击确定跳转至登录',
        ).then(
            () => {
                window.location.href = '/';
            }
        ).catch(
            () => {
                ElMessage.info('已取消');
                setTimeout(function () {
                    window.location.reload();
                }, 100);
            }
        );
    }

    // ElNotification({
    //     title: errorName,
    //     message: typeof error === 'string' ? error : error?.msg,
    //     type: "error",
    //     duration: 3000
    // });
};

export default errorHandler;