package models

type Comment struct {
	User_id     int    `valid:"numeric"`
	Issue_id    int    `valid:"alpha"`
	Description string `valid:"duck,required"`
}

type SuccessCommentMessage struct {
	Success string
	Message string
}
