import service from '@/utils/request'

// @Tags Policy
// @Summary 创建出价策略
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Policy true "创建出价策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /policy/createPolicy [post]
export const createPolicy = (data) => {
  return service({
    url: '/policy/createPolicy',
    method: 'post',
    data
  })
}

// @Tags Policy
// @Summary 删除出价策略
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Policy true "删除出价策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /policy/deletePolicy [delete]
export const deletePolicy = (data) => {
  return service({
    url: '/policy/deletePolicy',
    method: 'delete',
    data
  })
}

// @Tags Policy
// @Summary 批量删除出价策略
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除出价策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /policy/deletePolicy [delete]
export const deletePolicyByIds = (data) => {
  return service({
    url: '/policy/deletePolicyByIds',
    method: 'delete',
    data
  })
}

// @Tags Policy
// @Summary 更新出价策略
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Policy true "更新出价策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /policy/updatePolicy [put]
export const updatePolicy = (data) => {
  return service({
    url: '/policy/updatePolicy',
    method: 'put',
    data
  })
}

// @Tags Policy
// @Summary 用id查询出价策略
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Policy true "用id查询出价策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /policy/findPolicy [get]
export const findPolicy = (params) => {
  return service({
    url: '/policy/findPolicy',
    method: 'get',
    params
  })
}

// @Tags Policy
// @Summary 分页获取出价策略列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取出价策略列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /policy/getPolicyList [get]
export const getPolicyList = (params) => {
  return service({
    url: '/policy/getPolicyList',
    method: 'get',
    params
  })
}
