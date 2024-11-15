package controller

import (
	"backend/dao/mysql"
	"backend/logic"
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreateCategoryHandler 创建分类
func CreateCategoryHandler(c *gin.Context) {
	// 1、获取参数及校验参数
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil { // validator --> binding tag
		zap.L().Debug("c.ShouldBindJSON(category) err", zap.Any("err", err))
		zap.L().Error("create category with invalid parm")
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}

	// 获取作者ID，当前请求的UserID(从c取到当前发请求的用户ID)
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	category.UserId = userID
	//fmt.Println("category:", category)
	// 2、创建帖子
	err = logic.CreateCategory(&category)
	if err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		fmt.Println(err)
		return
	}
	// 3、返回响应
	ResponseSuccess(c, category.CategoryId)
}

// DeleteCategoryHandler 删除分类
func DeleteCategoryHandler(c *gin.Context) {
	// 1. 获取分类ID参数
	categoryID := c.Param("category_id")
	if categoryID == "" {
		ResponseErrorWithMsg(c, CodeInvalidParams, "分类ID不能为空")
		return
	}

	// 2. 逻辑操作
	err := logic.DeleteCategory(categoryID)
	if err != nil {
		zap.L().Error("mysql.DeleteCategory failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回成功响应
	ResponseSuccess(c, "分类删除成功")
}

// GetCategoryListHandler 获取作品列表
func GetCategoryListHandler(c *gin.Context) {

	// 获取作者ID，当前请求的UserID(从c取到当前发请求的用户ID)
	userID, err := getCurrentUserID(c)
	//user_id, err := strconv.ParseInt(c.DefaultQuery("user_id", "0"), 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 获取分页和排序参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	// 获取评论列表
	categorys, err := mysql.GetCategoryListByID(userID, page, size)
	if err != nil {
		zap.L().Error("mysql.GetCategoryListByID failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 格式化返回的数据，包含用户信息
	var commentResponses []models.Category
	for _, category := range categorys {
		commentResponses = append(commentResponses, models.Category{
			//UserId:       category.UserId,
			CategoryName: category.CategoryName,
			CategoryId:   category.CategoryId,
			//CategoryId:   category.CategoryId,
			CoverUrl: category.CoverUrl,
		})
	}
	ResponseSuccess(c, commentResponses)
}

// GetCategoryDetailHandler 获取作品列表
func GetCategoryDetailHandler(c *gin.Context) {
	categoryIDStr := c.Param("category_id") // 获取路径参数
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 获取分页和排序参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	//fmt.Println("category_id:", categoryID)
	// 获取评论列表
	works, err := mysql.GetCategoryDetailByID(categoryID, page, size)
	if err != nil {
		zap.L().Error("mysql.GetCategoryListByID failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		fmt.Println("err；", err)
		return
	}

	// 格式化返回的数据，包含用户信息
	var workResponses []models.WorkDetail
	for _, work := range works {
		workResponses = append(workResponses, models.WorkDetail{
			WorkId: work.WorkId,
			Url:    work.Url,
			Prompt: work.Prompt,
		})
	}
	ResponseSuccess(c, workResponses)
}
