package models

type Task struct {
	ID           int
	Name         string
	EventID      int
	OperateType  int
	StartTS      int64
	EndTS        int64
	CurrentIndex int64
	CurrentCount int
	TotalCount   int
	CreateTS     int64
}
