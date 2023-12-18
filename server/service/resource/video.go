package resource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resource"
	resourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/resource/request"
	"gorm.io/gorm"
)

type VideoService struct {
}

// CreateVideo 创建视频记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoService *VideoService) CreateVideo(video *resource.Video) (err error) {
	err = global.GVA_DB.Create(video).Error
	return err
}

// DeleteVideo 删除视频记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoService *VideoService) DeleteVideo(video resource.Video) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&resource.Video{}).Where("id = ?", video.ID).Update("deleted_by", video.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&video).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteVideoByIds 批量删除视频记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoService *VideoService) DeleteVideoByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&resource.Video{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&resource.Video{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateVideo 更新视频记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoService *VideoService) UpdateVideo(video resource.Video) (err error) {
	err = global.GVA_DB.Save(&video).Error
	return err
}

// GetVideo 根据id获取视频记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoService *VideoService) GetVideo(id uint) (video resource.Video, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&video).Error
	return
}

// GetVideoInfoList 分页获取视频记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoService *VideoService) GetVideoInfoList(info resourceReq.VideoSearch) (list []resource.Video, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&resource.Video{})
	var videos []resource.Video
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Width != nil {
		db = db.Where("width = ?", info.Width)
	}
	if info.Height != nil {
		db = db.Where("height = ?", info.Height)
	}
	if info.Format != "" {
		db = db.Where("format = ?", info.Format)
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

	err = db.Find(&videos).Error
	return videos, total, err
}
