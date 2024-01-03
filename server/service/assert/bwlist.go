package assert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/assert"
	assertReq "github.com/flipped-aurora/gin-vue-admin/server/model/assert/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type BlackWhiteListService struct {
}

// CreateBlackWhiteList 创建黑白名单记录
// Author [piexlmax](https://github.com/piexlmax)
func (bwlistService *BlackWhiteListService) CreateBlackWhiteList(bwlist *assert.BlackWhiteList) (err error) {
	err = global.GVA_DB.Create(bwlist).Error
	return err
}

// DeleteBlackWhiteList 删除黑白名单记录
// Author [piexlmax](https://github.com/piexlmax)
func (bwlistService *BlackWhiteListService) DeleteBlackWhiteList(bwlist assert.BlackWhiteList) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&assert.BlackWhiteList{}).Where("id = ?", bwlist.ID).Update("deleted_by", bwlist.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&bwlist).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteBlackWhiteListByIds 批量删除黑白名单记录
// Author [piexlmax](https://github.com/piexlmax)
func (bwlistService *BlackWhiteListService) DeleteBlackWhiteListByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&assert.BlackWhiteList{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&assert.BlackWhiteList{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateBlackWhiteList 更新黑白名单记录
// Author [piexlmax](https://github.com/piexlmax)
func (bwlistService *BlackWhiteListService) UpdateBlackWhiteList(bwlist assert.BlackWhiteList) (err error) {
	err = global.GVA_DB.Save(&bwlist).Error
	return err
}

// GetBlackWhiteList 根据id获取黑白名单记录
// Author [piexlmax](https://github.com/piexlmax)
func (bwlistService *BlackWhiteListService) GetBlackWhiteList(id uint) (bwlist assert.BlackWhiteList, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&bwlist).Error
	return
}

// GetBlackWhiteListInfoList 分页获取黑白名单记录
// Author [piexlmax](https://github.com/piexlmax)
func (bwlistService *BlackWhiteListService) GetBlackWhiteListInfoList(info assertReq.BlackWhiteListSearch) (list []assert.BlackWhiteList, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&assert.BlackWhiteList{})
	var bwlists []assert.BlackWhiteList
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("ID DESC")
	}

	err = db.Find(&bwlists).Error
	return bwlists, total, err
}
