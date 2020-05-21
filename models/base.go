package models

type Params struct {
	TaskID   int `json:"task_id"`
	EventID  int `json:"event_id"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
