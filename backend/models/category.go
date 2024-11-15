package models

import "encoding/json"

// Category 创建分类
type Category struct {
	UserId       uint64 `json:"user_id,string" db:"user_id"`
	CategoryName string `json:"category_name" db:"category_name"`
	CategoryId   uint64 `json:"category_id,string" db:"category_id"`
	Description  string `json:"description" db:"description"`
	Cover        string `json:"cover"`
	CoverUrl     string `json:"cover_url" db:"cover_url"`
}

// UnmarshalJSON 为Category类型实现自定义的UnmarshalJSON方法
func (c *Category) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserId       uint64 `json:"user_id,string"`
		CategoryId   uint64 `json:"category_id,string"` // 添加 CategoryId 字段
		CategoryName string `json:"category_name"`
		Description  string `json:"description"`
		Cover        string `json:"cover"`
		CoverUrl     string `json:"cover_url"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else {
		c.UserId = required.UserId
		c.CategoryId = required.CategoryId
		c.CategoryName = required.CategoryName
		c.Description = required.Description
		c.Cover = required.Cover
		c.CoverUrl = required.CoverUrl
	}
	return
}

//// CategoryDetailRes 用于返回分类详情和包含的作品列表
//type CategoryDetailRes struct {
//	CategoryId   uint64       `json:"category_id" db:"category_id"`
//	CategoryName string       `json:"category_name" db:"category_name"`
//	Description  string       `json:"description" db:"description"`
//	WorkDetail   []WorkDetail `json:"workDetail"` // 包含的作品列表
//}
