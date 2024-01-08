import service from '@/utils/request'

// @Tags Creative
// @Summary 创建创意表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Creative true "创建创意表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /creative/createCreative [post]
export const createCreative = (data) => {
  return service({
    url: '/creative/createCreative',
    method: 'post',
    data
  })
}

// @Tags Creative
// @Summary 创建创意表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Creative true "创建创意表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /creative/createCreatives [post]
export const createCreatives = (data) => {
  return service({
    url: '/creative/createCreatives',
    method: 'post',
    data
  })
}

// @Tags Creative
// @Summary 删除创意表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Creative true "删除创意表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /creative/deleteCreative [delete]
export const deleteCreative = (data) => {
  return service({
    url: '/creative/deleteCreative',
    method: 'delete',
    data
  })
}

// @Tags Creative
// @Summary 批量删除创意表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除创意表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /creative/deleteCreative [delete]
export const deleteCreativeByIds = (data) => {
  return service({
    url: '/creative/deleteCreativeByIds',
    method: 'delete',
    data
  })
}

// @Tags Creative
// @Summary 更新创意表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Creative true "更新创意表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /creative/updateCreative [put]
export const updateCreative = (data) => {
  return service({
    url: '/creative/updateCreative',
    method: 'put',
    data
  })
}

// @Tags Creative
// @Summary 用id查询创意表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Creative true "用id查询创意表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /creative/findCreative [get]
export const findCreative = (params) => {
  return service({
    url: '/creative/findCreative',
    method: 'get',
    params
  })
}

// @Tags Creative
// @Summary 分页获取创意表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取创意表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /creative/getCreativeList [get]
export const getCreativeList = (params) => {
  return service({
    url: '/creative/getCreativeList',
    method: 'get',
    params
  })
}
