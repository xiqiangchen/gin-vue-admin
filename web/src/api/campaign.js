import service from '@/utils/request'

// @Tags Campaign
// @Summary 创建活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Campaign true "创建活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /campaign/createCampaign [post]
export const createCampaign = (data) => {
  return service({
    url: '/campaign/createCampaign',
    method: 'post',
    data
  })
}

// @Tags Campaign
// @Summary 删除活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Campaign true "删除活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /campaign/deleteCampaign [delete]
export const deleteCampaign = (data) => {
  return service({
    url: '/campaign/deleteCampaign',
    method: 'delete',
    data
  })
}

// @Tags Campaign
// @Summary 批量删除活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /campaign/deleteCampaign [delete]
export const deleteCampaignByIds = (data) => {
  return service({
    url: '/campaign/deleteCampaignByIds',
    method: 'delete',
    data
  })
}

// @Tags Campaign
// @Summary 更新活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Campaign true "更新活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /campaign/updateCampaign [put]
export const updateCampaign = (data) => {
  return service({
    url: '/campaign/updateCampaign',
    method: 'put',
    data
  })
}

// @Tags Campaign
// @Summary 用id查询活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Campaign true "用id查询活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /campaign/findCampaign [get]
export const findCampaign = (params) => {
  return service({
    url: '/campaign/findCampaign',
    method: 'get',
    params
  })
}

// @Tags Campaign
// @Summary 分页获取活动列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取活动列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /campaign/getCampaignList [get]
export const getCampaignList = (params) => {
  return service({
    url: '/campaign/getCampaignList',
    method: 'get',
    params
  })
}
