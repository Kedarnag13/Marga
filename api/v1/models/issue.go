package models

import (
	"time"
)

type Issue struct {
	Id            int     `valid:"numeric"`
	Name          string  `valid:"duck,required"`
	Type          string  `valid:"alpha"`
	Description   string  `valid:"duck,required"`
	Latitude      float64 `valid:latitude,required`
	Longitude     float64 `valid:longitude,required`
	Image         string  `valid:"alphanum,required"`
	Status        bool    `valid:"required"`
	Address       string  `valid:"duck,required"`
	User_id       int     `valid:"numeric,required"`
	Corporator_id int     `valid:"numeric,required"`
	Created_at    time.Time
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

type WardDetails struct {
	Id           int
	Name         string
	Email        string
	Devise_token string
}

type WardList struct {
	Success      string
	No_Of_Wards  int
	Ward_Details []WardDetails
}

type Notification struct {
	Id          int
	Message     string
	Sender_id   int
	Reciever_id int
}

type NotificationSuccess struct {
	Success string
	Message string
}

type NotificationError struct {
	Success string
	Error   string
}
