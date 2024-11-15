package mysql

import (
	"backend/models"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

// CheckCategoryExist 检查指定作品分类是否存在
func CheckCategoryExist(category_id uint64) (error error) {
	sqlStr := `select count(category_id) from category where category_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, category_id); err != nil {
		return err
	}
	if count > 0 {
		return errors.New(ErrorUserExit)
	}
	return
}

// InsertWork 插入作品
func InsertWork(work models.Work) (error error) {
	// 执行SQL语句入库
	sqlstr := `insert into works(work_id,url,prompt,category_id) values(?,?,?,?)`
	_, err := db.Exec(sqlstr, work.WorkId, work.Url, work.Prompt, work.CategoryId)
	return err
}

// GetCategoryDetailByID 根据分类ID获取分类详情和作品列表
func GetCategoryDetailByID(categoryID uint64, page, size int) (workList []*models.WorkDetail, err error) {
	// 排序条件
	orderBy := "create_time DESC" // 默认按最新排序

	// 计算偏移量
	offset := (page - 1) * size

	// 构建查询语句，包含排序和分页
	sqlStr := fmt.Sprintf(`SELECT work_id, url, prompt
		FROM works
		WHERE category_id = ? 
		ORDER BY %s
		LIMIT ? OFFSET ?`, orderBy)

	// 查询评论列表
	err = db.Select(&workList, sqlStr, categoryID, size, offset)

	return workList, err
}

func DeleteWorksByCategoryID(categoryID string) error {
	// 直接通过 CategoryID 删除所有相关作品
	sqlStr := `DELETE FROM works WHERE category_id = ?`
	_, err := db.Exec(sqlStr, categoryID)
	if err != nil {
		zap.L().Error("DeleteWorksByCategoryID failed", zap.Error(err))
		return err
	}
	return nil
}
