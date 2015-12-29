package models

type Comment struct {
	User_id     int    `valid:"numeric"`
	Issue_id    int    `valid:"numeric"`
	Description string `valid:"duck,required"`
}

type SuccessCommentMessage struct {
	Success string
	Message string
}

type CommentDetails struct {
	Comment_message string `valid:"alphanum"`
	User_id         int    `valid:"numeric"`
	Name            string `valid:"alphanum"`
}

type CommentList struct {
	Success         string
	No_of_comments  int
	Comment_details []CommentDetails
}
