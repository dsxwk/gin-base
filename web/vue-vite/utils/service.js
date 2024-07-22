export default function createService(module, requestFunc, headers = {}) {
    const service = {};
    let hds = {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        ...headers,
    }

    Object.entries(module).forEach(([key, value]) => {
        if (value.token !== undefined && value.token.value) {
            hds[value?.token.name] = window.localStorage.getItem('token');
        }

        service[key] = async function (params) {
            const {url, method} = value;
            const config = {
                method,
                headers: hds,
            };
            if (method.toUpperCase() === 'POST' || method.toUpperCase() === 'PUT') {
                config.body = JSON.stringify(params);
            }

            let finalUrl = url;
            // 替换 URL 中的参数
            const paramRegex = /:(\w+)/g;
            const matches = url.match(paramRegex);
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

            return requestFunc(finalUrl, config);
        };
    });

    return service;
}

const isEmptyObject = (obj) => {
    return Object.keys(obj).length === 0;
};