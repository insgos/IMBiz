package users

import "time"

type (
	User struct {
		UserId       string    `json:"USER_ID"`
		Pw           string    `json:"PW"`
		Name         string    `json:"NAME"`
		Email        string    `json:"EMAIL"`
		Number       string    `json:"NUMBER"`
		DepartmentId string    `json:"DEPARTMENT_ID"`
		Position     string    `json:"POSITION"`
		LastAccessDt time.Time `json:"LAST_ACCESS_DT"`
		UpdateDt     time.Time `json:"UPDATE_DT"`
		CreateDt     time.Time `json:"CREATE_DT"`
	}

	Department struct {
		Id   string `json:"ID"`
		Name string `json:"NAME"`
	}
)
