package mysql

import (
	"backend/models"
	"fmt"
	"go.uber.org/zap"
)

// CreateCategory 创建Category
func CreateCategory(category *models.Category) (err error) {

	sqlStr := `insert into category(
	user_id, category_name, category_id,description,cover_url)
	values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, category.UserId, category.CategoryName,
		category.CategoryId, category.Description, category.CoverUrl)
	if err != nil {
		zap.L().Error("insert post failed", zap.Error(err))
		//err = ErrorInsertFailed
		return err
	}
	return nil
}

// GetCategoryListByID 根据分类ID获取分类列表
func GetCategoryListByID(uerID uint64, page, size int) (categoryList []*models.Category, err error) {
	// 排序条件
	orderBy := "create_time DESC" // 默认按最新排序

	// 计算偏移量
	offset := (page - 1) * size

	// 构建查询语句，包含排序和分页
	sqlStr := fmt.Sprintf(`SELECT category_id, category_name,cover_url
		FROM category
		WHERE user_id = ? 
		ORDER BY %s
		LIMIT ? OFFSET ?`, orderBy)

	// 查询评论列表
	err = db.Select(&categoryList, sqlStr, uerID, size, offset)
	return categoryList, err
}

func DeleteCategory(categoryID string) error {
	// 1. 删除分类本身
	sqlStr := `DELETE FROM category WHERE category_id = ?`
	_, err := db.Exec(sqlStr, categoryID)
	if err != nil {
		zap.L().Error("DeleteCategory failed", zap.Error(err))
		return err
	}
	return nil
}
