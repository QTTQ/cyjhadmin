import request from '@/utils/request'

// 查询SystemTasks列表
export function listSystemTasks(query) {
    return request({
        url: '/api/v1/systemtasks',
        method: 'get',
        params: query
    })
}

// 查询SystemTasks详细
export function getSystemTasks (ID) {
    return request({
        url: '/api/v1/systemtasks/' + ID,
        method: 'get'
    })
}


// 新增SystemTasks
export function addSystemTasks(data) {
    return request({
        url: '/api/v1/systemtasks',
        method: 'post',
        data: data
    })
}

// 修改SystemTasks
export function updateSystemTasks(data) {
    return request({
        url: '/api/v1/systemtasks/'+data.ID,
        method: 'put',
        data: data
    })
}

// 删除SystemTasks
export function delSystemTasks(data) {
    return request({
        url: '/api/v1/systemtasks',
        method: 'delete',
        data: data
    })
}

