package form

import "time"

type NsAddForm struct {
	Uid         uint       `form:"u_id" json:"u_id" binding:"gte=0"`
	Name        string     `form:"name" json:"name" binding:"required,min=3,max=16"`
	ExpiredTime *time.Time `form:"expired_time" json:"expired_time"`
	Cpu         string     `json:"cpu" form:"cpu" binding:"required"`
	Memory      string     `json:"memory" form:"memory" binding:"required"`
	Num         int        `json:"num" form:"num" binding:"required"`
}
