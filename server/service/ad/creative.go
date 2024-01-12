package ad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	adReq "github.com/flipped-aurora/gin-vue-admin/server/model/ad/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type CreativeService struct {
}

// CreateCreative 创建创意表记录
// Author [piexlmax](https://github.com/piexlmax)
func (creativeService *CreativeService) CreateCreative(creative *ad.Creative) (err error) {
	err = global.GVA_DB.Create(creative).Error
	return err
}

// CreateCreativeBatch 批量创建创意表记录
func (creativeService *CreativeService) CreateCreativeBatch(creatives []*ad.Creative) (err error) {
	err = global.GVA_DB.CreateInBatches(creatives, len(creatives)).Error
	return err
}

// DeleteCreative 删除创意表记录
// Author [piexlmax](https://github.com/piexlmax)
func (creativeService *CreativeService) DeleteCreative(creative ad.Creative) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ad.Creative{}).Where("id = ?", creative.ID).Update("deleted_by", creative.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&creative).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCreativeByIds 批量删除创意表记录
// Author [piexlmax](https://github.com/piexlmax)
func (creativeService *CreativeService) DeleteCreativeByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ad.Creative{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&ad.Creative{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCreative 更新创意表记录
// Author [piexlmax](https://github.com/piexlmax)
func (creativeService *CreativeService) UpdateCreative(creative ad.Creative) (err error) {
	err = global.GVA_DB.Save(&creative).Error
	return err
}

// GetCreative 根据id获取创意表记录
// Author [piexlmax](https://github.com/piexlmax)
func (creativeService *CreativeService) GetCreative(id uint) (creative ad.Creative, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Material").Preload("Plan").Preload("Campaign").First(&creative).Error
	return
}

// GetCreativeInfoList 分页获取创意表记录
// Author [piexlmax](https://github.com/piexlmax)
func (creativeService *CreativeService) GetCreativeInfoList(info adReq.CreativeSearch) (list []ad.Creative, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ad.Creative{})
	var creatives []ad.Creative
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.PlanId > 0 {
		db = db.Where("plan_id = ?", info.PlanId)
	}
	if info.CampaignId > 0 {
		db = db.Where("campaign_id = ?", info.CampaignId)
	}
	if info.MaterialId > 0 {
		db = db.Where("material_id = ?", info.MaterialId)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("ID DESC")
	}

	err = db.Preload("Material").Preload("Plan").Preload("Campaign").Find(&creatives).Error
	return creatives, total, err
}
