package models

//联系方式
type Contact struct {
	BaseModel
	Tel     string //联系电话
	Fax     string //传真
	Working string //
	Email   string //邮箱
	Address string //地址
	Weibo   string //微博
	Qrcode  string //二维码
}
