import service from '@/utils/request'

// @Tags Material
// @Summary 创建素材库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Material true "创建素材库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /material/createMaterial [post]
export const createMaterial = (data) => {
  return service({
    url: '/material/createMaterial',
    method: 'post',
    data
  })
}

// @Tags Material
// @Summary 删除素材库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Material true "删除素材库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /material/deleteMaterial [delete]
export const deleteMaterial = (data) => {
  return service({
    url: '/material/deleteMaterial',
    method: 'delete',
    data
  })
}

// @Tags Material
// @Summary 批量删除素材库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除素材库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /material/deleteMaterial [delete]
export const deleteMaterialByIds = (data) => {
  return service({
    url: '/material/deleteMaterialByIds',
    method: 'delete',
    data
  })
}

// @Tags Material
// @Summary 更新素材库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Material true "更新素材库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /material/updateMaterial [put]
export const updateMaterial = (data) => {
  return service({
    url: '/material/updateMaterial',
    method: 'put',
    data
  })
}

// @Tags Material
// @Summary 用id查询素材库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Material true "用id查询素材库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /material/findMaterial [get]
export const findMaterial = (params) => {
  return service({
    url: '/material/findMaterial',
    method: 'get',
    params
  })
}

// @Tags Material
// @Summary 分页获取素材库列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取素材库列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /material/getMaterialList [get]
export const getMaterialList = (params) => {
  return service({
    url: '/material/getMaterialList',
    method: 'get',
    params
  })
}
