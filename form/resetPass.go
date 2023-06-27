package form

type ResetPassForm struct {
	Password string `form:"password" json:"password" binding:"required,min=6,max=16"`
}
