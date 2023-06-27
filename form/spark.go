package form

import "time"

type SparkAddForm struct {
	Uid            uint       `form:"u_id" json:"u_id" binding:"required,gt=0"`
	MasterReplicas int32      `form:"master_replicas" json:"master_replicas" binding:"required,gte=1,lte=3"`
	WorkerReplicas int32      `form:"worker_replicas" json:"worker_replicas" binding:"required,gte=2,lte=10"`
	ExpiredTime    *time.Time `form:"expired_time" json:"expired_time"`
	Cpu            string     `json:"cpu" form:"cpu" binding:"required"`
	Memory         string     `json:"memory" form:"memory" binding:"required"`
}

type SparkUpdateForm struct {
	Name           string     `json:"name" form:"name" binding:"required"`
	Uid            uint       `form:"u_id" json:"u_id" binding:"required,gt=0"`
	MasterReplicas int32      `form:"master_replicas" json:"master_replicas" binding:"required,gte=1,lte=3"`
	WorkerReplicas int32      `form:"worker_replicas" json:"worker_replicas" binding:"required,gte=2,lte=10"`
	ExpiredTime    *time.Time `form:"expired_time" json:"expired_time"`
	Cpu            string     `json:"cpu" form:"cpu" binding:"required"`
	Memory         string     `json:"memory" form:"memory" binding:"required"`
}
