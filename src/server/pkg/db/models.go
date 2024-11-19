// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type ActivityActivitytype string

const (
	ActivityActivitytypeProject  ActivityActivitytype = "Project"
	ActivityActivitytypeWorkshop ActivityActivitytype = "Workshop"
)

func (e *ActivityActivitytype) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ActivityActivitytype(s)
	case string:
		*e = ActivityActivitytype(s)
	default:
		return fmt.Errorf("unsupported scan type for ActivityActivitytype: %T", src)
	}
	return nil
}

type NullActivityActivitytype struct {
	ActivityActivitytype ActivityActivitytype `json:"activity_activitytype"`
	Valid                bool                 `json:"valid"` // Valid is true if ActivityActivitytype is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullActivityActivitytype) Scan(value interface{}) error {
	if value == nil {
		ns.ActivityActivitytype, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ActivityActivitytype.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullActivityActivitytype) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ActivityActivitytype), nil
}

type Activity struct {
	Activityid        int32                `json:"activityid"`
	Title             string               `json:"title"`
	Proposer          int32                `json:"proposer"`
	Startdate         time.Time            `json:"startdate"`
	Enddate           time.Time            `json:"enddate"`
	Maxnumber         sql.NullInt32        `json:"maxnumber"`
	Format            sql.NullString       `json:"format"`
	Description       string               `json:"description"`
	Proposedatetime   time.Time            `json:"proposedatetime"`
	Acceptadmin       sql.NullInt32        `json:"acceptadmin"`
	Acceptdatetime    sql.NullTime         `json:"acceptdatetime"`
	Applicationstatus sql.NullString       `json:"applicationstatus"`
	Activitytype      ActivityActivitytype `json:"activitytype"`
}

type Activityregistration struct {
	Activityid  int32     `json:"activityid"`
	Memberid    int32     `json:"memberid"`
	Role        string    `json:"role"`
	Expectation string    `json:"expectation"`
	Datetime    time.Time `json:"datetime"`
}

type Activityrole struct {
	Activityid   int32  `json:"activityid"`
	Activityrole string `json:"activityrole"`
}

type Admin struct {
	Adminid int32  `json:"adminid"`
	Email   string `json:"email"`
}

type Chatdevad struct {
	Messageid   int32     `json:"messageid"`
	Adminid     int32     `json:"adminid"`
	Developerid int32     `json:"developerid"`
	Message     string    `json:"message"`
	Datetime    time.Time `json:"datetime"`
}

type Developer struct {
	Developerid int32  `json:"developerid"`
	Email       string `json:"email"`
}

type Feedback struct {
	Feedbackid       int32     `json:"feedbackid"`
	Activityid       int32     `json:"activityid"`
	Memberid         int32     `json:"memberid"`
	Feedbackmessage  string    `json:"feedbackmessage"`
	Feedbackdatetime time.Time `json:"feedbackdatetime"`
}

type Member struct {
	Memberid       int32         `json:"memberid"`
	Fname          string        `json:"fname"`
	Lname          string        `json:"lname"`
	Email          string        `json:"email"`
	Phone          string        `json:"phone"`
	Githuburl      string        `json:"githuburl"`
	Interest       string        `json:"interest"`
	Reason         string        `json:"reason"`
	Acceptdatetime sql.NullTime  `json:"acceptdatetime"`
	Acceptadmin    sql.NullInt32 `json:"acceptadmin"`
}

type Project struct {
	Projectid int32          `json:"projectid"`
	Advisor   sql.NullString `json:"advisor"`
}

type Workshop struct {
	Workshopid int32     `json:"workshopid"`
	Starttime  time.Time `json:"starttime"`
	Endtime    time.Time `json:"endtime"`
}
