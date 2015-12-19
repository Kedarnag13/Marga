package models

type Issue struct {
	Id          int     `valid:"numeric"`
	Name        string  `valid:"duck,required"`
	Type        string  `valid:"alpha"`
	Description string  `valid:"duck,required"`
	Latitude    float64 `valid:latitude,required`
	Longitude   float64 `valid:longitude,required`
	Image       string  `valid:"alphanum,required"`
	Status      bool    `valid:"required"`
	Address     string  `valid:"duck,required"`
	User_id     int     `valid:"numeric,required"`
}

type IssueDetails struct {
	Name        string  `valid:"duck,required"`
	Type        string  `valid:"alpha"`
	Description string  `valid:"duck,required"`
	Latitude    float64 `valid:latitude,required`
	Longitude   float64 `valid:longitude,required`
	Image       string  `valid:"alphanum,required"`
	Status      bool    `valid:"required"`
	Address     string  `valid:"duck,required"`
	User_id     int     `valid:"numeric,required"`
}

type IssueErrorMessage struct {
	Success string
	Error   string
}

type SuccessfulCreateIssue struct {
	Success string
	Message string
	Issue   Issue
}

type IssueList struct {
	Success       string
	No_Of_Issues  int
	Issue_Details []IssueDetails
}
