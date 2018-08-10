package models

// 品牌资讯
type News struct {
	BaseModel
	Title   string //资讯标题
	Des     string //描述
	Content string //资讯内容
}
