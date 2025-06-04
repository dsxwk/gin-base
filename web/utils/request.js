import {NextLoading} from '/@/utils/loading';
import {ElMessage} from 'element-plus';
import {Session} from '/@/utils/storage';
import pnotify from '/@/utils/pnotify/alert.js';
import pnotifyConfirm from '/@/utils/pnotify/confirm.js';
import {sprintf} from '/@/utils/commonFunction.js';

const API_URL = import.meta.env.VITE_API_URL;

// golang 测试地址
// API_URL = 'http://127.0.0.1:8080/api/v1';

export default async function request(path, config) {
    NextLoading.start();

    let response, res;

    try {
        response = await fetch(API_URL + path, config);
    } catch (error) {
        if (error.message === 'Failed to fetch') {
            ElMessage.error('获取接口失败');
        } else if (error.message.includes('timeout')) {
            ElMessage.error('网络超时');
        } else if (error.message === 'Network Error') {
            ElMessage.error('网络连接错误');
        } else {
            ElMessage.error(error.message);
        }
        NextLoading.done();

        return Promise.reject(error);
    }

    try {
        res = await response.json();
    } catch (error) {
        if (response.status === 404) {
            ElMessage.error('404 Not Found,接口地址不存在');
        } else {
            ElMessage.error(sprintf('http错误码:%s %s', response.status, error));
        }

        NextLoading.done();

        return Promise.reject(error);
    }

    if (res?.code !== 0) {
        NextLoading.done();

        if (res?.code === 401) {
            Session.clear();
            pnotifyConfirm.notice(
                '点击确定跳转至登录',
                res?.msg
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

        pnotify.error(res?.msg);

        return Promise.reject(res);
    }

    NextLoading.done();

    return res;
}