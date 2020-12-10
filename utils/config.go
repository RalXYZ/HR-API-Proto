package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Login struct {
	DbUser   string `json:"dbUser"`
	DbPasswd string `json:"dbPasswd"`
}

func GetLoginInfo(file string) (string, string) {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error: Requires a database configuration file!")
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var login Login
	json.Unmarshal([]byte(byteValue), &login)

	return login.DbUser, login.DbPasswd
}
