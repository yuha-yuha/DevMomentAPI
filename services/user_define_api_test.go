package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuha-yuha/DevMomentAPI/models"
	"github.com/yuha-yuha/DevMomentAPI/services"
)

func TestGetUserDefineAPIs(t *testing.T) {
	apis := services.GetUserDefineAPIs("../testdata/testdata.json")

	expected := []models.UserDefineAPI{
		{Path: "/huga", Response: map[string]interface{}{
			"current_user": "${user}",
			"newProduct":   "${product}",
			"count":        123.0,
		}},
		{Path: "/hoge", Response: map[string]interface{}{
			"aaaa": "ssss",
		}},
	}

	assert.Equal(t, expected, apis)

}
