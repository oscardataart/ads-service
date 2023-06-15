package controllers

import (
	"app/managers"
	"app/managers/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdsController_GetUserAd(t *testing.T) {
	mockAdsManager := &mocks.Advertisement{}
	adsController := NewAdsController(mockAdsManager)
	mockUrl := "http://example.com/ad"
	mockId := "mockAdId"
	mgrAd := managers.Advertise{
		Url: mockUrl,
		Id:  mockId,
	}

	mockAdsManager.On("GetUserAdvertisement", mock.Anything,
		userFromInputParams("123", "US", "en")).
		Return(mgrAd, nil).Once()

	req, err := http.NewRequest(http.MethodGet, "/user-ad?userId=123&country=US&language=en", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	adsController.GetUserAd(rr, req)

	expectedResponse := `{"id": "` + mockId + `", "url": "` + mockUrl + `"}`
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, expectedResponse, rr.Body.String())
}

func TestAdsController_GetUserAd_NoUserIDError(t *testing.T) {
	mockAdsManager := &mocks.Advertisement{}
	adsController := NewAdsController(mockAdsManager)

	req, err := http.NewRequest(http.MethodGet, "/user-ad?country=US&language=en", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	adsController.GetUserAd(rr, req)

	expectedResponse := "userId is mandatory\n"
	assert.Equal(t, expectedResponse, rr.Body.String())
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAdsController_GetUserAd_NoCountryError(t *testing.T) {
	mockAdsManager := &mocks.Advertisement{}
	adsController := NewAdsController(mockAdsManager)

	req, err := http.NewRequest(http.MethodGet, "/user-ad?userId=123&language=en", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	adsController.GetUserAd(rr, req)

	expectedResponse := "country is mandatory\n"
	assert.Equal(t, expectedResponse, rr.Body.String())
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAdsController_GetUserAd_NoLanguageError(t *testing.T) {
	mockAdsManager := &mocks.Advertisement{}
	adsController := NewAdsController(mockAdsManager)

	req, err := http.NewRequest(http.MethodGet, "/user-ad?country=US&userId=123", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	adsController.GetUserAd(rr, req)

	expectedResponse := "language is mandatory\n"
	assert.Equal(t, expectedResponse, rr.Body.String())
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAdsController_GetUserAd_GetUserAdvertisementError(t *testing.T) {
	mockAdsManager := &mocks.Advertisement{}
	adsController := NewAdsController(mockAdsManager)

	mockAdsManager.On("GetUserAdvertisement", mock.Anything, mock.Anything).
		Return(managers.Advertise{}, errors.New("some")).Once()

	req, err := http.NewRequest(http.MethodGet, "/user-ad?userId=123&country=US&language=en", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	adsController.GetUserAd(rr, req)

	expectedResponse := "something was wrong\n"
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, expectedResponse, rr.Body.String())
}
