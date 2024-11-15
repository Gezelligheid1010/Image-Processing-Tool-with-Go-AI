package logic

import (
	"backend/dao/api"
	"backend/dao/mysql"
	"backend/models"
	"backend/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
)

// CreateCategory 创建Category
func CreateCategory(c *models.Category) (err error) {
	// 1、 生成post_id(生成分类ID)
	CategoryId, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		return
	}
	c.CategoryId = CategoryId

	// 2、将封面上传
	//base64Work, err := utils.ConverUrl2Base64(c.Cover)
	c.CoverUrl, err = api.UploadImageToSMMS(c.Cover)
	if err != nil {
		fmt.Println("图像上传失败:", err)
		return err
	}

	fmt.Println("图像上传成功，URL:", c.CoverUrl)

	// 4、保存进数据库
	return mysql.CreateCategory(c)
}

func DeleteCategory(categoryID string) error {

	// 1. 删除该分类下的所有作品
	err := mysql.DeleteWorksByCategoryID(categoryID)
	if err != nil {
		return err
	}

	// 3. 删除分类本身
	err = mysql.DeleteCategory(categoryID)
	if err != nil {
		return err
	}

	return nil
}
