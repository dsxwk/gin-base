/**
 * 是否外链
 * @type {[{label: string, value: boolean},{label: string, value: boolean}]}
 */
export const isLinkDict = [
    { label: '是', value: true },
    { label: '否', value: false }
];

/**
 * 是否隐藏
 * @type {[{label: string, value: boolean},{label: string, value: boolean}]}
 */
export const isHideDict = [
    { label: '隐藏', value: true },
    { label: '不隐藏', value: false }
];

/**
 * 是否为链接
 * @type {[{label: string, value: number},{label: string, value: number}]}
 */
export const actionIsLinkEnum = [
    { label: '是', value: true },
    { label: '否', value: false }
];

/**
 * 功能类型
 * @type {[{label: string, value: number},{label: string, value: number}]}
 */
export const actionTypeDict = [
    { label: 'header', value: 1 },
    { label: 'operation', value: 2 },
];

/**
 * 页面缓存
 * @type {[{label: string, value: boolean},{label: string, value: boolean}]}
 */
export const isKeepAliveDict = [
    { label: '缓存', value: true },
    { label: '不缓存', value: false },
];

/**
 * 是否固定
 * @type {[{label: string, value: boolean},{label: string, value: boolean}]}
 */
export const isAffixDict = [
    { label: '固定', value: true },
    { label: '不固定', value: false },
];

/**
 * 是否内嵌
 * @type {[{label: string, value: boolean},{label: string, value: boolean}]}
 */
export const isIframeDict = [
    { label: '是', value: true },
    { label: '否', value: false },
];

// 按钮类型
export const btnTypeDict = [
    { label: '按钮', value: 'btn' },
    { label: '文本', value: 'text' },
];

// 是否确认
export const isConfirmDict = [
    { label: '是', value: 1 },
    { label: '否', value: 2 },
];

// 按钮样式
export const btnStyleDict = [
    { label: 'primary', value: 'primary' },
    { label: 'success', value: 'success' },
    { label: 'warning', value: 'warning' },
    { label: 'danger', value: 'danger' },
    { label: 'info', value: 'info' },
    { label: 'default', value: 'default' },
];

// 按钮尺寸
export const btnSizeDict = [
    { label: 'small', value: 'small' },
    { label: 'default', value: 'default' },
];