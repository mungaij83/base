package models

import (
	"github.com/hexya-erp/hexya/src/models/types"
	"time"
)

type IntervalTypeWrapper struct {
}

func (IntervalTypeWrapper) Values() types.Selection {
	return types.Selection{
		"minutes": "Minutes",
		"hours":   "Hours",
		"days":    "Days",
		"weeks":   "Weeks",
		"months":  "Months",
	}
}

type CronModel struct {
	Name           string              `hexya:"required"`
	User           UserModel           `hexya:"required;many2one=id"`
	Active         bool                `hexya:"default=true"`
	IntervalNumber int                 `hexya:"default=1;help=Repeat every x.;goType=int"`
	IntervalType   IntervalTypeWrapper `hexya:"display_name=Interval Unit;default=months"`
	NextCall       *time.Time          `hexya:"type=datetime;display_name=Next Execution Date;required;default=date;help=Next planned execution date for this job."`
	Model          string              `hexya:"required"`
	Method         string              `hexya:"required"`
	RecordsIds     string              `hexya:"type=text;default=[];help=Use a JSON list format (e.g. [1, 2])"`
	Arguments      string              `hexya:"type=text;default=[];help=Use a JSON list format (e.g. [[1, 2], \"My string value\", true]). For relation fields, pass the ID or the list of IDs"`
}
