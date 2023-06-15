package controllers

import (
	"app/managers"
	"net/http"
)

type AdsController struct {
	advertiserManager managers.Advertisement
}

func NewAdsController(adsManager managers.Advertisement) *AdsController {
	return &AdsController{
		advertiserManager: adsManager,
	}
}

// GetUserAd returns available add for given user, country and language.
func (ac *AdsController) GetUserAd(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "userId is mandatory", http.StatusBadRequest)
		return
	}

	country := r.URL.Query().Get("country")
	if country == "" {
		http.Error(w, "country is mandatory", http.StatusBadRequest)
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		http.Error(w, "language is mandatory", http.StatusBadRequest)
		return
	}

	adResponse, err := ac.advertiserManager.GetUserAdvertisement(r.Context(),
		userFromInputParams(userId, country, language))
	if err != nil {
		http.Error(w, "something was wrong", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, http.StatusOK, toUserAdResponse(adResponse))
}
