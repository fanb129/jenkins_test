package form

import "time"

// BatchSparkAddForm 支持批量添加
type BatchSparkAddForm struct {
	Uid            []uint     `form:"u_id" json:"u_id" binding:"required"`
	MasterReplicas int32      `form:"master_replicas" json:"master_replicas" binding:"required,gte=1,lte=3"`
	WorkerReplicas int32      `form:"worker_replicas" json:"worker_replicas" binding:"required,gte=2,lte=10"`
	ExpiredTime    *time.Time `form:"expired_time" json:"expired_time"`
	Cpu            string     `json:"cpu" form:"cpu" binding:"required"`
	Memory         string     `json:"memory" form:"memory" binding:"required"`
}

type BatchLinuxAddForm struct {
	Uid         []uint     `form:"u_id" json:"u_id" binding:"required"`
	Kind        uint       `form:"kind" json:"kind" binding:"required,gt=0"`
	ExpiredTime *time.Time `form:"expired_time" json:"expired_time"`
	Cpu         string     `json:"cpu" form:"cpu" binding:"required"`
	Memory      string     `json:"memory" form:"memory" binding:"required"`
}

// BatchHadoopAddForm 批量添加
type BatchHadoopAddForm struct {
	Uid                []uint     `form:"u_id" json:"u_id" binding:"required"`
	HdfsMasterReplicas int32      `form:"hdfs_master_replicas" json:"hdfs_master_replicas" binding:"required,gte=1,lte=3"`
	DatanodeReplicas   int32      `form:"datanode_replicas" json:"datanode_replicas" binding:"required,gte=2,lte=10"`
	YarnMasterReplicas int32      `form:"yarn_master_replicas" json:"yarn_master_replicas" binding:"required,gte=1,lte=3"`
	YarnNodeReplicas   int32      `form:"yarn_node_replicas" json:"yarn_node_replicas" binding:"required,gte=2,lte=10"`
	ExpiredTime        *time.Time `form:"expired_time" json:"expired_time"`
	Cpu                string     `json:"cpu" form:"cpu" binding:"required"`
	Memory             string     `json:"memory" form:"memory" binding:"required"`
}
