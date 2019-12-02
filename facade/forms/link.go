package forms

// LinkForm POST参数表单验证
type LinkForm struct {
	URL string `form:"url" json:"url" binding:"required,max=256"`
}
