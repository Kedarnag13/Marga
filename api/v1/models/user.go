package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name                 string
	Email                string
	MobileNumber         string
	Latitude             float64
	Longitude            float64
	Password             string
	PasswordConfirmation string
	City                 string
	DeviseToken          string
	Type                 string
}

type Session struct {
	gorm.Model
	UserId      uint
	DeviseToken string
	User        User
	Device      Device
}

type Device struct {
	gorm.Model
	UserId    uint
	SessionId uint
}

type Response struct {
	Message string
	Error   string
	Success bool
}

// type User struct {
// 	Id                    int     `valid:"numeric"`
// 	Name                  string  `valid:"alphanum,required"`
// 	Username              string  `valid:"alphanum,required"`
// 	Email                 string  `valid:"email"`
// 	Mobile_number         string  `valid:"alphanum,required"`
// 	Latitude              float64 `valid:latitude`
// 	Longitude             float64 `valid:longitude`
// 	Password              string  `valid:"duck,required"`
// 	Password_confirmation string  `valid:"duck,required"`
// 	City                  string  `valid:"alphanum"`
// 	Devise_token          string  `valid:"alphanum,required"`
// 	Type                  string  `valid:"string"`
// }

// type SuccessfulSignIn struct {
// 	Success string
// 	Message string
// 	User    User
// 	Session SessionDetails
// }

// type SessionDetails struct {
// 	SessionId   int
// 	DeviseToken string
// }

// // Message struct [controllers/account]
// // Common for sign_up, session and password
// type Message struct {
// 	Success string
// 	Message string
// 	User    User
// }

// type ErrorMessage struct {
// 	Success string
// 	Error   string
// }

// type LogErrorMessage struct {
// 	Success string
// 	Error   error
// }
