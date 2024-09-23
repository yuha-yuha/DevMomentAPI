package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuha-yuha/DevMomentAPI/controllers/handlers"
	"github.com/yuha-yuha/DevMomentAPI/models"
)

func TestCreateUserDefineHandler(t *testing.T) {
	apis := []models.UserDefineAPI{
		{Path: "/hoge", Response: map[string]interface{}{"hello": "world"}},
		{Path: "/huga", Response: map[string]interface{}{"user": "${user}"}},
	}
	defineModels := []models.UserDefineModel{
		{Name: "user", Content: map[string]interface{}{"name": "alice", "age": 18}},
	}

	defineHandlers := handlers.CreateUserDefineHandler(apis, defineModels)

	case1, _ := json.Marshal(map[string]interface{}{
		"hello": "world",
	})

	case2, _ := json.Marshal(map[string]map[string]interface{}{
		"user": {
			"name": "alice",
			"age":  18,
		},
	})

	expected := [][]byte{
		case1,
		case2,
	}
	for i, handler := range defineHandlers {

		req := httptest.NewRequest(http.MethodGet, handler.Path, nil)
		rec := httptest.NewRecorder()

		handler.HandlerFunc.ServeHTTP(rec, req)

		assert.Equal(t, string(expected[i])+"\n", rec.Body.String())
	}
}
