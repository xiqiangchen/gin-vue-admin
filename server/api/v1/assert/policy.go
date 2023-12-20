package assert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/assert"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    assertReq "github.com/flipped-aurora/gin-vue-admin/server/model/assert/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type PolicyApi struct {
}

var policyService = service.ServiceGroupApp.AssertServiceGroup.PolicyService


// CreatePolicy 创建出价策略
// @Tags Policy
// @Summary 创建出价策略
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assert.Policy true "创建出价策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /policy/createPolicy [post]
func (policyApi *PolicyApi) CreatePolicy(c *gin.Context) {
	var policy assert.Policy
	err := c.ShouldBindJSON(&policy)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    policy.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Price":{utils.NotEmpty()},
    }
	if err := utils.Verify(policy, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := policyService.CreatePolicy(&policy); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePolicy 删除出价策略
// @Tags Policy
// @Summary 删除出价策略
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assert.Policy true "删除出价策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /policy/deletePolicy [delete]
func (policyApi *PolicyApi) DeletePolicy(c *gin.Context) {
	var policy assert.Policy
	err := c.ShouldBindJSON(&policy)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    policy.DeletedBy = utils.GetUserID(c)
	if err := policyService.DeletePolicy(policy); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePolicyByIds 批量删除出价策略
// @Tags Policy
// @Summary 批量删除出价策略
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除出价策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /policy/deletePolicyByIds [delete]
func (policyApi *PolicyApi) DeletePolicyByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := policyService.DeletePolicyByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePolicy 更新出价策略
// @Tags Policy
// @Summary 更新出价策略
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assert.Policy true "更新出价策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /policy/updatePolicy [put]
func (policyApi *PolicyApi) UpdatePolicy(c *gin.Context) {
	var policy assert.Policy
	err := c.ShouldBindJSON(&policy)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    policy.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Price":{utils.NotEmpty()},
      }
    if err := utils.Verify(policy, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := policyService.UpdatePolicy(policy); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPolicy 用id查询出价策略
// @Tags Policy
// @Summary 用id查询出价策略
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query assert.Policy true "用id查询出价策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /policy/findPolicy [get]
func (policyApi *PolicyApi) FindPolicy(c *gin.Context) {
	var policy assert.Policy
	err := c.ShouldBindQuery(&policy)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repolicy, err := policyService.GetPolicy(policy.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repolicy": repolicy}, c)
	}
}

// GetPolicyList 分页获取出价策略列表
// @Tags Policy
// @Summary 分页获取出价策略列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query assertReq.PolicySearch true "分页获取出价策略列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /policy/getPolicyList [get]
func (policyApi *PolicyApi) GetPolicyList(c *gin.Context) {
	var pageInfo assertReq.PolicySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := policyService.GetPolicyInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
