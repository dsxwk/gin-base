import toast from '@/utils/toast';
import {API_URL} from '@/config';
import Functions from '@/utils/functions';
import pnotify from '@/utils/pnotify/alert';
import pnotifyConfirm from '@/utils/pnotify/confirm';

const funcs = new Functions();

// golang 测试地址
// API_URL = 'http://127.0.0.1:8080/api/v1';

export default async function request(path, config) {
    toast.loading();

    const response = await fetch(API_URL + path, config);

    const data = await response.json();

    if (data.code !== 0) {
        toast.clear();

        if (data.code === 401) {
            let redirect = location.pathname + location.search;
            pnotifyConfirm(funcs.lang('Login unauthorized or expired, please log in again'), 'error').then(res => {
                if (res) {
                    location.href = '/login?lang=' + funcs.getLang() + `&redirect=${funcs.urlencode(redirect)}`;
                } else {
                    console.log('取消');
                }
            });
            return;
        }

        if (data.code === 400) {
            pnotify(data.msg, 'notice');
            return;
        }

        pnotify(data.msg, 'error');
        // 抛出异常，中断执行
        throw new Error(data.msg);
    }

    toast.clear();

    return data;
}