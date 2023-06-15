package controllers

import "app/managers"

type UserAd struct {
	Id  string `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
}

func toUserAdResponse(advertise managers.Advertise) UserAd {
	return UserAd{
		Id:  advertise.Id,
		Url: advertise.Url,
	}
}

func userFromInputParams(userId string, country string, language string) managers.User {
	return managers.User{
		Id:       userId,
		Country:  country,
		Language: language,
	}
}
