const userModule = {
    list: {
        name: '用户列表',
        url: '/user',
        method: 'GET',
        params: {},
        token: {name: 'token', value: true}
    },
    create: {
        name: '新增用户',
        url: '/user',
        method: 'POST',
        params: {},
        token: {name: 'token', value: true}
    },
    update: {
        name: '更新用户',
        url: '/user/:id',
        method: 'PUT',
        params: {},
        token: {name: 'token', value: true}
    },
    delete: {
        name: '删除用户',
        url: '/user/:id',
        method: 'DELETE',
        params: {},
        token: {name: 'token', value: true}
    }
};

export default userModule;