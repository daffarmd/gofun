package jsonfun_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type UserResponse struct {
	Name string `json:"name_user"`
	Age  int    `json:"age"`
}

func TestJsonTag(t *testing.T) {
	user := UserResponse{Name: "Eko", Age: 12}
	b, _ := json.Marshal(user)
	fmt.Println(string(b))
}

func TestJsonDecode(t *testing.T) {
	userRequest := `{"name_user":"Eko","age":12}`
	user := UserResponse{}
	jsonBytes := []byte(userRequest)
	json.Unmarshal(jsonBytes, &user)
	fmt.Println(user)
}
