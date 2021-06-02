package middlewere

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Form(c echo.Context) error {
	postForm, err := c.FormParams()
	var data map[string]interface{}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
	}
	if len(postForm) > 0 {
		jsonData := getParamsFromJsonData(postForm)
		if jsonData == nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"message": "Json parsing error"})
		}
		data = jsonData.(map[string]interface{})
	} else {
		return nil
	}
	c.Set("params", data)
	return nil
}

func getParamsFromJsonData(postForm map[string][]string) interface{} {
	var data interface{}
	var dataString = postForm["data"][0]
	json.Unmarshal([]byte(dataString), &data)
	return data
}
