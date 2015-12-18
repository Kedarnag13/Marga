package models

type Issue struct {
	Id          int     `valid:"numeric"`
	Name        string  `valid:"alphanum,required"`
	Type        string  `valid:"alpha"`
	Description string  `valid:"duck,required"`
	Latitude    float64 `valid:latitude,required`
	Longitude   float64 `valid:longitude,required`
	Image       string  `valid:"alphanum,required"`
	Status      bool    `valid:"required"`
}
