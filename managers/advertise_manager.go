package managers

import (
	"app/config"
	"app/datasources"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type advertiseManager struct {
	httpClient  datasources.HTTPClient
	redisClient datasources.Cache
}

func NewAdvertiserManager(httpClient datasources.HTTPClient, redisClient datasources.Cache) Advertisement {
	return &advertiseManager{
		httpClient:  httpClient,
		redisClient: redisClient,
	}
}

// GetUserAdvertisement will follow next flow in order to retrieve an advertisement for the given user
//  1. Call the endpoint to get Ads list
//  2. Filter Ads retrieved by endpoint
//     a. By start and end hour
//     b. By language
//     c. By country
//     d. By not shown before in the current calendar day
//  4. Async save key userId-addId in redis cache with a TTL that will expire at 00:00:00 next day (UTC)
func (am *advertiseManager) GetUserAdvertisement(context context.Context, user User) (ad Advertise, err error) {
	ads, err := am.getAdsFromService(context)
	if err != nil {
		return
	}

	ads = ads.filterByTimeLanguageAndCountry(user.Language, user.Country)
	advertisement := am.getFirstAdNotShownInTheDay(user.Id, ads)

	if advertisement != nil {
		go am.setAdCache(user.Id, advertisement.Id)
		return advertisement.toAdvertisementOutput(), nil
	}

	return Advertise{}, nil
}

func (am *advertiseManager) getFirstAdNotShownInTheDay(userID string, inputAds adsResponse) *advertiseItem {
	for _, ad := range inputAds.Ads {
		key := fmt.Sprintf("%s-%s", userID, ad.Id)
		exists, err := am.redisClient.KeyExists(key)
		if err != nil {
			fmt.Errorf("error getting item from redis %s, %v", key, err)
		}
		if !exists {
			return &ad
		}
	}

	return nil
}

func (am *advertiseManager) getAdsFromService(context context.Context) (ads adsResponse, err error) {
	requestInfo := &datasources.RequestInfo{
		Url:        config.ServiceConfig.Ads.URL,
		HTTPMethod: http.MethodGet,
	}

	response, err := am.httpClient.Request(context, requestInfo)
	if err != nil {
		log.Printf("error getting Ads %v", err)
		return
	}

	err = json.Unmarshal(response.Body, &ads)
	if err != nil {
		log.Printf("error decoding Ads %v", err)
		return
	}

	return
}

func (am *advertiseManager) setAdCache(userID string, adID string) {
	key := fmt.Sprintf("%s-%s", userID, adID)
	am.redisClient.Set(key, true, durationToUTCMidnight())
}

func durationToUTCMidnight() time.Duration {
	now := time.Now().UTC()
	midnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)
	duration := midnight.Sub(now)
	return duration
}
