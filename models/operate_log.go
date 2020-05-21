package models

type OperateLog struct {
	ID            int64
	UID           int64
	EventID       int
	Operate       int
	OperateTS     int64
	OperateStatus int8
	Data          string
}
