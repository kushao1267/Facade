package forms

type LinkForm struct { // POST参数表单验证
	Url string `form:"url" json:"url" binding:"required,max=256"`
}