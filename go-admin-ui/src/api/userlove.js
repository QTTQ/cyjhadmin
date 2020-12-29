import request from '@/utils/request'

// 查询UserLove列表
export function listUserLove(query) {
    return request({
        url: '/api/v1/userlove',
        method: 'get',
        params: query
    })
}

// 查询UserLove详细
export function getUserLove (ID) {
    return request({
        url: '/api/v1/userlove/' + ID,
        method: 'get'
    })
}


// 新增UserLove
export function addUserLove(data) {
    return request({
        url: '/api/v1/userlove',
        method: 'post',
        data: data
    })
}

// 修改UserLove
export function updateUserLove(data) {
    return request({
        url: '/api/v1/userlove/'+data.ID,
        method: 'put',
        data: data
    })
}

// 删除UserLove
export function delUserLove(data) {
    return request({
        url: '/api/v1/userlove',
        method: 'delete',
        data: data
    })
}

