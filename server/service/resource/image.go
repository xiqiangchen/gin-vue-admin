package resource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resource"
	resourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/resource/request"
	"gorm.io/gorm"
)

type ImageService struct {
}

// CreateImage 创建图片库记录
// Author [piexlmax](https://github.com/piexlmax)
func (imageService *ImageService) CreateImage(image *resource.Image) (err error) {
	err = global.GVA_DB.Create(image).Error
	return err
}

// DeleteImage 删除图片库记录
// Author [piexlmax](https://github.com/piexlmax)
func (imageService *ImageService) DeleteImage(image resource.Image) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&resource.Image{}).Where("id = ?", image.ID).Update("deleted_by", image.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&image).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteImageByIds 批量删除图片库记录
// Author [piexlmax](https://github.com/piexlmax)
func (imageService *ImageService) DeleteImageByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&resource.Image{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&resource.Image{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateImage 更新图片库记录
// Author [piexlmax](https://github.com/piexlmax)
func (imageService *ImageService) UpdateImage(image resource.Image) (err error) {
	err = global.GVA_DB.Save(&image).Error
	return err
}

// GetImage 根据id获取图片库记录
// Author [piexlmax](https://github.com/piexlmax)
func (imageService *ImageService) GetImage(id uint) (image resource.Image, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&image).Error
	return
}

// GetImageInfoList 分页获取图片库记录
// Author [piexlmax](https://github.com/piexlmax)
func (imageService *ImageService) GetImageInfoList(info resourceReq.ImageSearch) (list []resource.Image, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&resource.Image{})
	var images []resource.Image
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Format != "" {
		db = db.Where("format = ?", info.Format)
	}
	if info.Width != nil {
		db = db.Where("width = ?", info.Width)
	}
	if info.Height != nil {
		db = db.Where("height = ?", info.Height)
	}
	if info.Comment != "" {
		db = db.Where("comment LIKE ?", "%"+info.Comment+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&images).Error
	return images, total, err
}
