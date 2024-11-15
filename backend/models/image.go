package models

import (
	"encoding/json"
	"errors"
)

// Image 上传需要处理的手绘图
type Image struct {
	Prompt   string `json:"prompt"`
	OriImage string `json:"ori_image"`
	AuthorId uint64 `json:"author_id,string"`
	OriUrl   string `json:"ori_url"`
}

// UnmarshalJSON 为Post类型实现自定义的UnmarshalJSON方法
func (p *Image) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Prompt   string `json:"prompt"`
		OriImage string `json:"ori_image"`
		AuthorId uint64 `json:"author_id,string"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Prompt) == 0 {
		err = errors.New("Prompt不能为空")
	} else {
		p.Prompt = required.Prompt
		p.OriImage = required.OriImage
		p.AuthorId = required.AuthorId
	}
	return
}
