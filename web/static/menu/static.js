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
            "path": "/home",
            "name": "home",
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
                "icon": "iconfont icon-shouye",
                "tagsViewName": "内嵌 iframe2"
            },
            "title": "首页",
            "menuType": "menu",
            "menuSort": 25
        },
        {
            "path": "/system",
            "name": "system",
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
                        "icon": "iconfont icon-caidan",
                        "tagsViewName": "工作流"
                    },
                    "title": "菜单管理"
                },
                {
                    "path": "/system/role",
                    "name": "systemRole",
                    "meta": {
                        "title": "message.router.systemRole",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin"
                        ],
                        "icon": "ele-ColdDrink"
                    },
                    "title": "角色管理"
                },
                {
                    "path": "/system/user",
                    "name": "systemUser",
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
                    "title": "用户管理"
                },
                {
                    "path": "/system/dept",
                    "name": "systemDept",
                    "meta": {
                        "title": "message.router.systemDept",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin"
                        ],
                        "icon": "ele-OfficeBuilding"
                    },
                    "title": "部门管理"
                },
                {
                    "path": "/system/dic",
                    "name": "systemDic",
                    "meta": {
                        "title": "message.router.systemDic",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin"
                        ],
                        "icon": "ele-SetUp"
                    },
                    "title": "字典管理"
                }
            ],
            "title": "系统设置"
        },
        {
            "path": "/limits",
            "name": "limits",
            "redirect": "/limits/frontEnd",
            "meta": {
                "title": "message.router.limits",
                "isLink": "",
                "isHide": false,
                "isKeepAlive": true,
                "isAffix": false,
                "isIframe": false,
                "roles": [
                    "admin",
                    "common"
                ],
                "icon": "iconfont icon-quanxian"
            },
            "children": [
                {
                    "path": "/limits/frontEnd",
                    "name": "limitsFrontEnd",
                    "redirect": "/limits/frontEnd/page",
                    "meta": {
                        "title": "message.router.limitsFrontEnd",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": ""
                    },
                    "children": [
                        {
                            "path": "/limits/frontEnd/page",
                            "name": "limitsFrontEndPage",
                            "meta": {
                                "title": "message.router.limitsFrontEndPage",
                                "isLink": "",
                                "isHide": false,
                                "isKeepAlive": true,
                                "isAffix": false,
                                "isIframe": false,
                                "roles": [
                                    "admin",
                                    "common"
                                ],
                                "icon": ""
                            },
                            "title": "页面权限"
                        },
                        {
                            "path": "/limits/frontEnd/btn",
                            "name": "limitsFrontEndBtn",
                            "meta": {
                                "title": "message.router.limitsFrontEndBtn",
                                "isLink": "",
                                "isHide": false,
                                "isKeepAlive": true,
                                "isAffix": false,
                                "isIframe": false,
                                "roles": [
                                    "admin",
                                    "common"
                                ],
                                "icon": ""
                            },
                            "title": "按钮权限"
                        }
                    ],
                    "title": "前端控制"
                },
                {
                    "path": "/limits/backEnd",
                    "name": "limitsBackEnd",
                    "meta": {
                        "title": "message.router.limitsBackEnd",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": ""
                    },
                    "children": [
                        {
                            "path": "/limits/backEnd/page",
                            "name": "limitsBackEndEndPage",
                            "meta": {
                                "title": "message.router.limitsBackEndEndPage",
                                "isLink": "",
                                "isHide": false,
                                "isKeepAlive": true,
                                "isAffix": false,
                                "isIframe": false,
                                "roles": [
                                    "admin",
                                    "common"
                                ],
                                "icon": ""
                            },
                            "title": "页面权限"
                        }
                    ],
                    "title": "后端控制"
                }
            ],
            "title": "权限管理"
        },
        {
            "path": "/menu",
            "name": "menu",
            "redirect": "/menu/menu1",
            "meta": {
                "title": "message.router.menu",
                "isLink": "",
                "isHide": false,
                "isKeepAlive": true,
                "isAffix": false,
                "isIframe": false,
                "roles": [
                    "admin",
                    "common"
                ],
                "icon": "iconfont icon-caidan"
            },
            "children": [
                {
                    "path": "/menu/menu1",
                    "name": "menu1",
                    "redirect": "/menu/menu1/menu11",
                    "meta": {
                        "title": "message.router.menu1",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-caidan"
                    },
                    "children": [
                        {
                            "path": "/menu/menu1/menu11",
                            "name": "menu11",
                            "meta": {
                                "title": "message.router.menu11",
                                "isLink": "",
                                "isHide": false,
                                "isKeepAlive": true,
                                "isAffix": false,
                                "isIframe": false,
                                "roles": [
                                    "admin",
                                    "common"
                                ],
                                "icon": "iconfont icon-caidan"
                            },
                            "title": "菜单11"
                        },
                        {
                            "path": "/menu/menu1/menu12",
                            "name": "menu12",
                            "redirect": "/menu/menu1/menu12/menu121",
                            "meta": {
                                "title": "message.router.menu12",
                                "isLink": "",
                                "isHide": false,
                                "isKeepAlive": true,
                                "isAffix": false,
                                "isIframe": false,
                                "roles": [
                                    "admin",
                                    "common"
                                ],
                                "icon": "iconfont icon-caidan"
                            },
                            "children": [
                                {
                                    "path": "/menu/menu1/menu12/menu121",
                                    "name": "menu121",
                                    "meta": {
                                        "title": "message.router.menu121",
                                        "isLink": "",
                                        "isHide": false,
                                        "isKeepAlive": true,
                                        "isAffix": false,
                                        "isIframe": false,
                                        "roles": [
                                            "admin",
                                            "common"
                                        ],
                                        "icon": "iconfont icon-caidan"
                                    },
                                    "title": "菜单121"
                                },
                                {
                                    "path": "/menu/menu1/menu12/menu122",
                                    "name": "menu122",
                                    "meta": {
                                        "title": "message.router.menu122",
                                        "isLink": "",
                                        "isHide": false,
                                        "isKeepAlive": true,
                                        "isAffix": false,
                                        "isIframe": false,
                                        "roles": [
                                            "admin",
                                            "common"
                                        ],
                                        "icon": "iconfont icon-caidan"
                                    },
                                    "title": "菜单122"
                                }
                            ],
                            "title": "菜单12"
                        },
                        {
                            "path": "/menu/menu1/menu13",
                            "name": "menu13",
                            "meta": {
                                "title": "message.router.menu13",
                                "isLink": "",
                                "isHide": false,
                                "isKeepAlive": true,
                                "isAffix": false,
                                "isIframe": false,
                                "roles": [
                                    "admin",
                                    "common"
                                ],
                                "icon": "iconfont icon-caidan"
                            },
                            "title": "菜单13"
                        }
                    ],
                    "title": "菜单1"
                },
                {
                    "path": "/menu/menu2",
                    "name": "menu2",
                    "meta": {
                        "title": "message.router.menu2",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-caidan"
                    },
                    "title": "菜单2"
                }
            ],
            "title": "菜单嵌套"
        },
        {
            "path": "/fun",
            "name": "funIndex",
            "redirect": "/fun/tagsView",
            "meta": {
                "title": "message.router.funIndex",
                "isLink": "",
                "isHide": false,
                "isKeepAlive": true,
                "isAffix": false,
                "isIframe": false,
                "roles": [
                    "admin",
                    "common"
                ],
                "icon": "iconfont icon-crew_feature"
            },
            "children": [
                {
                    "path": "/fun/tagsView",
                    "name": "funTagsView",
                    "meta": {
                        "title": "message.router.funTagsView",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "ele-Pointer"
                    },
                    "title": "tagsView 操作"
                },
                {
                    "path": "/fun/countup",
                    "name": "funCountup",
                    "meta": {
                        "title": "message.router.funCountup",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "ele-Odometer"
                    },
                    "title": "数字滚动"
                },
                {
                    "path": "/fun/wangEditor",
                    "name": "funWangEditor",
                    "meta": {
                        "title": "message.router.funWangEditor",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-fuwenbenkuang"
                    },
                    "title": "Editor 编辑器"
                },
                {
                    "path": "/fun/cropper",
                    "name": "funCropper",
                    "meta": {
                        "title": "message.router.funCropper",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-caijian"
                    },
                    "title": "图片裁剪"
                },
                {
                    "path": "/fun/qrcode",
                    "name": "funQrcode",
                    "meta": {
                        "title": "message.router.funQrcode",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-ico"
                    },
                    "title": "二维码生成"
                },
                {
                    "path": "/fun/echartsMap",
                    "name": "funEchartsMap",
                    "meta": {
                        "title": "message.router.funEchartsMap",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-ditu"
                    },
                    "title": "地理坐标/地图"
                },
                {
                    "path": "/fun/printJs",
                    "name": "funPrintJs",
                    "meta": {
                        "title": "message.router.funPrintJs",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "ele-Printer"
                    },
                    "title": "页面打印"
                },
                {
                    "path": "/fun/clipboard",
                    "name": "funClipboard",
                    "meta": {
                        "title": "message.router.funClipboard",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "ele-DocumentCopy"
                    },
                    "title": "复制剪切"
                },
                {
                    "path": "/fun/gridLayout",
                    "name": "funGridLayout",
                    "meta": {
                        "title": "message.router.funGridLayout",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-tuodong"
                    },
                    "title": "拖拽布局"
                },
                {
                    "path": "/fun/splitpanes",
                    "name": "funSplitpanes",
                    "meta": {
                        "title": "message.router.funSplitpanes",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon--chaifenlie"
                    },
                    "title": "窗格拆分器"
                }
            ],
            "title": "功能"
        },
        {
            "path": "/pages",
            "name": "pagesIndex",
            "redirect": "/pages/filtering",
            "meta": {
                "title": "message.router.pagesIndex",
                "isLink": "",
                "isHide": false,
                "isKeepAlive": true,
                "isAffix": false,
                "isIframe": false,
                "roles": [
                    "admin",
                    "common"
                ],
                "icon": "iconfont icon-fuzhiyemian"
            },
            "children": [
                {
                    "path": "/pages/filtering",
                    "name": "pagesFiltering",
                    "meta": {
                        "title": "message.router.pagesFiltering",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "ele-Sell",
                        "tagsViewName": "内嵌 iframe2"
                    },
                    "children": [
                        {
                            "path": "/pages/filtering/details",
                            "name": "pagesFilteringDetails",
                            "meta": {
                                "title": "message.router.pagesFilteringDetails",
                                "isLink": "",
                                "isHide": true,
                                "isKeepAlive": false,
                                "isAffix": false,
                                "isIframe": false,
                                "roles": [
                                    "admin",
                                    "common"
                                ],
                                "icon": "ele-Sunny"
                            },
                            "title": "过滤筛选组件详情"
                        }
                    ],
                    "title": "过滤筛选组件"
                },
                {
                    "path": "/pages/filtering/details1",
                    "name": "pagesFilteringDetails1",
                    "meta": {
                        "title": "message.router.pagesFilteringDetails1",
                        "isLink": "",
                        "isHide": true,
                        "isKeepAlive": false,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "ele-Sunny"
                    },
                    "title": "过滤筛选组件详情111"
                },
                {
                    "path": "/pages/iocnfont",
                    "name": "pagesIocnfont",
                    "meta": {
                        "title": "message.router.pagesIocnfont",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "ele-Present"
                    },
                    "title": "ali 字体图标"
                },
                {
                    "path": "/pages/element",
                    "name": "pagesElement",
                    "meta": {
                        "title": "message.router.pagesElement",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "ele-Eleme"
                    },
                    "title": "ele 字体图标"
                },
                {
                    "path": "/pages/awesome",
                    "name": "pagesAwesome",
                    "meta": {
                        "title": "message.router.pagesAwesome",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "ele-SetUp"
                    },
                    "title": "awe 字体图标"
                },
                {
                    "path": "/pages/formAdapt",
                    "name": "pagesFormAdapt",
                    "meta": {
                        "title": "message.router.pagesFormAdapt",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-biaodan",
                        "tagsViewName": "过滤筛选组件"
                    },
                    "title": "表单自适应"
                },
                {
                    "path": "/pages/tableRules",
                    "name": "pagesTableRules",
                    "meta": {
                        "title": "message.router.pagesTableRules",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-jiliandongxuanzeqi",
                        "tagsViewName": "表单自适应"
                    },
                    "title": "表单表格验证"
                },
                {
                    "path": "/pages/formI18n",
                    "name": "pagesFormI18n",
                    "meta": {
                        "title": "message.router.pagesFormI18n",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-diqiu",
                        "tagsViewName": "表单表格验证"
                    },
                    "title": "表单国际化"
                },
                {
                    "path": "/pages/formRules",
                    "name": "pagesFormRules",
                    "meta": {
                        "title": "message.router.pagesFormRules",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-shuxing",
                        "tagsViewName": "表单国际化"
                    },
                    "title": "多表单验证"
                },
                {
                    "path": "/pages/listAdapt",
                    "name": "pagesListAdapt",
                    "meta": {
                        "title": "message.router.pagesListAdapt",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-chazhaobiaodanliebiao",
                        "tagsViewName": "多表单验证"
                    },
                    "title": "列表自适应"
                },
                {
                    "path": "/pages/waterfall",
                    "name": "pagesWaterfall",
                    "meta": {
                        "title": "message.router.pagesWaterfall",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-zidingyibuju",
                        "tagsViewName": "列表自适应"
                    },
                    "title": "瀑布屏"
                },
                {
                    "path": "/pages/steps",
                    "name": "pagesSteps",
                    "meta": {
                        "title": "message.router.pagesSteps",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-step",
                        "tagsViewName": "瀑布屏"
                    },
                    "title": "步骤条"
                },
                {
                    "path": "/pages/preview",
                    "name": "pagesPreview",
                    "meta": {
                        "title": "message.router.pagesPreview",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-15tupianyulan",
                        "tagsViewName": "步骤条"
                    },
                    "title": "大图预览"
                },
                {
                    "path": "/pages/waves",
                    "name": "pagesWaves",
                    "meta": {
                        "title": "message.router.pagesWaves",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-bolangneng",
                        "tagsViewName": "大图预览"
                    },
                    "title": "波浪效果"
                },
                {
                    "path": "/pages/tree",
                    "name": "pagesTree",
                    "meta": {
                        "title": "message.router.pagesTree",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-shuxingtu",
                        "tagsViewName": "菜单管理"
                    },
                    "title": "树形改表格"
                },
                {
                    "path": "/pages/drag",
                    "name": "pagesDrag",
                    "meta": {
                        "title": "message.router.pagesDrag",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "ele-Pointer",
                        "tagsViewName": "树形改表格"
                    },
                    "title": "拖动指令"
                },
                {
                    "path": "/pages/lazyImg",
                    "name": "pagesLazyImg",
                    "meta": {
                        "title": "message.router.pagesLazyImg",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin"
                        ],
                        "icon": "ele-PictureFilled"
                    },
                    "title": "图片懒加载"
                },
                {
                    "path": "/pages/dynamicForm",
                    "name": "pagesDynamicForm",
                    "meta": {
                        "title": "message.router.pagesDynamicForm",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin"
                        ],
                        "icon": "iconfont icon-wenducanshu-05",
                        "tagsViewName": "树形改表格"
                    },
                    "title": "动态复杂表单"
                },
                {
                    "path": "/pages/workflow",
                    "name": "pagesWorkflow",
                    "meta": {
                        "title": "message.router.pagesWorkflow",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin"
                        ],
                        "icon": "ele-Connection",
                        "tagsViewName": "动态复杂表单"
                    },
                    "title": "工作流"
                }
            ],
            "title": "页面"
        },
        {
            "path": "/make",
            "name": "makeIndex",
            "redirect": "/make/selector",
            "meta": {
                "title": "message.router.makeIndex",
                "isLink": "",
                "isHide": false,
                "isKeepAlive": true,
                "isAffix": false,
                "isIframe": false,
                "roles": [
                    "admin"
                ],
                "icon": "iconfont icon-siweidaotu"
            },
            "children": [
                {
                    "path": "/make/selector",
                    "name": "makeSelector",
                    "meta": {
                        "title": "message.router.makeSelector",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-xuanzeqi"
                    },
                    "title": "图标选择器"
                },
                {
                    "path": "/make/noticeBar",
                    "name": "makeNoticeBar",
                    "meta": {
                        "title": "message.router.makeNoticeBar",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "ele-Bell"
                    },
                    "title": "滚动通知栏"
                },
                {
                    "path": "/make/svgDemo",
                    "name": "makeSvgDemo",
                    "meta": {
                        "title": "message.router.makeSvgDemo",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "fa fa-thumbs-o-up"
                    },
                    "title": "svgIcon 演示"
                },
                {
                    "path": "/make/tableDemo",
                    "name": "makeTableDemo",
                    "meta": {
                        "title": "message.router.makeTableDemo",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "icon": "iconfont icon-shuju"
                    },
                    "title": "表格封装演示"
                }
            ],
            "title": "组件封装"
        },
        {
            "path": "/params",
            "name": "paramsIndex",
            "redirect": "/params/common",
            "meta": {
                "title": "message.router.paramsIndex",
                "isLink": "",
                "isHide": false,
                "isKeepAlive": true,
                "isAffix": false,
                "isIframe": false,
                "roles": [
                    "admin"
                ],
                "icon": "iconfont icon-zhongduancanshu"
            },
            "children": [
                {
                    "path": "/params/common",
                    "name": "paramsCommon",
                    "meta": {
                        "title": "message.router.paramsCommon",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin"
                        ],
                        "icon": "iconfont icon-putong"
                    },
                    "title": "普通路由"
                },
                {
                    "path": "/params/common/details",
                    "name": "paramsCommonDetails",
                    "meta": {
                        "title": "message.router.paramsCommonDetails",
                        "isLink": "",
                        "isHide": true,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin"
                        ],
                        "icon": "ele-Comment"
                    },
                    "title": "普通路由详情"
                },
                {
                    "path": "/params/dynamic",
                    "name": "paramsDynamic",
                    "meta": {
                        "title": "message.router.paramsDynamic",
                        "isLink": "",
                        "isHide": false,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin"
                        ],
                        "icon": "iconfont icon-dongtai"
                    },
                    "title": "动态路由"
                },
                {
                    "path": "/params/dynamic/details/:t/:id/:tagsViewName",
                    "name": "paramsDynamicDetails",
                    "meta": {
                        "title": "message.router.paramsDynamicDetails",
                        "isLink": "",
                        "isHide": true,
                        "isKeepAlive": true,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin"
                        ],
                        "icon": "ele-Lightning",
                        "isDynamic": true,
                        "isDynamicPath": "/params/dynamic/details/:t/:id/:tagsViewName"
                    },
                    "title": "动态路由详情"
                }
            ],
            "title": "路由参数"
        },
        {
            "path": "/visualizing",
            "name": "visualizingIndex",
            "redirect": "/visualizing/visualizingLinkDemo1",
            "meta": {
                "title": "message.router.visualizingIndex",
                "isLink": "",
                "isHide": false,
                "isKeepAlive": true,
                "isAffix": false,
                "isIframe": false,
                "roles": [
                    "admin"
                ],
                "icon": "ele-ChatLineRound"
            },
            "children": [
                {
                    "path": "/visualizing/visualizingLinkDemo1",
                    "name": "visualizingLinkDemo1",
                    "meta": {
                        "title": "message.router.visualizingLinkDemo1",
                        "isLink": "/visualizingDemo1",
                        "isHide": false,
                        "isKeepAlive": false,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin"
                        ],
                        "icon": "iconfont icon-caozuo-wailian"
                    },
                    "title": "数据可视化演示1"
                },
                {
                    "path": "/visualizing/visualizingLinkDemo2",
                    "name": "visualizingLinkDemo2",
                    "meta": {
                        "title": "message.router.visualizingLinkDemo2",
                        "isLink": "/visualizingDemo2",
                        "isHide": false,
                        "isKeepAlive": false,
                        "isAffix": false,
                        "isIframe": false,
                        "roles": [
                            "admin"
                        ],
                        "icon": "iconfont icon-caozuo-wailian"
                    },
                    "title": "数据可视化演示2"
                }
            ],
            "title": "数据可视化"
        },
        {
            "path": "/chart",
            "name": "chartIndex",
            "meta": {
                "title": "message.router.chartIndex",
                "isLink": "",
                "isHide": false,
                "isKeepAlive": true,
                "isAffix": false,
                "isIframe": false,
                "roles": [
                    "admin",
                    "common"
                ],
                "icon": "iconfont icon-ico_shuju"
            },
            "title": "大数据图表"
        },
        {
            "path": "/personal",
            "name": "personal",
            "meta": {
                "title": "message.router.personal",
                "isLink": "",
                "isHide": false,
                "isKeepAlive": true,
                "isAffix": false,
                "isIframe": false,
                "roles": [
                    "admin",
                    "common"
                ],
                "icon": "iconfont icon-gerenzhongxin"
            },
            "title": "个人中心"
        },
        {
            "path": "/tools",
            "name": "tools",
            "meta": {
                "title": "message.router.tools",
                "isLink": "",
                "isHide": false,
                "isKeepAlive": true,
                "isAffix": false,
                "isIframe": false,
                "roles": [
                    "admin",
                    "common"
                ],
                "icon": "iconfont icon-gongju"
            },
            "title": "工具类集合"
        },
        {
            "path": "/link",
            "name": "layoutLinkView",
            "meta": {
                "title": "message.router.layoutLinkView",
                "isLink": "https://element-plus.gitee.io/#/zh-CN/component/installation",
                "isHide": false,
                "isKeepAlive": false,
                "isAffix": false,
                "isIframe": false,
                "roles": [
                    "admin"
                ],
                "icon": "iconfont icon-caozuo-wailian"
            },
            "title": "外链"
        },
        {
            "path": "/iframesOne",
            "name": "layoutIframeViewOne",
            "meta": {
                "title": "message.router.layoutIframeViewOne",
                "isLink": "https://nodejs.org/zh-cn/",
                "isHide": false,
                "isKeepAlive": true,
                "isAffix": true,
                "isIframe": true,
                "roles": [
                    "admin"
                ],
                "icon": "iconfont icon-neiqianshujuchucun",
                "isIframeOpen": true,
                "loading": false,
                "tagsViewName": "内嵌 iframe2"
            },
            "title": "内嵌 iframe1"
        },
        {
            "path": "/iframesTwo",
            "name": "layoutIframeViewTwo",
            "meta": {
                "title": "message.router.layoutIframeViewTwo",
                "isLink": "https://undraw.co/illustrations",
                "isHide": false,
                "isKeepAlive": true,
                "isAffix": true,
                "isIframe": true,
                "roles": [
                    "admin"
                ],
                "icon": "iconfont icon-neiqianshujuchucun",
                "isIframeOpen": true,
                "loading": false,
                "tagsViewName": "首页"
            },
            "title": "内嵌 iframe2"
        }
    ]
};