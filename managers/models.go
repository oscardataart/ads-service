package managers

import "time"

type advertiseItem struct {
	Id        string `json:"id,omitempty"`
	VideoUrl  string `json:"video_url,omitempty"`
	Country   string `json:"country,omitempty"`
	Language  string `json:"lang,omitempty"`
	StartHour int    `json:"start_hour,omitempty"`
	EndHour   int    `json:"end_hour,omitempty"`
}

type adsResponse struct {
	Ads []advertiseItem `json:"ads,omitempty"`
}

func (ar adsResponse) filterByTimeLanguageAndCountry(language string, country string) adsResponse {
	var filteredAds []advertiseItem
	currentHour := time.Now().UTC().Hour()
	for _, ad := range ar.Ads {
		if ad.Country == country && ad.Language == language &&
			validTimeWindow(currentHour, ad.StartHour, ad.EndHour) {
			filteredAds = append(filteredAds, ad)
		}
	}
	return adsResponse{Ads: filteredAds}
}

func validTimeWindow(currentHour, startHour, endHour int) bool {
	// for range 23 - 2: valid hours are 23, 24, 1
	// for range 2 - 5: valid hours are 2, 3, 4
	if endHour >= 24 || endHour < 0 || startHour >= 24 || startHour < 0 {
		return false
	}
	if endHour <= startHour {
		if currentHour >= startHour || currentHour < endHour {
			return true
		}
	} else if currentHour >= startHour && currentHour < endHour {
		return true
	}

	return false
}

func (item *advertiseItem) toAdvertisementOutput() Advertise {
	return Advertise{
		Url: item.VideoUrl,
		Id:  item.Id,
	}
}
