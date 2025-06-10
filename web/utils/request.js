import {NextLoading} from '/@/utils/loading';
import errorHandler from '/@/utils/error/handle';

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
        errorHandler(err);

        return Promise.reject(err);
    }

    try {
        res = await response?.json();
    } catch (error) {
        NextLoading.done();
        errorHandler(error);

        return Promise.reject(error);
    }

    if (res?.code !== 0) {
        NextLoading.done();

        // throw new Error(res?.msg)
        return Promise.reject(res);
    }

    NextLoading.done();

    return res;
}