package managers

import (
	"app/datasources"
	"app/datasources/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestAdvertiseManager_GetUserAdvertisement(t *testing.T) {
	mockHTTPClient := &mocks.HTTPClient{}
	mockRedisClient := &mocks.Cache{}

	manager := NewAdvertiserManager(mockHTTPClient, mockRedisClient)

	user := User{
		Id:       "testUserID",
		Country:  "us",
		Language: "eng",
	}

	mockResponse := []byte(`{
    "ads": [
        {
            "id": "59d4fb16",
            "video_url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
            "country": "us",
            "lang": "eng",
            "start_hour": 0,
            "end_hour": 23
        },
        {
            "id": "f75d5ea7",
            "video_url": "https://www.youtube.com/watch?v=QH2-TGUlwu4",
            "country": "us",
            "lang": "eng",
            "start_hour": 0,
            "end_hour": 23
        }
	]}`)

	mockHTTPClient.On("Request", mock.Anything, mock.Anything).
		Return(&datasources.ResponseInfo{Body: mockResponse}, nil).Once()

	mockRedisClient.On("KeyExists", mock.Anything).Return(false, nil)
	mockRedisClient.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ad, err := manager.GetUserAdvertisement(context.Background(), user)
	time.Sleep(1 * time.Second)
	assert.NoError(t, err)

	expectedAd := Advertise{Id: "59d4fb16", Url: "https://www.youtube.com/watch?v=dQw4w9WgXcQ"}
	assert.Equal(t, expectedAd, ad)

	mockHTTPClient.AssertExpectations(t)
	mockRedisClient.AssertExpectations(t)
}

func TestAdvertiseManager_GetUserAdvertisement_ErrorGettingAds(t *testing.T) {
	mockHTTPClient := &mocks.HTTPClient{}
	mockRedisClient := &mocks.Cache{}

	manager := NewAdvertiserManager(mockHTTPClient, mockRedisClient)

	user := User{
		Id:       "testUserID",
		Country:  "us",
		Language: "eng",
	}

	mockHTTPClient.On("Request", mock.Anything, mock.Anything).
		Return(nil, errors.New("some error")).Once()

	ad, err := manager.GetUserAdvertisement(context.Background(), user)
	time.Sleep(1 * time.Second)
	assert.Error(t, err)

	expectedAd := Advertise{}
	assert.Equal(t, expectedAd, ad)

	mockHTTPClient.AssertExpectations(t)
	mockRedisClient.AssertExpectations(t)
}

func TestAdvertiseManager_GetUserAdvertisement_ErrorDecodingAds(t *testing.T) {
	mockHTTPClient := &mocks.HTTPClient{}
	mockRedisClient := &mocks.Cache{}

	manager := NewAdvertiserManager(mockHTTPClient, mockRedisClient)

	user := User{
		Id:       "testUserID",
		Country:  "us",
		Language: "eng",
	}

	mockResponse := []byte(`wrong json`)

	mockHTTPClient.On("Request", mock.Anything, mock.Anything).
		Return(&datasources.ResponseInfo{Body: mockResponse}, nil).Once()

	ad, err := manager.GetUserAdvertisement(context.Background(), user)
	time.Sleep(1 * time.Second)
	assert.Error(t, err)

	expectedAd := Advertise{}
	assert.Equal(t, expectedAd, ad)

	mockHTTPClient.AssertExpectations(t)
}
