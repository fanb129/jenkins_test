package form

type UpdateForm struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required,min=1,max=16"` // 昵称
	Avatar   string `form:"avatar" json:"avatar"`
}
