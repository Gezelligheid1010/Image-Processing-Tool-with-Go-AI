package logic

import (
	"backend/dao/api"
	"backend/dao/mysql"
	"backend/models"
	"backend/pkg/snowflake"
	"fmt"
)

// UploadWork 上传作品
func UploadWork(w *models.Work) (err error) {
	//// 1、判断分类存不存在
	//err = mysql.CheckCategoryExist(w.CategoryId)
	//if err != nil {
	//	// 数据库查询出错
	//	return err
	//}

	// 2、生成WorkID
	workId, err := snowflake.GetID()
	if err != nil {
		return mysql.ErrorGenIDFailed
	}

	// 3、上传作品到图床
	//base64Work, err := utils.ConverUrl2Base64(w.Url)
	URL, err := api.UploadImageToSMMS(w.WorkImage)
	if err != nil {
		fmt.Println("图像上传失败:", err)
		return err
	}

	//fmt.Println("图像上传成功，URL:", URL)

	// 构造一个User实例
	work := models.Work{
		WorkId:     workId,
		Url:        URL,
		CategoryId: w.CategoryId,
		Prompt:     w.Prompt,
	}
	// 4、保存进数据库
	return mysql.InsertWork(work)
}
