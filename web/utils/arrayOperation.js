/**
 * 判断两数组字符串是否相同（用于按钮权限验证），数组字符串中存在相同时会自动去重（按钮权限标识不会重复）
 * @returns 两数组相同返回 `true`，反之则反
 * @param newArr
 * @param oldArr
 */
export function judementSameArr(newArr, oldArr) {
    const news = removeDuplicate(newArr);
    const olds = removeDuplicate(oldArr);
    let count = 0;
    const length = news.length;
    for (let i in olds) {
        for (let j in news) {
            if (olds[i] === news[j]) count++;
        }
    }
    return count === length;
}

/**
 * 判断两个对象是否相同
 * @param a 要比较的对象一
 * @param b 要比较的对象二
 * @returns 相同返回 true，反之则反
 */
export function isObjectValueEqual(a, b) {
    if (!a || !b) return false;
    let aProps = Object.getOwnPropertyNames(a);
    let bProps = Object.getOwnPropertyNames(b);
    if (aProps.length !== bProps.length) return false;
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
 * 数组、数组对象去重
 * @param arr 数组内容
 * @returns
 */
export function removeDuplicate(arr)
{
    if (!Object.keys(arr).length) {
        return arr;
    } else {
        let attr;
        if (attr) {
            const obj = {};
            return arr.reduce((cur, item) => {
                obj[item[attr]] ? '' : (obj[item[attr]] = item[attr] && cur.push(item));
                return cur;
            }, []);
        } else {
            return [...new Set(arr)];
        }
    }
}
