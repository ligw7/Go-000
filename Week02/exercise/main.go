package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	api := NewAPI()
	var studentId uint64
	// userId = 19, not found
	studentId = 19
	response := api.GetUserInfo(studentId)
	jsonObj, _ := json.MarshalIndent(response, "", "    ")
	fmt.Printf("\r\nunnormal responsed json:\r\n %s", string(jsonObj))
	// userId = 21, should be ok
	studentId = 21
	response = api.GetUserInfo(studentId)
	jsonObj, _ = json.MarshalIndent(response, "", "    ")
	fmt.Printf("\r\nnormal responsed json:\r\n %s", string(jsonObj))
}
