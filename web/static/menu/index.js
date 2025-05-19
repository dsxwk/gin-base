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
            "component": "layout/routerView/parent",
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