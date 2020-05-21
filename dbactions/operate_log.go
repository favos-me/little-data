package dbactions

import "github.com/favos-me/little-data/models"

func CreateOperateLog(log *models.OperateLog) error {
	return db.Create(log).Error
}

func CreateOperateLogs(logs []models.OperateLog) error {
	return nil
}

func QueryOperateLogs(params *models.Params) ([]models.OperateLog, error) {
	return nil, nil
}
