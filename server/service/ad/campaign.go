package ad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	adReq "github.com/flipped-aurora/gin-vue-admin/server/model/ad/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type CampaignService struct {
}

// CreateCampaign 创建活动记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) CreateCampaign(campaign *ad.Campaign) (err error) {
	/*
		err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {

			// 根据素材创建创意
			// 将创意
			return nil
		})*/

	err = global.GVA_DB.Create(campaign).Error
	return err
}

// DeleteCampaign 删除活动记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) DeleteCampaign(campaign ad.Campaign) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ad.Campaign{}).Where("id = ?", campaign.ID).Update("deleted_by", campaign.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&campaign).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCampaignByIds 批量删除活动记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) DeleteCampaignByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ad.Campaign{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&ad.Campaign{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCampaign 更新活动记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) UpdateCampaign(campaign ad.Campaign) (err error) {
	err = global.GVA_DB.Save(&campaign).Error
	return err
}

// GetCampaign 根据id获取活动记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) GetCampaign(id uint) (campaign ad.Campaign, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Plan").First(&campaign).Error
	return
}

// GetCampaignInfoList 分页获取活动记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) GetCampaignInfoList(info adReq.CampaignSearch) (list []ad.Campaign, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ad.Campaign{})
	var campaigns []ad.Campaign
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.PlanID != 0 {
		db = db.Where("plan_id = ?", info.PlanID)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.BidMethod != nil {
		db = db.Where("bid_method = ?", info.BidMethod)
	}
	if info.BidMode != nil {
		db = db.Where("bid_mode = ?", info.BidMode)
	}
	if info.Brand != "" {
		db = db.Where("brand LIKE ?", "%"+info.Brand+"%")
	}
	if info.AllowVirtually != nil {
		db = db.Where("allow_virtually = ?", info.AllowVirtually)
	}
	if info.CreativeMode != nil {
		db = db.Where("creative_mode = ?", info.CreativeMode)
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

	err = db.Preload("Plan").Preload("Creatives").Find(&campaigns).Error
	return campaigns, total, err
}
