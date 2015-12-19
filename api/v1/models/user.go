package models

type User struct {
	Id                    int     `valid:"numeric"`
	Name                  string  `valid:"alphanum,required"`
	Username              string  `valid:"alphanum,required"`
	Email                 string  `valid:"email"`
	Mobile_number         int     `valid:"numeric,required"`
	Latitude              float64 `valid:latitude`
	Longitude             float64 `valid:longitude`
	Password              string  `valid:"alphanum,required"`
	Password_confirmation string  `valid:"alphanum,required"`
	City                  string  `valid:"alphanum"`
	Devise_token          string  `valid:"alphanum,required"`
	Ward_id               int
	Type                  string `valid:"duck"`
}
