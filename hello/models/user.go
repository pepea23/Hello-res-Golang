package models

type User struct {
	Name    string `json:"name"`
	SurName string `json:"sur_name"`
}

func NewParamsFromData(params map[string]interface{}) *User {
	user := new(User)
	if v, ok := params["name"]; ok {
		user.Name = v.(string)
	}

	if v, ok := params["surname"]; ok {
		user.SurName = v.(string)
	}

	return user
}
