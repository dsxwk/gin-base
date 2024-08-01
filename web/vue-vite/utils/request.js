import toast from "@utils/toast";
import confs from "@config/configs";
import pnotify from "@utils/pnotify/alert";

let apiUrl = confs.apiUrl;
let data = null;

// golang 测试地址
// apiUrl = 'http://127.0.0.1:8080/api/v1';

export default async function request(path, config) {
    toast.loading();

    const response = await fetch(apiUrl + path, config);

    data = await response.json();

    if (data.code !== 0) {
        toast.clear();

        pnotify(data.msg, 'error');
        // 抛出异常，中断执行
        throw new Error(data.msg);
    }

    toast.clear();

    return data;
}