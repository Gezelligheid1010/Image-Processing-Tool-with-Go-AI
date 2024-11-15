package models

import (
	"encoding/json"
)

// Work 上传需要处理的手绘图
type Work struct {
	WorkId     uint64 `json:"work_id" db:"work_id"`
	Url        string `json:"url" db:"url"`
	WorkImage  string `json:"work_image"`
	CategoryId uint64 `json:"category_id,string" db:"category_id"`
	Prompt     string `json:"prompt" db:"prompt"`
}

// UnmarshalJSON 为Work类型实现自定义的UnmarshalJSON方法
func (w *Work) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Url        string `json:"url"`
		WorkImage  string `json:"work_image"`
		CategoryId uint64 `json:"category_id,string"`
		Prompt     string `json:"prompt"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else {
		w.Url = required.Url
		w.WorkImage = required.WorkImage
		w.CategoryId = required.CategoryId
		w.Prompt = required.Prompt
	}
	return
}

// WorkDetail 获取分类中的手绘图
type WorkDetail struct {
	WorkId uint64 `json:"work_id,string" db:"work_id"`
	Url    string `json:"url" db:"url"`
	Prompt string `json:"prompt" db:"prompt"`
}
