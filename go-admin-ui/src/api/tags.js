import request from '@/utils/request'

// 查询Tags列表
export function listTags(query) {
    return request({
        url: '/api/v1/tags',
        method: 'get',
        params: query
    })
}

// 查询Tags详细
export function getTags (ID) {
    return request({
        url: '/api/v1/tags/' + ID,
        method: 'get'
    })
}


// 新增Tags
export function addTags(data) {
    return request({
        url: '/api/v1/tags',
        method: 'post',
        data: data
    })
}

// 修改Tags
export function updateTags(data) {
    return request({
        url: '/api/v1/tags/'+data.ID,
        method: 'put',
        data: data
    })
}

// 删除Tags
export function delTags(data) {
    return request({
        url: '/api/v1/tags',
        method: 'delete',
        data: data
    })
}

