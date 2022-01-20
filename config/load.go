package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type JSON struct {
	Key       string `json:"key"`
	Hash      string `json:"hash"`
	MasterUrl string `json:"master-url"`
}

type Config interface {
	Load()
}

func (receiver *JSON) Load() *JSON {
	// TODO : link to cli interface
	JsonFile, ErrReadJsonFile := ioutil.ReadFile("config/conf-prod.json")
	if ErrReadJsonFile != nil {
		log.Fatalf("JSON JSON struct ErrReadJsonFile: " + ErrReadJsonFile.Error())
	}

	ErrUnmarshalJsonFile := json.Unmarshal(JsonFile, receiver)
	if ErrUnmarshalJsonFile != nil {
		log.Fatalf("JSON JSON struct ErrUnmarshalJsonFile: " + ErrUnmarshalJsonFile.Error())
	}

	return receiver
}
