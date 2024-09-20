const menuModule = {
    list: {
        name: '菜单列表',
        url: '/menu',
        method: 'GET',
        params: {},
        token: {name: 'token', value: true}
    },
    detail: {
        name: '菜单详情',
        url: '/menu/:id',
        method: 'GET',
        params: {},
        token: {name: 'token', value: true}
    },
    create: {
        name: '新增菜单',
        url: '/menu',
        method: 'POST',
        params: {},
        token: {name: 'token', value: true}
    },
    update: {
        name: '更新菜单',
        url: '/menu/:id',
        method: 'PUT',
        params: {},
        token: {name: 'token', value: true}
    },
    delete: {
        name: '删除菜单',
        url: '/menu/:id',
        method: 'DELETE',
        params: {},
        token: {name: 'token', value: true}
    }
};

export default menuModule;