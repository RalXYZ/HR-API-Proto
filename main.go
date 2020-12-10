package main

import (
	"HR-API-proto/database"
	"HR-API-proto/http"
)

func main() {
	database.Init()
	http.InitWebFramework()
	http.StartServer()
}
