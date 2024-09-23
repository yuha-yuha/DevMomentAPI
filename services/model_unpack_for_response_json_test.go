package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuha-yuha/DevMomentAPI/models"
	"github.com/yuha-yuha/DevMomentAPI/services"
)

func TestModelUnpackforResponseJson(t *testing.T) {
	apis := []models.UserDefineAPI{
		{Path: "/hoge", Response: map[string]interface{}{"hello": "world"}},
		{Path: "/huga", Response: map[string]interface{}{"user": "${user}"}},
	}
	defineModels := []models.UserDefineModel{
		{Name: "user", Content: map[string]interface{}{"name": "alice", "age": 18}},
	}

	expected := []*models.UserDefineAPI{
		{Path: "/hoge", Response: map[string]interface{}{"hello": "world"}},
		{Path: "/huga", Response: map[string]interface{}{"user": map[string]interface{}{
			"name": "alice", "age": 18,
		}}},
	}

	apiPointers := []*models.UserDefineAPI{}

	for _, api := range apis {
		apiPointers = append(apiPointers, &api)
	}
	services.ModelUnpackforResponseJson(apiPointers, defineModels)

	assert.Equal(t, expected, apiPointers)
}
