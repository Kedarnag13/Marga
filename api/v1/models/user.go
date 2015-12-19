package models

type User struct {
	Id                    int     `valid:"numeric"`
	Name                  string  `valid:"alphanum,required"`
	Username              string  `valid:"alphanum,required"`
	Email                 string  `valid:"email"`
	Mobile_number         string  `valid:"alphanum,required"`
	Latitude              float64 `valid:latitude`
	Longitude             float64 `valid:longitude`
	Password              string  `valid:"alphanum,required"`
	Password_confirmation string  `valid:"alphanum,required"`
	City                  string  `valid:"alphanum"`
	Devise_token          string  `valid:"alphanum,required"`
	Ward_id               int     `valid:"numeric"`
	Type                  string  `valid:"duck"`
}

type SignIn struct {
	Success string
	Message string
	User    User
	Session Session
}

// Session struct [account/session]
type Session struct {
	SessionId int
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
