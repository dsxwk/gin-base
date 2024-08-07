import Cookies from 'js-cookie';
import data from '@utils/lang/lang.json';

export default class Functions {
    /**
     * 设置cookie
     *
     * @param key
     * @param value
     * @param expires
     */
    setCookie(key, value, expires = 1) {
        Cookies.set(key, value, {expires: expires, path: '/'})
    };

    /**
     * 获取cookie
     *
     * @param key
     * @returns {*}
     */
    getCookie(key) {
        return Cookies.get(key)
    };

    /**
     * 删除cookie
     *
     * @param key
     */
    delCookie(key) {
        Cookies.remove(key, {path: '/'})
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
     * @param key
     */
    removeLocalStorage(key) {
        localStorage.removeItem(key)
    };

    /**
     * 获取本地存储
     *
     * @param key
     * @returns {string}
     */
    getLocalStorage(key) {
        const value = localStorage.getItem(key);
        try {
            return JSON.parse(localStorage.getItem(key));
        } catch (error) {
            return value;
        }
    };

    /**
     * 设置本地存储
     *
     * @param key
     * @param value
     */
    setLocalStorage(key, value) {
        localStorage.setItem(key, JSON.stringify(value));
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
     * @description 判断数据类型
     * @param {*} val 需要判断类型的数据
     * @returns {String}
     */
    isType(val) {
        if (val === null) {
            return 'null'
        }
        if (typeof val !== 'object') {
            return typeof val
        } else {
            return Object.prototype.toString.call(val).slice(8, -1).toLocaleLowerCase()
        }
    }

    /**
     * @description 生成唯一 uuid
     * @returns {String}
     */
    generateUUID() {
        let uuid = '';
        for (let i = 0; i < 32; i++) {
            let random = (Math.random() * 16) | 0;
            if (i === 8 || i === 12 || i === 16 || i === 20) uuid += '-';
            uuid += (i === 12 ? 4 : i === 16 ? (random & 3) | 8 : random).toString(16);
        }
        return uuid;
    }

    /**
     * 判断两个对象是否相同
     * @param {Object} a 要比较的对象一
     * @param {Object} b 要比较的对象二
     * @returns {Boolean} 相同返回 true，反之 false
     */
    isObjectValueEqual(a, b) {
        if (!a || !b) return false;
        let aProps = Object.getOwnPropertyNames(a);
        let bProps = Object.getOwnPropertyNames(b);
        if (aProps.length != bProps.length) return false;
        for (let i = 0; i < aProps.length; i++) {
            let propName = aProps[i];
            let propA = a[propName];
            let propB = b[propName];
            if (!b.hasOwnProperty(propName)) return false;
            if (propA instanceof Object) {
                if (!isObjectValueEqual(propA, propB)) return false;
            } else if (propA !== propB) {
                return false;
            }
        }
        return true;
    }

    /**
     * @description 获取浏览器默认语言
     * @returns {String}
     */
    getBrowserLang() {
        let browserLang = navigator.language ? navigator.language : navigator.browserLanguage;
        let defaultBrowserLang = '';
        if (['cn', 'zh', 'zh-cn'].includes(browserLang.toLowerCase())) {
            defaultBrowserLang = 'zh-cn';
        } else {
            defaultBrowserLang = 'en-us';
        }
        return defaultBrowserLang;
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
            var vars = query.split('&');
            for (var i = 0; i < vars.length; i++) {
                var name = vars[i].split('=');
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