const articleModule = {
    list: {
        name: '文章列表',
        url: '/article',
        method: 'GET',
        params: {},
        token: {name: 'token', value: true}
    },
    detail: {
        name: '文章详情',
        url: '/article/:id',
        method: 'GET',
        params: {},
        token: {name: 'token', value: true}
    },
    create: {
        name: '新增文章',
        url: '/article',
        method: 'POST',
        params: {},
        token: {name: 'token', value: true}
    },
    update: {
        name: '更新文章',
        url: '/article/:id',
        method: 'PUT',
        params: {},
        token: {name: 'token', value: true}
    },
    delete: {
        name: '删除文章',
        url: '/article/:id',
        method: 'DELETE',
        params: {},
        token: {name: 'token', value: true}
    }
};

export default articleModule;