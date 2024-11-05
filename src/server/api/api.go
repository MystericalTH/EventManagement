package api

type Member struct {
	MemberID int    `json:"memberID"`
	FName    string `json:"fName"`
}

type Error struct {
	Message string `json:"message"`
}
