package job

import (
	"encoding/json"
	"gcron-api/config/cache"
	"gcron-api/config/env"
	"gcron-api/model/Job"
	"gcron-api/util/lib"
	"gcron-api/util/request"
	"github.com/artfoxe6/cron_expression"
	"time"
)

func Add(r *request.Request) bool {
	err := r.Validate([]string{
		"name", "expression", "method", "url", "args", "location_name", "location_offset", "ttl",
	})
	if err != nil {
		return r.Error(err.Error())
	}
	inputs := r.Posts()
	job := Job.Job{
		Name:           inputs["name"],
		Expression:     inputs["expression"],
		Method:         inputs["method"],
		Url:            inputs["url"],
		Args:           inputs["args"],
		Status:         1,
		Header:         inputs["header"],
		LocationName:   inputs["location_name"],
		LocationOffset: lib.Int(inputs["location_offset"]),
		TTL:            0,
	}
	err = job.Add()
	//同步到redis中
	SyncRedis(job)
	if err != nil {
		return r.Error(err.Error())
	}
	return r.Success(nil)
}

//同步到redis中
func SyncRedis(job Job.Job) {
	expr := cron_expression.NewExpression(job.Expression, job.LocationName, job.LocationOffset)
	dst := make([]string, 0)
	_ = expr.Next(time.Now(), 1, &dst)
	local := time.FixedZone(job.LocationName, job.LocationOffset)
	nextAt, _ := time.ParseInLocation("2006-01-02 15:04:05", dst[0], local)
	_, _ = cache.Instance().Do("ZADD", env.Redis.Zset, nextAt.Unix(), job.ID)
	jobByte, _ := json.Marshal(job)
	_, _ = cache.Instance().Do("HSET", env.Redis.Hash, job.ID, string(jobByte))
}
