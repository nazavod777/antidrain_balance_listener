package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"main/custom_types"
	"os"
)

func ReadConfig() custom_types.ConfigStruct {
	var configData custom_types.ConfigStruct

	jsonFile, err := os.Open("data/config.json")

	if err != nil {
		log.Panicf("Error When Opening Config File: %s", err.Error())
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Panicf("Error When Reading Config File: %s", err.Error())
	}

	err = json.Unmarshal(byteValue, &configData)

	if err != nil {
		log.Panicf("Error When Decoding Config File: %s", err.Error())
	}

	return configData
}
