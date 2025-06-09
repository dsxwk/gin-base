// import {ElNotification} from 'element-plus';
import pnotify from '/@/utils/pnotify/alert.js';
import pnotifyConfirm from '/@/utils/pnotify/confirm.js';
import {Session} from '/@/utils/storage';
import {ElMessage} from 'element-plus';

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
        URIError: 'URI错误'
    };

    if (error && error.code) {
        switch (error.code) {
            case 400:
                errorMap[error.code] = '请求错误';
                break;
            case 401:
                errorMap[error.code] = '请求未授权';
                break;
            case 403:
                errorMap[error.code] = '禁止访问';
                break;
            case 404:
                errorMap[error.code] = '请求未找到';
                break;
            case 500:
                errorMap[error.code] = 'System Error';
                break;
            default:
                break;
        }
    }

    let errorName = errorMap[error.name] || errorMap[error.code] || '未知错误';

    pnotify.error(
        typeof error === 'string' ? error : error?.msg,
        errorName
    );

    if (error?.code === 401) {
        Session.clear();
        pnotifyConfirm.notice(
            typeof error === 'string' ? error : error?.msg,
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