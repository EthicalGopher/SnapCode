package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Storage struct {
	Lang string `json:"lang"`
	Code string `json:"code"`
}

var Store Storage

func load() Storage {
	dataByte, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}
	var dataJson Storage
	if err = json.Unmarshal(dataByte, &dataJson); err != nil {
		log.Fatal(err)
	}
	return dataJson
}
func (s *Storage) Save() {
	dataByte, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	if err = os.WriteFile("data.json", dataByte, os.FileMode(os.O_CREATE|os.O_APPEND)); err != nil {
		log.Fatal(err)
	}

}
func init() {
	Store = load()
}
func main() {
	defer Store.Save()
	fmt.Println(Store)
}
