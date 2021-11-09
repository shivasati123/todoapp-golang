package main

import (
	"fmt"
	"todoapp-golang/database"
	"todoapp-golang/server"
)

func main() {
	err := database.ConnectAndMigrate("localhost", "5433", "To_Do_App", "local", "local", database.SSLModeDisable)
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")
	srv := server.SetupRoutes()
	err = srv.Run(":8080")
	if err != nil {
		panic(err)
	}
}