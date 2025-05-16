export const menuJson = {
    "code": 0,
    "type": "adminMenu",
    "data": [
        {
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
            }
        },
        {
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
                "icon": "iconfont icon-xitongshezhi"
            },
            "children": [
                {
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
                    }
                },
                {
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
                    }
                },
                {
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
                },
                {
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
                },
            ]
        },
    ]
};