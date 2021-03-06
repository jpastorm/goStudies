package main

import (
	"encoding/json"
	"fmt"
)

type MongoDbChangeEvent struct {
	Id     string `json:"id"`
	Source string `json:"source"`
	// context   string
	// database  string
	// field     string
	// operation string
	// updatedBy string
	// oldValue  string
	// newValue  string
}

func main() {
	//	var doc MongoDbChangeEvent
	//	jsonString := `{"id":"5f9bdc94de3af92965dc4f55","source":"mongodb","context":null,"database":"test.test","field":"hello","operation":"u","timestamp":1604050097,"updatedBy":null,"oldValue":"World4","newValue":"World5"}`
	//sample := MongoDbChangeEvent{id: "asd", source: "asd"}
	var sample MongoDbChangeEvent
	sample.Id = "123"
	sample.Source = "12243"
	sampleJson, _ := json.Marshal(sample)
	fmt.Println(string(sampleJson))
	// err := json.Unmarshal([]byte(sampleJson), &doc)
	// if err != nil {
	// 	return
	// }
	// newJson, _ := json.Marshal(doc)
	// fmt.Println(string(newJson))

}
