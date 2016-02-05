package models

type Password struct {
	MobileNumber int `valid:"numeric"`
}

type ResetPassword struct {
	User_id     int    `valid:"numeric"`
	OldPassword string `valid:"duck,required"`
	NewPassword string `valid:"duck,required"`
}
