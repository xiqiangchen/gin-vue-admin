import service from '@/utils/request'

// @Tags Image
// @Summary 创建图片库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Image true "创建图片库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /image/createImage [post]
export const createImage = (data) => {
  return service({
    url: '/image/createImage',
    method: 'post',
    data
  })
}

// @Tags Image
// @Summary 删除图片库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Image true "删除图片库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /image/deleteImage [delete]
export const deleteImage = (data) => {
  return service({
    url: '/image/deleteImage',
    method: 'delete',
    data
  })
}

// @Tags Image
// @Summary 批量删除图片库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除图片库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /image/deleteImage [delete]
export const deleteImageByIds = (data) => {
  return service({
    url: '/image/deleteImageByIds',
    method: 'delete',
    data
  })
}

// @Tags Image
// @Summary 更新图片库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Image true "更新图片库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /image/updateImage [put]
export const updateImage = (data) => {
  return service({
    url: '/image/updateImage',
    method: 'put',
    data
  })
}

// @Tags Image
// @Summary 用id查询图片库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Image true "用id查询图片库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /image/findImage [get]
export const findImage = (params) => {
  return service({
    url: '/image/findImage',
    method: 'get',
    params
  })
}

// @Tags Image
// @Summary 分页获取图片库列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取图片库列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /image/getImageList [get]
export const getImageList = (params) => {
  return service({
    url: '/image/getImageList',
    method: 'get',
    params
  })
}
