import {defineStore} from 'pinia';

/**
 * 路由列表
 * @methods setRoutesList 设置路由数据
 * @methods setColumnsMenuHover 设置分栏布局菜单鼠标移入 boolean
 * @methods setColumnsNavHover 设置分栏布局最左侧导航鼠标移入 boolean
 */
export const useRoutesList = defineStore('routesList', {
    state: () => ({
        routesList: [],
        isColumnsMenuHover: false,
        isColumnsNavHover: false,
    }),
    actions: {
        async setRoutesList(data) {
            this.routesList = data;
        },
        async setColumnsMenuHover(bool) {
            this.isColumnsMenuHover = bool;
        },
        async setColumnsNavHover(bool) {
            this.isColumnsNavHover = bool;
        },
    },
});
