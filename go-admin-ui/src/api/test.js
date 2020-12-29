import request from '@/utils/request'

// 查询Test列表
export function listTest(query) {
    return request({
        url: '/api/v1/test',
        method: 'get',
        params: query
    })
}

// 查询Test详细
export function getTest (ID) {
    return request({
        url: '/api/v1/test/' + ID,
        method: 'get'
    })
}


// 新增Test
export function addTest(data) {
    return request({
        url: '/api/v1/test',
        method: 'post',
        data: data
    })
}

// 修改Test
export function updateTest(data) {
    return request({
        url: '/api/v1/test/'+data.ID,
        method: 'put',
        data: data
    })
}

// 删除Test
export function delTest(data) {
    return request({
        url: '/api/v1/test',
        method: 'delete',
        data: data
    })
}

