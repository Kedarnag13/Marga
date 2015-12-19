package models

type Mypoints struct {
	SenderId  int `valid:"numeric,required"`
	ReciverId int `valid:"numeric,required"`
}

type ProfileErrorMessage struct {
	Success string
	Error   string
}

type ShowErrorMessage struct {
	Success string
	Error   string
}

type SuccessMessage struct {
	Success string
	Message string
}
