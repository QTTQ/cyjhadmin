import request from '@/utils/request'

// 查询UserInfo列表
export function listUserInfo(query) {
    return request({
        url: '/api/v1/userinfo',
        method: 'get',
        params: query
    })
}

// 查询UserInfo详细
export function getUserInfo (ID) {
    return request({
        url: '/api/v1/userinfo/' + ID,
        method: 'get'
    })
}


// 新增UserInfo
export function addUserInfo(data) {
    return request({
        url: '/api/v1/userinfo',
        method: 'post',
        data: data
    })
}

// 修改UserInfo
export function updateUserInfo(data) {
    return request({
        url: '/api/v1/userinfo/'+data.ID,
        method: 'put',
        data: data
    })
}

// 删除UserInfo
export function delUserInfo(data) {
    return request({
        url: '/api/v1/userinfo',
        method: 'delete',
        data: data
    })
}

