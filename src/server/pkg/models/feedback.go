package models

import "time"

type Feedback struct {
	Feedbackid       int32     `json:"feedbackId"`
	Activityid       int32     `json:"activityId"`
	Memberid         int32     `json:"memberId"`
	Feedbackmessage  string    `json:"feedbackMessage"`
	Feedbackdatetime time.Time `json:"feedbackDateTime"`
}
