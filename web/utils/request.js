import {NextLoading} from '/@/utils/loading';
import errorHandler from '/@/utils/error/handle';
import {Session} from '/@/utils/storage.js';

const API_URL = import.meta.env.VITE_API_URL;

// golang 测试地址
// API_URL = 'http://127.0.0.1:8080/api/v1';

const err = {
    code : 0,
    msg : ''
};

export default async function request(path, config) {
    NextLoading.start();

    let response, res;
    try {
        response = await fetch(API_URL + path, config);
    } catch (error) {
        NextLoading.done();
        errorHandler(error);

        return Promise.reject(error);
    }

    if (response.status !== 200) {
        NextLoading.done();
        err.code = response.status;
        err.msg = response.statusText;

        return Promise.reject(err);
    }

    try {
        res = await response?.json();
    } catch (error) {
        NextLoading.done();

        return Promise.reject(error);
    }

    if (res?.code !== 0) {
        NextLoading.done();
        if (res?.code === 401) {
            try {
                await refreshToken();
                response = await fetch(API_URL + path, config);
                res = await response?.json();
            } catch (err) {
                errorHandler(res);

                return Promise.reject(res);
            }
        } else {
            errorHandler(res);
            // throw new Error(res?.msg)
            return Promise.reject(res);
        }
    }

    NextLoading.done();

    return res;
}

async function refreshToken() {
    const refreshToken = Session.get('refreshToken');
    if (!refreshToken) throw new Error('没有 refresh token');

    const res = await fetch(API_URL + '/refresh-token', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'token': refreshToken
        }
    });

    if (res.status !== 200) throw new Error('刷新失败');

    const data = await res.json();
    if (data.code !== 0) throw new Error(data.msg);

    Session.set('token', data?.data?.accessToken);
    Session.set('refreshToken', data?.data?.refreshToken);

    return data?.data?.accessToken;
}