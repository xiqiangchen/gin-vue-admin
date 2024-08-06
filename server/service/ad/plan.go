package ad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	adReq "github.com/flipped-aurora/gin-vue-admin/server/model/ad/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type PlanService struct {
}

// CreatePlan 创建广告计划记录
// Author [piexlmax](https://github.com/piexlmax)
func (planService *PlanService) CreatePlan(plan *ad.Plan) (err error) {
	err = global.GVA_DB.Create(plan).Error
	return err
}

// DeletePlan 删除广告计划记录
// Author [piexlmax](https://github.com/piexlmax)
func (planService *PlanService) DeletePlan(plan ad.Plan) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ad.Plan{}).Where("id = ?", plan.ID).Update("deleted_by", plan.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&plan).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeletePlanByIds 批量删除广告计划记录
// Author [piexlmax](https://github.com/piexlmax)
func (planService *PlanService) DeletePlanByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ad.Plan{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&ad.Plan{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdatePlan 更新广告计划记录
// Author [piexlmax](https://github.com/piexlmax)
func (planService *PlanService) UpdatePlan(plan ad.Plan) (err error) {
	err = global.GVA_DB.Save(&plan).Error
	return err
}

// GetPlan 根据id获取广告计划记录
// Author [piexlmax](https://github.com/piexlmax)
func (planService *PlanService) GetPlan(id uint) (plan ad.Plan, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&plan).Error
	return
}

// GetPlanInfoList 分页获取广告计划记录
// Author [piexlmax](https://github.com/piexlmax)
func (planService *PlanService) GetPlanInfoList(info adReq.PlanSearch) (list []ad.Plan, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ad.Plan{})
	var plans []ad.Plan
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CreatedBy > 0 {
		db = db.Where("created_by = ?", info.CreatedBy)
	}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.Filter != nil {
		db = db.Where("filter = ?", info.Filter)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("ID DESC")
	}

	err = db.Find(&plans).Error
	return plans, total, err
}
