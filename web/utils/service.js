import request from '/@/utils/request.js';
import {Session, Local} from '/@/utils/storage';

export default function createService(module, headers = {}) {
    const themeConfig = Local.get('themeConfig');
    let lang = themeConfig.globalI18n ?? 'zh-Cn'
    if (lang.startsWith('en')) {
        lang = 'en-US';
    }
    headers['Accept-Language'] = lang;
    const service = {};
    let hds = {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        ...headers,
    }

    Object.entries(module).forEach(([key, value]) => {
        service[key] = async function (params) {
            // 每次请求都动态获取 token
            let dynamicHeaders = {...hds};
            if (value.token !== undefined && value.token.value) {
                dynamicHeaders[value?.token.name] = Session.get('token');
            }

            const {url, method} = value;
            const config = {
                method,
                headers: dynamicHeaders,
            };
            if (method.toUpperCase() === 'POST' || method.toUpperCase() === 'PUT') {
                config.body = JSON.stringify(params);
            }

            let finalUrl = url;
            // 替换 URL 中的参数
            const matches = url.match(/:(\w+)/g);
            if (matches) {
                matches.forEach((match) => {
                    // 去除 ":"
                    const paramKey = match.slice(1);
                    if (params && params[paramKey] !== undefined) {
                        finalUrl = finalUrl.replace(match, params[paramKey]);
                        delete params[paramKey];
                    }
                });
            }

            // 当 method 为 GET 时才拼接查询参数
            if (method.toUpperCase() === 'GET' && params) {
                let queryParams = new URLSearchParams();
                Object.entries(params).forEach(([paramKey, paramValue]) => {
                    finalUrl = finalUrl.replace(`:${paramKey}`, paramValue);
                    queryParams.append(paramKey, paramValue);
                });

                // 添加判断条件
                if (!isEmptyObject(params)) {
                    const delimiter = finalUrl.includes('?') ? '&' : '?';
                    finalUrl += delimiter + queryParams.toString();
                }
            }

            return request(finalUrl, config);
        };
    });

    return service;
}

const isEmptyObject = (obj) => {
    return Object.keys(obj).length === 0;
};