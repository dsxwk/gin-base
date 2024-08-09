/**
 * @description 📚 路由参数配置简介
 * @param path ==> 路由菜单访问路径
 * @param name ==> 路由 name (对应页面组件 name, 可用作 KeepAlive 缓存标识 && 按钮权限筛选)
 * @param redirect ==> 路由重定向地址
 * @param component ==> 视图文件路径
 * @param meta ==> 路由菜单元信息
 * @param meta.icon ==> 菜单和面包屑对应的图标
 * @param meta.title ==> 路由标题 (用作 document.title || 菜单的名称)
 * @param meta.activeMenu ==> 当前路由为详情页时，需要高亮的菜单
 * @param meta.isLink ==> 路由外链时填写的访问地址
 * @param meta.isHide ==> 是否在菜单中隐藏 (通常列表详情页需要隐藏)
 * @param meta.isFull ==> 菜单是否全屏 (示例：数据大屏页面)
 * @param meta.isAffix ==> 菜单是否固定在标签页中 (首页通常是固定项)
 * @param meta.isKeepAlive ==> 当前路由是否缓存
 * */
export const menuJson = [
  {
    "path": "/dashboard",
    "name": "home",
    "component": "@/views/home/index.vue",
    "meta": {
      "icon": "HomeFilled",
      "title": "首页",
      "isLink": "",
      "isHide": false,
      "isFull": false,
      "isAffix": true,
      "isKeepAlive": true
    }
  },
  {
    "path": "/",
    "name": "user",
    "redirect": "/user",
    "meta": {
      "icon": "User",
      "title": "用户管理",
      "isLink": "",
      "isHide": false,
      "isFull": false,
      "isAffix": false,
      "isKeepAlive": true
    },
    "children": [
      {
        "path": "/user",
        "name": "userList",
        "component": "@/views/user/index.vue",
        "meta": {
          "icon": "",
          "title": "用户列表",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        },
        "children": []
      }
    ]
  },
  {
    "path": "/dataScreen",
    "name": "dataScreen",
    "component": "/dataScreen/index",
    "meta": {
      "icon": "Histogram",
      "title": "数据大屏",
      "isLink": "",
      "isHide": false,
      "isFull": true,
      "isAffix": false,
      "isKeepAlive": true
    }
  },
  {
    "path": "/proTable",
    "name": "proTable",
    "redirect": "/proTable/useProTable",
    "meta": {
      "icon": "MessageBox",
      "title": "超级表格",
      "isLink": "",
      "isHide": false,
      "isFull": false,
      "isAffix": false,
      "isKeepAlive": true
    },
    "children": [
      {
        "path": "/proTable/useProTable",
        "name": "useProTable",
        "component": "/proTable/useProTable/index",
        "meta": {
          "icon": "Menu",
          "title": "使用 ProTable",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        },
        "children": []
      },
      {
        "path": "/proTable/useTreeFilter",
        "name": "useTreeFilter",
        "component": "/proTable/useTreeFilter/index",
        "meta": {
          "icon": "Menu",
          "title": "使用 TreeFilter",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/proTable/useSelectFilter",
        "name": "useSelectFilter",
        "component": "/proTable/useSelectFilter/index",
        "meta": {
          "icon": "Menu",
          "title": "使用 SelectFilter",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/proTable/treeProTable",
        "name": "treeProTable",
        "component": "/proTable/treeProTable/index",
        "meta": {
          "icon": "Menu",
          "title": "树形 ProTable",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/proTable/complexProTable",
        "name": "complexProTable",
        "component": "/proTable/complexProTable/index",
        "meta": {
          "icon": "Menu",
          "title": "复杂 ProTable",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/proTable/document",
        "name": "proTableDocument",
        "component": "/proTable/document/index",
        "meta": {
          "icon": "Menu",
          "title": "ProTable 文档",
          "isLink": "https://juejin.cn/post/7166068828202336263/#heading-14",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      }
    ]
  },
  {
    "path": "/auth",
    "name": "auth",
    "redirect": "/auth/menu",
    "meta": {
      "icon": "Lock",
      "title": "权限管理",
      "isLink": "",
      "isHide": false,
      "isFull": false,
      "isAffix": false,
      "isKeepAlive": true
    },
    "children": [
      {
        "path": "/auth/menu",
        "name": "authMenu",
        "component": "/auth/menu/index",
        "meta": {
          "icon": "Menu",
          "title": "菜单权限",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/auth/button",
        "name": "authButton",
        "component": "/auth/button/index",
        "meta": {
          "icon": "Menu",
          "title": "按钮权限",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      }
    ]
  },
  {
    "path": "/assembly",
    "name": "assembly",
    "redirect": "/assembly/guide",
    "meta": {
      "icon": "Briefcase",
      "title": "常用组件",
      "isLink": "",
      "isHide": false,
      "isFull": false,
      "isAffix": false,
      "isKeepAlive": true
    },
    "children": [
      {
        "path": "/assembly/guide",
        "name": "guide",
        "component": "/assembly/guide/index",
        "meta": {
          "icon": "Menu",
          "title": "引导页",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/assembly/tabs",
        "name": "tabs",
        "component": "/assembly/tabs/index",
        "meta": {
          "icon": "Menu",
          "title": "标签页操作",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        },
        "children": []
      },
      {
        "path": "/assembly/selectIcon",
        "name": "selectIcon",
        "component": "/assembly/selectIcon/index",
        "meta": {
          "icon": "Menu",
          "title": "图标选择器",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/assembly/selectFilter",
        "name": "selectFilter",
        "component": "/assembly/selectFilter/index",
        "meta": {
          "icon": "Menu",
          "title": "分类筛选器",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/assembly/treeFilter",
        "name": "treeFilter",
        "component": "/assembly/treeFilter/index",
        "meta": {
          "icon": "Menu",
          "title": "树形筛选器",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/assembly/svgIcon",
        "name": "svgIcon",
        "component": "/assembly/svgIcon/index",
        "meta": {
          "icon": "Menu",
          "title": "SVG 图标",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/assembly/uploadFile",
        "name": "uploadFile",
        "component": "/assembly/uploadFile/index",
        "meta": {
          "icon": "Menu",
          "title": "文件上传",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/assembly/batchImport",
        "name": "batchImport",
        "component": "/assembly/batchImport/index",
        "meta": {
          "icon": "Menu",
          "title": "批量添加数据",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/assembly/wangEditor",
        "name": "wangEditor",
        "component": "/assembly/wangEditor/index",
        "meta": {
          "icon": "Menu",
          "title": "富文本编辑器",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/assembly/draggable",
        "name": "draggable",
        "component": "/assembly/draggable/index",
        "meta": {
          "icon": "Menu",
          "title": "拖拽组件",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      }
    ]
  },
  {
    "path": "/dashboard",
    "name": "dashboard",
    "redirect": "/dashboard/dataVisualize",
    "meta": {
      "icon": "Odometer",
      "title": "Dashboard",
      "isLink": "",
      "isHide": false,
      "isFull": false,
      "isAffix": false,
      "isKeepAlive": true
    },
    "children": [
      {
        "path": "/dashboard/dataVisualize",
        "name": "dataVisualize",
        "component": "/dashboard/dataVisualize/index",
        "meta": {
          "icon": "Menu",
          "title": "数据可视化",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      }
    ]
  },
  {
    "path": "/form",
    "name": "form",
    "redirect": "/form/proForm",
    "meta": {
      "icon": "Tickets",
      "title": "表单 Form",
      "isLink": "",
      "isHide": false,
      "isFull": false,
      "isAffix": false,
      "isKeepAlive": true
    },
    "children": [
      {
        "path": "/form/proForm",
        "name": "proForm",
        "component": "/form/proForm/index",
        "meta": {
          "icon": "Menu",
          "title": "超级 Form",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/form/basicForm",
        "name": "basicForm",
        "component": "/form/basicForm/index",
        "meta": {
          "icon": "Menu",
          "title": "基础 Form",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/form/validateForm",
        "name": "validateForm",
        "component": "/form/validateForm/index",
        "meta": {
          "icon": "Menu",
          "title": "校验 Form",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/form/dynamicForm",
        "name": "dynamicForm",
        "component": "/form/dynamicForm/index",
        "meta": {
          "icon": "Menu",
          "title": "动态 Form",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      }
    ]
  },
  {
    "path": "/echarts",
    "name": "echarts",
    "redirect": "/echarts/waterChart",
    "meta": {
      "icon": "TrendCharts",
      "title": "ECharts",
      "isLink": "",
      "isHide": false,
      "isFull": false,
      "isAffix": false,
      "isKeepAlive": true
    },
    "children": [
      {
        "path": "/echarts/waterChart",
        "name": "waterChart",
        "component": "/echarts/waterChart/index",
        "meta": {
          "icon": "Menu",
          "title": "水型图",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/echarts/columnChart",
        "name": "columnChart",
        "component": "/echarts/columnChart/index",
        "meta": {
          "icon": "Menu",
          "title": "柱状图",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/echarts/lineChart",
        "name": "lineChart",
        "component": "/echarts/lineChart/index",
        "meta": {
          "icon": "Menu",
          "title": "折线图",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/echarts/pieChart",
        "name": "pieChart",
        "component": "/echarts/pieChart/index",
        "meta": {
          "icon": "Menu",
          "title": "饼图",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/echarts/radarChart",
        "name": "radarChart",
        "component": "/echarts/radarChart/index",
        "meta": {
          "icon": "Menu",
          "title": "雷达图",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/echarts/nestedChart",
        "name": "nestedChart",
        "component": "/echarts/nestedChart/index",
        "meta": {
          "icon": "Menu",
          "title": "嵌套环形图",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      }
    ]
  },
  {
    "path": "/directives",
    "name": "directives",
    "redirect": "/directives/copyDirect",
    "meta": {
      "icon": "Stamp",
      "title": "自定义指令",
      "isLink": "",
      "isHide": false,
      "isFull": false,
      "isAffix": false,
      "isKeepAlive": true
    },
    "children": [
      {
        "path": "/directives/copyDirect",
        "name": "copyDirect",
        "component": "/directives/copyDirect/index",
        "meta": {
          "icon": "Menu",
          "title": "复制指令",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/directives/watermarkDirect",
        "name": "watermarkDirect",
        "component": "/directives/watermarkDirect/index",
        "meta": {
          "icon": "Menu",
          "title": "水印指令",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/directives/dragDirect",
        "name": "dragDirect",
        "component": "/directives/dragDirect/index",
        "meta": {
          "icon": "Menu",
          "title": "拖拽指令",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/directives/debounceDirect",
        "name": "debounceDirect",
        "component": "/directives/debounceDirect/index",
        "meta": {
          "icon": "Menu",
          "title": "防抖指令",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/directives/throttleDirect",
        "name": "throttleDirect",
        "component": "/directives/throttleDirect/index",
        "meta": {
          "icon": "Menu",
          "title": "节流指令",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/directives/longpressDirect",
        "name": "longpressDirect",
        "component": "/directives/longpressDirect/index",
        "meta": {
          "icon": "Menu",
          "title": "长按指令",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      }
    ]
  },
  {
    "path": "/menu",
    "name": "menu",
    "redirect": "/menu/menu1",
    "meta": {
      "icon": "List",
      "title": "菜单嵌套",
      "isLink": "",
      "isHide": false,
      "isFull": false,
      "isAffix": false,
      "isKeepAlive": true
    },
    "children": [
      {
        "path": "/menu/menu1",
        "name": "menu1",
        "component": "/menu/menu1/index",
        "meta": {
          "icon": "Menu",
          "title": "菜单1",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/menu/menu2",
        "name": "menu2",
        "redirect": "/menu/menu2/menu21",
        "meta": {
          "icon": "Menu",
          "title": "菜单2",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        },
        "children": [
          {
            "path": "/menu/menu2/menu21",
            "name": "menu21",
            "component": "/menu/menu2/menu21/index",
            "meta": {
              "icon": "Menu",
              "title": "菜单2-1",
              "isLink": "",
              "isHide": false,
              "isFull": false,
              "isAffix": false,
              "isKeepAlive": true
            }
          },
          {
            "path": "/menu/menu2/menu22",
            "name": "menu22",
            "redirect": "/menu/menu2/menu22/menu221",
            "meta": {
              "icon": "Menu",
              "title": "菜单2-2",
              "isLink": "",
              "isHide": false,
              "isFull": false,
              "isAffix": false,
              "isKeepAlive": true
            },
            "children": [
              {
                "path": "/menu/menu2/menu22/menu221",
                "name": "menu221",
                "component": "/menu/menu2/menu22/menu221/index",
                "meta": {
                  "icon": "Menu",
                  "title": "菜单2-2-1",
                  "isLink": "",
                  "isHide": false,
                  "isFull": false,
                  "isAffix": false,
                  "isKeepAlive": true
                }
              },
              {
                "path": "/menu/menu2/menu22/menu222",
                "name": "menu222",
                "component": "/menu/menu2/menu22/menu222/index",
                "meta": {
                  "icon": "Menu",
                  "title": "菜单2-2-2",
                  "isLink": "",
                  "isHide": false,
                  "isFull": false,
                  "isAffix": false,
                  "isKeepAlive": true
                }
              }
            ]
          },
          {
            "path": "/menu/menu2/menu23",
            "name": "menu23",
            "component": "/menu/menu2/menu23/index",
            "meta": {
              "icon": "Menu",
              "title": "菜单2-3",
              "isLink": "",
              "isHide": false,
              "isFull": false,
              "isAffix": false,
              "isKeepAlive": true
            }
          }
        ]
      },
      {
        "path": "/menu/menu3",
        "name": "menu3",
        "component": "/menu/menu3/index",
        "meta": {
          "icon": "Menu",
          "title": "菜单3",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      }
    ]
  },
  {
    "path": "/system",
    "name": "system",
    "redirect": "/system/accountManage",
    "meta": {
      "icon": "Tools",
      "title": "系统管理",
      "isLink": "",
      "isHide": false,
      "isFull": false,
      "isAffix": false,
      "isKeepAlive": true
    },
    "children": [
      {
        "path": "/system/accountManage",
        "name": "accountManage",
        "component": "/system/accountManage/index",
        "meta": {
          "icon": "Menu",
          "title": "账号管理",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/system/roleManage",
        "name": "roleManage",
        "component": "/system/roleManage/index",
        "meta": {
          "icon": "Menu",
          "title": "角色管理",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/system/menuMange",
        "name": "menuMange",
        "component": "/system/menuMange/index",
        "meta": {
          "icon": "Menu",
          "title": "菜单管理",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/system/departmentManage",
        "name": "departmentManage",
        "component": "/system/departmentManage/index",
        "meta": {
          "icon": "Menu",
          "title": "部门管理",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/system/dictManage",
        "name": "dictManage",
        "component": "/system/dictManage/index",
        "meta": {
          "icon": "Menu",
          "title": "字典管理",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/system/timingTask",
        "name": "timingTask",
        "component": "/system/timingTask/index",
        "meta": {
          "icon": "Menu",
          "title": "定时任务",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/system/systemLog",
        "name": "systemLog",
        "component": "/system/systemLog/index",
        "meta": {
          "icon": "Menu",
          "title": "系统日志",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      }
    ]
  },
  {
    "path": "/link",
    "name": "link",
    "redirect": "/link/bing",
    "meta": {
      "icon": "Paperclip",
      "title": "外部链接",
      "isLink": "",
      "isHide": false,
      "isFull": false,
      "isAffix": false,
      "isKeepAlive": true
    },
    "children": [
      {
        "path": "/link/bing",
        "name": "bing",
        "component": "/link/bing/index",
        "meta": {
          "icon": "Menu",
          "title": "Bing 内嵌",
          "isLink": "",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/link/gitee",
        "name": "gitee",
        "component": "/link/gitee/index",
        "meta": {
          "icon": "Menu",
          "title": "Gitee 仓库",
          "isLink": "https://gitee.com/HalseySpicy/Geeker-Admin",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/link/github",
        "name": "github",
        "component": "/link/github/index",
        "meta": {
          "icon": "Menu",
          "title": "GitHub 仓库",
          "isLink": "https://github.com/HalseySpicy/Geeker-Admin",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/link/docs",
        "name": "docs",
        "component": "/link/docs/index",
        "meta": {
          "icon": "Menu",
          "title": "项目文档",
          "isLink": "https://docs.spicyboy.cn",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      },
      {
        "path": "/link/juejin",
        "name": "juejin",
        "component": "/link/juejin/index",
        "meta": {
          "icon": "Menu",
          "title": "掘金主页",
          "isLink": "https://juejin.cn/user/3263814531551816/posts",
          "isHide": false,
          "isFull": false,
          "isAffix": false,
          "isKeepAlive": true
        }
      }
    ]
  },
  {
    "path": "/about/index",
    "name": "about",
    "component": "/about/index",
    "meta": {
      "icon": "InfoFilled",
      "title": "关于项目",
      "isLink": "",
      "isHide": false,
      "isFull": false,
      "isAffix": false,
      "isKeepAlive": true
    }
  }
];