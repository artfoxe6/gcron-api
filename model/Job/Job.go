package Job

import (
	"gcron-api/config/db"
	"gcron-api/model"
)

type Job struct {
	model.Base
	Name           string `json:"name"`
	Expression     string `json:"expression"`
	Method         string `json:"method"`
	Url            string `json:"url"`
	Args           string `json:"args"`
	Status         int    `json:"status"`
	Header         string `json:"header"`
	LocationName   string `json:"location_name"`
	LocationOffset int    `json:"location_offset"`
	TTL            int    `json:"ttl"`
}

func (Job) TableName() string {
	return "job"
}

func (job *Job) Add() error {

	if err := db.Instance().Create(job).Error; err != nil {
		return err
	}
	return nil
}
