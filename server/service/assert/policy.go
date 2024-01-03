package assert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/assert"
	assertReq "github.com/flipped-aurora/gin-vue-admin/server/model/assert/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type PolicyService struct {
}

// CreatePolicy 创建出价策略记录
// Author [piexlmax](https://github.com/piexlmax)
func (policyService *PolicyService) CreatePolicy(policy *assert.Policy) (err error) {
	err = global.GVA_DB.Create(policy).Error
	return err
}

// DeletePolicy 删除出价策略记录
// Author [piexlmax](https://github.com/piexlmax)
func (policyService *PolicyService) DeletePolicy(policy assert.Policy) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&assert.Policy{}).Where("id = ?", policy.ID).Update("deleted_by", policy.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&policy).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeletePolicyByIds 批量删除出价策略记录
// Author [piexlmax](https://github.com/piexlmax)
func (policyService *PolicyService) DeletePolicyByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&assert.Policy{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&assert.Policy{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdatePolicy 更新出价策略记录
// Author [piexlmax](https://github.com/piexlmax)
func (policyService *PolicyService) UpdatePolicy(policy assert.Policy) (err error) {
	err = global.GVA_DB.Save(&policy).Error
	return err
}

// GetPolicy 根据id获取出价策略记录
// Author [piexlmax](https://github.com/piexlmax)
func (policyService *PolicyService) GetPolicy(id uint) (policy assert.Policy, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&policy).Error
	return
}

// GetPolicyInfoList 分页获取出价策略记录
// Author [piexlmax](https://github.com/piexlmax)
func (policyService *PolicyService) GetPolicyInfoList(info assertReq.PolicySearch) (list []assert.Policy, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&assert.Policy{})
	var policys []assert.Policy
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Platform != nil {
		db = db.Where("platform = ?", info.Platform)
	}
	if info.Spot != nil {
		db = db.Where("spot = ?", info.Spot)
	}
	if info.Bundle != nil {
		db = db.Where("bundle = ?", info.Bundle)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("ID DESC")
	}

	err = db.Find(&policys).Error
	return policys, total, err
}
