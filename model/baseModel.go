package model

type ConfigDept struct {
	ID   uint64 `xorm:"depno autoincr pk"`
	Name string `xorm:"depname not null"`
	Addr string `xorm:"addr varchar(50)"`
	Tag  int    `xorm:"tag"`
}

func (ConfigDept) TableName() string {
	return "dept"
}
