package assert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/assert"
	assertReq "github.com/flipped-aurora/gin-vue-admin/server/model/assert/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type TargetService struct {
}

// CreateTarget 创建定向包记录
// Author [piexlmax](https://github.com/piexlmax)
func (targetService *TargetService) CreateTarget(target *assert.Target) (err error) {
	err = global.GVA_DB.Create(target).Error
	return err
}

// DeleteTarget 删除定向包记录
// Author [piexlmax](https://github.com/piexlmax)
func (targetService *TargetService) DeleteTarget(target assert.Target) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&assert.Target{}).Where("id = ?", target.ID).Update("deleted_by", target.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&target).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTargetByIds 批量删除定向包记录
// Author [piexlmax](https://github.com/piexlmax)
func (targetService *TargetService) DeleteTargetByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&assert.Target{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&assert.Target{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTarget 更新定向包记录
// Author [piexlmax](https://github.com/piexlmax)
func (targetService *TargetService) UpdateTarget(target assert.Target) (err error) {
	err = global.GVA_DB.Save(&target).Error
	return err
}

// GetTarget 根据id获取定向包记录
// Author [piexlmax](https://github.com/piexlmax)
func (targetService *TargetService) GetTarget(id uint) (target assert.Target, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&target).Error
	return
}

// GetTargetInfoList 分页获取定向包记录
// Author [piexlmax](https://github.com/piexlmax)
func (targetService *TargetService) GetTargetInfoList(info assertReq.TargetSearch) (list []assert.Target, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&assert.Target{})
	var targets []assert.Target
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Desc != "" {
		db = db.Where("desc LIKE ?", "%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("ID DESC")
	}

	err = db.Find(&targets).Error
	return targets, total, err
}
