package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuha-yuha/DevMomentAPI/models"
	"github.com/yuha-yuha/DevMomentAPI/services"
)

func TestGetUserDefineModels(t *testing.T) {
	defineModels := services.GetUserDefineModels("../testdata/testdata.json")

	expected := []models.UserDefineModel{
		{Name: "user", Content: map[string]interface{}{
			"name": "ya",
			"id":   "12345",
		}},
		{Name: "product", Content: map[string]interface{}{
			"user": "${user}",
			"name": "どら焼き",
		},
		},
	}

	assert.Equal(t, expected, defineModels)

}
