package services

import (
	"regexp"

	"github.com/yuha-yuha/DevMomentAPI/models"
)

func ModelUnpackforResponseJson(apis []*models.UserDefineAPI, models []models.UserDefineModel) {

	for i, model := range models {
		for contentKey, contentValue := range model.Content {
			ValueIsMap(&contentValue, models)
			models[i].Content[contentKey] = contentValue
		}
	}

	for _, api := range apis {
		ValueIsMap(&(api.Response), models)
	}
}

func ValueIsMap(v *interface{}, models []models.UserDefineModel) {
	if m, ok := (*v).(map[string]interface{}); ok {
		for k, v := range m {
			ValueIsMap(&v, models)
			m[k] = v
		}
	} else {
		pattern := `\$\{[^}]*\}`
		re, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}

		if str, ok := (*v).(string); ok {
			if re.MatchString(str) {
				modelStr := str[2 : len(str)-1]
				for _, model := range models {
					if model.Name == modelStr {
						*v = model.Content
					}
				}
			}
		}
	}
}
