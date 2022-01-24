package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	// Check configuration file existance
	if CheckConfigFile() {

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

	return nil
}

func CheckConfigFile() bool {
	if CheckConfigDirectory() {

		filename := "conf-prod.json"

		// Open the file, this will append to the today's file if server restarted.
		file, err := os.OpenFile("config/"+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {

			panic(err)
		}
		defer file.Close()

		File, _ := file.Stat()
		if File.Size() == 0 {
			fmt.Println("Please, fill config file. Located at -> config/config-prod.json")

			os.Exit(1)

			return false
		}
	}

	return true
}

func CheckConfigDirectory() bool {
	path := "config"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)

			os.Exit(1)

			return false
		}
	}

	return true
}
