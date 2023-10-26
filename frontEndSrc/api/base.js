import instance from './index';

export function login(params) {
    return instance.request({
        method: 'get',
        url: '/v1/login',
        params:params
    })
};

export function getProjects(params) {
    return instance.request({
        method: 'get',
        url: '/v1/projects',
        params:params
    })
};

export function getColumns(params) {
    return instance.request({
        method: 'get',
        url: '/v1/data/columns',
        params:params
    })
};

export function getData(params) {
    return instance.request({
        method: 'get',
        url: '/v1/data',
        params:params
    })
};

export default {
    login,
    getProjects,
    getColumns,
    getData
};
