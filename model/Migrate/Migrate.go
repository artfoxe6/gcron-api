package Migrate

import (
	"gcron-api/config/db"
	"gcron-api/model/Job"
)

func Run() {
	db.Instance().AutoMigrate(
		&Job.Job{},
	)
}
