package models

import (
	"github.com/hexya-erp/hexya/src/models/loader"
	"github.com/hexya-erp/hexya/src/models/types"
	"time"
)

type QueueJobStates struct {
}

func (QueueJobStates) Values() map[string]string {
	return types.Selection{
		"pending":  "Pending",
		"enqueued": "Enqueued",
		"started":  "Started",
		"done":     "Done",
		"failed":   "Failed",
	}
}

type QueueChannelModel struct {
	Name     string `hexya:"required;index;unique"`
	Capacity int    `hexya:"required;default=1;goType=int"`
}

type QueueJobModel struct {
	Name              string            `hexya:"required,index"`
	Model             string            `hexya:"required"`
	Method            string            `hexya:"required"`
	RecordsIds        string            `hexya:"type=text;default=[];help=Use a JSON list format (e.g. [1, 2])"`
	Arguments         string            `hexya:"default=[];type=text;help=Use a JSON list format (e.g. [[1, 2], \"My string value\", true]) For relation fields, pass the ID or the list of IDs"`
	User              UserModel         `hexya:"many2one=id;required"`
	Company           CompanyModel      `hexya:"required;many2one=id"`
	Channel           QueueChannelModel `hexya:"many2one=id"`
	Priority          int
	ExecuteAfterJob   *QueueChannelModel `hexya:"many2one=id;display_name=Execute only after;help=Execute the current job only after this one has been correctly executed"`
	ExecuteBeforeJobs *QueueJobModel     `hexya:"one2many=id;reverseFk=ExecuteAfterJob;help=List of jobs that will be executed after the current one"`
	State             QueueJobStates     `hexya:"required;index;readOnly;default=pending"`
	ExcInfo           string             `hexya:"type=text;display_name=Exception Info;readOnly"`
	Result            string             `hexya:"type=text;readOnly"`
	DateStarted       *time.Time         `hexya:"type=datetime;readOnly"`
	DateEnqueued      *time.Time         `hexya:"type=datetime;readOnly"`
	DateDone          *time.Time         `hexya:"type=datetime;readOnly"`
	ETA               *time.Time         `hexya:"type=datetime;display_name=Execute only after"`
	Retry             int                `hexya:"display_name=Current trye;default=30"`
	MaxRetries        int                `hexya:"help=The job will fail if the number of tries reach the max. retries.Retries are infinite when equals zero."`
}

func InitJobQueues() {
	loader.NewTypedModel(QueueChannelModel{})
	loader.NewTypedModel(QueueJobModel{})
	loader.NewTypedModel(CronModel{})
}
