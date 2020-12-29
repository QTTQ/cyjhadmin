import request from '@/utils/request'

// 查询Dress列表
export function listDress(query) {
    return request({
        url: '/api/v1/dress',
        method: 'get',
        params: query
    })
}

// 查询Dress详细
export function getDress (ID) {
    return request({
        url: '/api/v1/dress/' + ID,
        method: 'get'
    })
}


// 新增Dress
export function addDress(data) {
    return request({
        url: '/api/v1/dress',
        method: 'post',
        data: data,
        headers: {'Content-Type': "application/x-www-form-urlencoded"}//设置header信息
    })
}

// 修改Dress
export function updateDress(data) {
    return request({
        url: '/api/v1/dress/'+data.ID,
        method: 'put',
        data: data
    })
}

// 删除Dress
export function delDress(data) {
    return request({
        url: '/api/v1/dress',
        method: 'delete',
        data: data
    })
}

