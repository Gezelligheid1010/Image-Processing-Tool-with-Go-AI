package logic

import (
	"backend/dao/api"
	"backend/dao/mysql"
	"backend/models"
	"backend/pkg/jwt"
	"backend/pkg/snowflake"
	"fmt"
)

// SignUp 注册业务逻辑
func SignUp(p *models.RegisterForm) (error error) {
	// 1、判断用户存不存在
	err := mysql.CheckUserExist(p.UserName)
	if err != nil {
		// 数据库查询出错
		return err
	}

	// 2、生成UID
	userId, err := snowflake.GetID()
	if err != nil {
		return mysql.ErrorGenIDFailed
	}

	//fmt.Printf("fo: %v\n", fo)

	// 3.上传图片
	imageURL, err := api.UploadImageToSMMS(p.Avatar)
	if err != nil {
		fmt.Println("图像上传失败:", err)
	} else {
		fmt.Println("图像上传成功，URL:", imageURL)
	}

	// 构造一个User实例
	u := models.User{
		UserID:   userId,
		UserName: p.UserName,
		Password: p.Password,
		Email:    p.Email,
		Gender:   p.Gender,
		Avatar:   imageURL,
	}
	// 4、保存进数据库
	return mysql.InsertUser(u)
}

// Login 登录业务逻辑代码
func Login(p *models.LoginForm) (user *models.User, error error) {
	user = &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	//return jwt.GenToken(user.UserID,user.UserName)
	accessToken, refreshToken, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return
}
