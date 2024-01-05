package resource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resource"
	resourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/resource/request"
	"gorm.io/gorm"
)

type MaterialService struct {
}

// CreateMaterial 创建素材库记录
// Author [piexlmax](https://github.com/piexlmax)
func (materialService *MaterialService) CreateMaterial(material *resource.Material) (err error) {
	err = global.GVA_DB.Create(material).Error
	return err
}

// DeleteMaterial 删除素材库记录
// Author [piexlmax](https://github.com/piexlmax)
func (materialService *MaterialService) DeleteMaterial(material resource.Material) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&resource.Material{}).Where("id = ?", material.ID).Update("deleted_by", material.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&material).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMaterialByIds 批量删除素材库记录
// Author [piexlmax](https://github.com/piexlmax)
func (materialService *MaterialService) DeleteMaterialByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&resource.Material{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&resource.Material{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMaterial 更新素材库记录
// Author [piexlmax](https://github.com/piexlmax)
func (materialService *MaterialService) UpdateMaterial(material resource.Material) (err error) {
	err = global.GVA_DB.Save(&material).Error
	return err
}

// GetMaterial 根据id获取素材库记录
// Author [piexlmax](https://github.com/piexlmax)
func (materialService *MaterialService) GetMaterial(id uint) (material resource.Material, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&material).Error
	return
}

// GetMaterialInfoList 分页获取素材库记录
// Author [piexlmax](https://github.com/piexlmax)
func (materialService *MaterialService) GetMaterialInfoList(info resourceReq.MaterialSearch) (list []resource.Material, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&resource.Material{})
	var materials []resource.Material
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Type > 0 {
		db = db.Where("type = ?", info.Type)
	}
	if info.Format != "" {
		db = db.Where("format = ?", info.Format)
	}
	if info.Width > 0 {
		db = db.Where("width = ?", info.Width)
	}
	if info.Height > 0 {
		db = db.Where("height = ?", info.Height)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&materials).Error
	return materials, total, err
}
