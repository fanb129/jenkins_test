package form

type EditForm struct {
	Role uint `form:"role" json:"role" binding:"required,oneof=1 2 3"`
}
