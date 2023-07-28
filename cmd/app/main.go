package main

import (
	"fmt"
	"main/core/config"
	"main/core/database"
	"main/core/routes"
	"main/core/utils"
	"net/http"
)

func main() {
	configs := config.GetConfig()
	db := database.GetDatabase(configs.DatabaseURI)
	repository := database.NewAdapter(db, "products")

	utils.Info("waiting for the service to start...", nil)

	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouters(repository)
	utils.Info("service is running on port:", port)

	err := http.ListenAndServe(port, router)
	utils.Panic("failed to start the server:", err)
}
