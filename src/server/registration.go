import (
	"fmt",
	"database/sql",
	"net/http",
	"html/template",
	"log",

	_ "github.com/go-sql-driver/mysql"
)

type pjRegist struct {
	memberID int(11),
	activityID int(11),
	role varchar(30),
	datetime datetime,
	expectation text,
	status varchar(30)
}

type wsRegist struct {
	memberID int(11),
	activityID int(11), // change to wsID
	role varchar(30),
	datetime datetime,
	expectation text,
	status varchar(30)
}

func postPjRegist(w http.ResponseWriter, r *http.Request) {
	// Get project registration information from the request
	pjRegist := pjRegist{
		memberID: r.FormValue("memberID"),
		activityID: r.FormValue("activityID"),
		role: r.FormValue("role"),
		datetime: r.FormValue("datetime"),
		expectation: r.FormValue("expectation"),
		status: r.FormValue("status")
	}
}

func postWsRegist(w http.ResponseWriter, r *http.Request) {
	// Get workshop registration information from the request
	wsRegist := wsRegist{
		memberID: r.FormValue("memberID"),
		activityID: r.FormValue("activityID"),
		role: r.FormValue("role"),
		datetime: r.FormValue("datetime"),
		expectation: r.FormValue("expectation"),
		status: r.FormValue("status")
	}
}