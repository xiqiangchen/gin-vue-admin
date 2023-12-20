import service from '@/utils/request'

// @Tags Plan
// @Summary 创建广告计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Plan true "创建广告计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /plan/createPlan [post]
export const createPlan = (data) => {
  return service({
    url: '/plan/createPlan',
    method: 'post',
    data
  })
}

// @Tags Plan
// @Summary 删除广告计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Plan true "删除广告计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /plan/deletePlan [delete]
export const deletePlan = (data) => {
  return service({
    url: '/plan/deletePlan',
    method: 'delete',
    data
  })
}

// @Tags Plan
// @Summary 批量删除广告计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除广告计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /plan/deletePlan [delete]
export const deletePlanByIds = (data) => {
  return service({
    url: '/plan/deletePlanByIds',
    method: 'delete',
    data
  })
}

// @Tags Plan
// @Summary 更新广告计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Plan true "更新广告计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /plan/updatePlan [put]
export const updatePlan = (data) => {
  return service({
    url: '/plan/updatePlan',
    method: 'put',
    data
  })
}

// @Tags Plan
// @Summary 用id查询广告计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Plan true "用id查询广告计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /plan/findPlan [get]
export const findPlan = (params) => {
  return service({
    url: '/plan/findPlan',
    method: 'get',
    params
  })
}

// @Tags Plan
// @Summary 分页获取广告计划列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取广告计划列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /plan/getPlanList [get]
export const getPlanList = (params) => {
  return service({
    url: '/plan/getPlanList',
    method: 'get',
    params
  })
}
