import Cookies from 'js-cookie';
import data from '@utils/lang/lang.json';
// import {useRoute} from "vue-router";

export class Functions {
    /**
     * 设置cookie
     *
     * @param name
     * @param value
     * @param expires
     */
    setCookie(name, value, expires = 1) {
        Cookies.set(name, value, {expires: expires, path: '/'})
    };

    /**
     * 获取cookie
     *
     * @param name
     * @returns {*}
     */
    getCookie(name) {
        return Cookies.get(name)
    };

    /**
     * 删除cookie
     *
     * @param name
     */
    delCookie(name) {
        Cookies.remove(name, {path: '/'})
    };

    /**
     * 清除本地存储
     */
    clearLocalStorage() {
        localStorage.clear()
    };

    /**
     * 删除本地存储
     *
     * @param name
     */
    removeLocalStorage(name) {
        localStorage.removeItem(name)
    };

    /**
     * 获取本地存储
     *
     * @param name
     * @returns {string}
     */
    getLocalStorage(name) {
        return localStorage.getItem(name)
    };

    /**
     * 设置本地存储
     *
     * @param name
     * @param value
     */
    setLocalStorage(name, value) {
        localStorage.setItem(name, value)
    };

    /**
     * 清空所有缓存
     */
    clearAllCache() {
        clearLocalStorage();
        delCookie('loginInfo');
    };

    /**
     * urlencode
     * @param str
     * @returns {string}
     */
    urlencode (str) {
        str = (str + '').toString();

        return encodeURIComponent(str).replace(/!/g, '%21').replace(/'/g, '%27').replace(/\(/g, '%28').
        replace(/\)/g, '%29').replace(/\*/g, '%2A').replace(/%20/g, '+');
    }

    /**
     * 获取url参数
     *
     * @param key
     * @returns {string|null}
     */
    getUrlParam(key) {
        let url = window.location.search;
        if (url.indexOf('?') > -1) {
            var query = url.substring(url.indexOf('?') + 1);
            var vars = query.split("&");
            for (var i = 0; i < vars.length; i++) {
                var name = vars[i].split("=");
                if (name[0] == key) {
                    return name[1]
                }
            }
        } else {
            return null
        }
    };

    /**
     * 设置url参数
     *
     * @param url
     * @param key
     * @param val
     * @returns {boolean}
     */
    setUrlParam(url, key, val) {
        if (url != 'javascript:void(0);') {
            if (url.indexOf(key) > -1) {
                var replace = eval('/(' + key + '=)([^&]*)/gi');
                url = url.replace(replace, key + '=' + val)
            } else {
                var paraStr = key + '=' + val;

                var idx = url.indexOf('?');

                if (idx < 0) {
                    url = url + '?' + paraStr;
                } else if (idx >= 0 && idx != url.length - 1) {
                    url = url + '&' + paraStr;
                }
            }

            return url;
        } else {
            return url;
        }
    };

    /**
     * 获取语言
     *
     * @returns {*}
     */
    getLang() {
        if (this.getUrlParam('lang') == null) {
            this.setCookie('lang', 'zh-cn');
        } else {
            this.setCookie('lang', this.getUrlParam('lang'));
        }

        return this.getCookie('lang');
    };

    /**
     * 添加语言到url
     *
     * @param url
     * @param lang
     * @returns {string|*}
     */
    addLang(url, lang) {
        if (url != 'javascript:void(0);') {
            lang = (lang == '') ? 'zh-cn' : lang;
            let currentLang = 'lang=' + lang;
            if (url.search) {
                let arrQuery = _.remove(url.search.substring(1).split('&'), function (el) {
                    return el.split('=')[0] != 'lang'
                });
                let query = '';
                if (arrQuery.join('&') == '') {
                    query = ['lang=' + lang].join('&')
                } else {
                    query = ['lang=' + lang, arrQuery.join('&')].join('&')
                }
                if (url == null) return 'javascript:void(0);';
                return url.split('?')[0] + '?' + query
            } else {
                let query = ['lang=' + lang].join('&');
                return url.split('?')[0] + '?' + query
            }
        } else {
            return url
        }
    };

    /**
     * 切换语言
     *
     * @param lang
     */
    /*switchLang(lang)
    {
        let route = useRoute();
        funcs.setCookie('lang', lang);

        // 获取当前路由地址和查询参数
        let currentPath = route.path;
        let currentQuery = route.query;

        // 构造查询参数对象
        let query = new URLSearchParams();
        for (let key in currentQuery) {
            query.set(key, currentQuery[key]);
        }
        query.set('lang', lang);

        // 构造新的 URI 地址
        location.href = `${currentPath}?${query.toString()}`;
    }*/

    /**
     *
     * 翻译
     *
     * @param key
     * @returns {*}
     */
    lang(key) {
        let lang = this.getLang();

        if (data[key]) {
            return data[key][lang];
        } else {
            return key;
        }
    };
}