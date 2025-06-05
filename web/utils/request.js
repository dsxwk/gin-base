import {NextLoading} from '/@/utils/loading';
import {ElMessage} from 'element-plus';
import {Session} from '/@/utils/storage';
import pnotify from '/@/utils/pnotify/alert.js';
import pnotifyConfirm from '/@/utils/pnotify/confirm.js';

const API_URL = import.meta.env.VITE_API_URL;

// golang 测试地址
// API_URL = 'http://127.0.0.1:8080/api/v1';

export default async function request(path, config) {
    NextLoading.start();

    const response = await fetch(API_URL + path, config).catch((error) => {
        if (error.message === 'Failed to fetch') {
            pnotify.error('获取接口失败');
        } else if (error.message.indexOf('timeout') !== -1) {
            pnotify.error('网络超时');
        } else if (error.message === 'Network Error') {
            pnotify.error('网络连接错误');
        }
    });

    const res = await response?.json();
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

        return Promise.reject(res);
    }

    NextLoading.done();

    return res;
}