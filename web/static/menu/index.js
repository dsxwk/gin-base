/**
 * 建议：路由 path 路径与文件夹名称相同，找文件可浏览器地址找，方便定位文件位置
 *
 * 路由meta对象参数说明
 * meta: {
 *      title:          菜单栏及 tagsView 栏、菜单搜索名称（国际化）
 *      isLink：        是否超链接菜单，开启外链条件，`1、isLink: 链接地址不为空 2、isIframe:false`
 *      isHide：        是否隐藏此路由
 *      isKeepAlive：   是否缓存组件状态
 *      isAffix：       是否固定在 tagsView 栏上
 *      isIframe：      是否内嵌窗口，开启条件，`1、isIframe:true 2、isLink：链接地址不为空`
 *      roles：         当前路由权限标识，取角色管理。控制路由显示、隐藏。超级管理员：admin 普通角色：common
 *      icon：          菜单、tagsView 图标，阿里：加 `iconfont xxx`，fontawesome：加 `fa xxx`
 * }
 */
export const menuJson = {
    "code": 0,
    "type": "adminMenu",
    "data": [
        {
            "id": 1,
            "pid": 0,
            "path": "/home",
            "name": "home",
            "component": "home/index",
            "meta": {
                "title": "message.router.home",
                "isLink": "",
                "isHide": false,
                "isKeepAlive": true,
                "isAffix": true,
                "isIframe": false,
                "roles": [
                    "admin",
                    "common"
                ],
                "icon": "iconfont icon-shouye"
            },
            "title": "首页",
            "menuType": "menu",
            "children": []
        },
        {
            "id": 2,
            "pid": 0,
            "path": "/system",
            "name": "system",
            "component": "layouts/routerView/parent",
            "redirect": "/system/menu",
            "meta": {
                "title": "message.router.system",
                "isLink": "",
                "isHide": false,
                "isKeepAlive": true,
                "isAffix": false,
                "isIframe": false,
                "roles": [
                    "admin"
                ],
                "icon": "iconfont icon-xitongshezhi",
            },
            "children": [
                {
                    "id": 3,
                    "pid": 0,
                    "path": "/system/menu",
                    "name": "systemMenu",
                    "component": "system/menu/index",
                    "meta": {
                        "title": "message.router.systemMenu",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin"
                        ],
                        "icon": "iconfont icon-caidan"
                    },
                    "title": "菜单管理",
                    "children": []
                },
                {
                    "id": 4,
                    "pid": 0,
                    "path": "/system/user",
                    "name": "systemUser",
                    "component": "system/user/index",
                    "meta": {
                        "title": "message.router.systemUser",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin"
                        ],
                        "icon": "iconfont icon-icon-"
                    },
                    "title": "用户管理",
                    "children": []
                },
                {
                    "id": 5,
                    "pid": 0,
                    "path": "/system/role",
                    "name": "systemRole",
                    "component": "system/role/index",
                    "meta": {
                        "title": "message.router.systemRole",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": ["admin"],
                        "icon": "ele-ColdDrink",
                    },
                    "title": "角色管理",
                    "children": []
                },
                {
                    "id": 6,
                    "pid": 0,
                    "path": "/system/dic",
                    "name": "systemDic",
                    "component": "system/dic/index",
                    "meta": {
                        "title": "message.router.systemDic",
                        "isLink":"",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": ["admin"],
                        "icon": "ele-SetUp",
                    },
                    "title": "字典管理",
                    "children": []
                }
            ],
            "title": "系统设置"
        }
    ]
};