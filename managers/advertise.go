package managers

import "context"

type Advertisement interface {
	GetUserAdvertisement(context context.Context, user User) (Advertise, error)
}

type Advertise struct {
	Url string
	Id  string
}

type User struct {
	Id       string
	Country  string
	Language string
}
