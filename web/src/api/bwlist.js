import service from '@/utils/request'

// @Tags BlackWhiteList
// @Summary 创建黑白名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BlackWhiteList true "创建黑白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bwlist/createBlackWhiteList [post]
export const createBlackWhiteList = (data) => {
  return service({
    url: '/bwlist/createBlackWhiteList',
    method: 'post',
    data
  })
}

// @Tags BlackWhiteList
// @Summary 删除黑白名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BlackWhiteList true "删除黑白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bwlist/deleteBlackWhiteList [delete]
export const deleteBlackWhiteList = (data) => {
  return service({
    url: '/bwlist/deleteBlackWhiteList',
    method: 'delete',
    data
  })
}

// @Tags BlackWhiteList
// @Summary 批量删除黑白名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除黑白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bwlist/deleteBlackWhiteList [delete]
export const deleteBlackWhiteListByIds = (data) => {
  return service({
    url: '/bwlist/deleteBlackWhiteListByIds',
    method: 'delete',
    data
  })
}

// @Tags BlackWhiteList
// @Summary 更新黑白名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BlackWhiteList true "更新黑白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bwlist/updateBlackWhiteList [put]
export const updateBlackWhiteList = (data) => {
  return service({
    url: '/bwlist/updateBlackWhiteList',
    method: 'put',
    data
  })
}

// @Tags BlackWhiteList
// @Summary 用id查询黑白名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BlackWhiteList true "用id查询黑白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bwlist/findBlackWhiteList [get]
export const findBlackWhiteList = (params) => {
  return service({
    url: '/bwlist/findBlackWhiteList',
    method: 'get',
    params
  })
}

// @Tags BlackWhiteList
// @Summary 分页获取黑白名单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取黑白名单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bwlist/getBlackWhiteListList [get]
export const getBlackWhiteListList = (params) => {
  return service({
    url: '/bwlist/getBlackWhiteListList',
    method: 'get',
    params
  })
}
