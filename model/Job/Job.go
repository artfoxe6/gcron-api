package Job

import "gcron-api/model"

type Job struct {
	model.Base
	name string
	expression string
	method string
	status int
}
