package models

type User struct {
	Id                    int     `valid:"numeric"`
	Name                  string  `valid:"alphanum,required"`
	Username              string  `valid:"alphanum,required"`
	Email                 string  `valid:"email"`
	Mobile_number         string  `valid:"alphanum,required"`
	Latitude              float64 `valid:latitude`
	Longitude             float64 `valid:longitude`
	Password              string  `valid:"duck,required"`
	Password_confirmation string  `valid:"duck,required"`
	City                  string  `valid:"alphanum"`
	Devise_token          string  `valid:"alphanum,required"`
	Ward_id               int     `valid:"numeric"`
	Type                  string  `valid:"string"`
}

type SuccessfulSignIn struct {
	Success string
	Message string
	User    User
	Session SessionDetails
}

type SessionDetails struct {
	SessionId   int
	DeviseToken string
}

// Message struct [controllers/account]
// Common for sign_up, session and password
type Message struct {
	Success string
	Message string
	User    User
}

type ErrorMessage struct {
	Success string
	Error   string
}

type LogErrorMessage struct {
	Success string
	Error   error
}
