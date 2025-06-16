import { judementSameArr } from '/@/utils/arrayOperation';

const route = el.__vueParentComponent?.appContext.config.globalProperties.$route;
// 只用 meta.authBtnList 判断
const metaAuthBtnList = route.meta?.authBtnList?.map(item => item?.authValue).filter(Boolean) || [];

/**
 * 单个权限验证 auth('btn.add')
 * @param value 权限值
 * @returns 有权限，返回 `true`，反之则反
 */
export function auth(value) {
    return metaAuthBtnList.some((v) => v === value);
}

/**
 * 多个权限验证，满足一个则为 true auths(['btn.add'])
 * @param value 权限值
 * @returns 有权限，返回 `true`，反之则反
 */
export function auths(value) {
    let flag = false;
    metaAuthBtnList.map((val) => {
        value.map((v) => {
            if (val === v) flag = true;
        });
    });
    return flag;
}

/**
 * 多个权限验证，全部满足则为 true authAll(['btn.add'])
 * @param value 权限值
 * @returns 有权限，返回 `true`，反之则反
 */
export function authAll(value) {
    return judementSameArr(value, metaAuthBtnList);
}