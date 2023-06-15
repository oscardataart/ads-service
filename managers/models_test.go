package managers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidTimeWindow(t *testing.T) {
	// Test valid time window from 10 to 15
	assert.True(t, validTimeWindow(12, 10, 15))

	// Test valid time window from 22 to 3 (crossing midnight)
	assert.False(t, validTimeWindow(12, 22, 3))

	// Test valid time window from 5 to 2 (crossing midnight)
	assert.True(t, validTimeWindow(12, 5, 2))

	// Test for 0 hour
	assert.True(t, validTimeWindow(0, 5, 2))
	assert.True(t, validTimeWindow(0, 23, 1))

	// Test invalid time window from 18 to 2 (crossing midnight)
	assert.False(t, validTimeWindow(12, 18, 2))

	// Test invalid time window with invalid hours
	assert.False(t, validTimeWindow(12, 25, 30))
	assert.False(t, validTimeWindow(12, 24, 10))
	assert.False(t, validTimeWindow(12, 0, 24))
	assert.False(t, validTimeWindow(12, -1, 5))
	assert.False(t, validTimeWindow(12, 10, -1))
}

func TestAdsResponse_FilterByTimeLanguageAndCountry(t *testing.T) {
	ads := adsResponse{
		Ads: []advertiseItem{
			{Country: "US", Language: "en", StartHour: 0, EndHour: 23},
			{Country: "US", Language: "es", StartHour: 0, EndHour: 23},
			{Country: "UK", Language: "en", StartHour: 0, EndHour: 23},
		},
	}

	// Test filtering by valid time, language, and country
	filteredAds := ads.filterByTimeLanguageAndCountry("en", "US")
	expectedAds := adsResponse{
		Ads: []advertiseItem{
			{Country: "US", Language: "en", StartHour: 0, EndHour: 23},
		},
	}
	assert.Equal(t, expectedAds, filteredAds)

	// Test filtering by valid time, language, and different country
	filteredAds = ads.filterByTimeLanguageAndCountry("es", "UK")
	expectedAds = adsResponse{}
	assert.Equal(t, expectedAds, filteredAds)

	// Test filtering by invalid language
	filteredAds = ads.filterByTimeLanguageAndCountry("fr", "US")
	expectedAds = adsResponse{}
	assert.Equal(t, expectedAds, filteredAds)
}
