import service from '@/utils/request'

// @Tags Target
// @Summary 创建定向包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Target true "创建定向包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /target/createTarget [post]
export const createTarget = (data) => {
  return service({
    url: '/target/createTarget',
    method: 'post',
    data
  })
}

// @Tags Target
// @Summary 删除定向包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Target true "删除定向包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /target/deleteTarget [delete]
export const deleteTarget = (data) => {
  return service({
    url: '/target/deleteTarget',
    method: 'delete',
    data
  })
}

// @Tags Target
// @Summary 批量删除定向包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除定向包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /target/deleteTarget [delete]
export const deleteTargetByIds = (data) => {
  return service({
    url: '/target/deleteTargetByIds',
    method: 'delete',
    data
  })
}

// @Tags Target
// @Summary 更新定向包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Target true "更新定向包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /target/updateTarget [put]
export const updateTarget = (data) => {
  return service({
    url: '/target/updateTarget',
    method: 'put',
    data
  })
}

// @Tags Target
// @Summary 用id查询定向包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Target true "用id查询定向包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /target/findTarget [get]
export const findTarget = (params) => {
  return service({
    url: '/target/findTarget',
    method: 'get',
    params
  })
}

// @Tags Target
// @Summary 分页获取定向包列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取定向包列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /target/getTargetList [get]
export const getTargetList = (params) => {
  return service({
    url: '/target/getTargetList',
    method: 'get',
    params
  })
}
