package main

import (
	"fmt"
	"go-layered/config"
	"go-layered/repository"
	"go-layered/controllers"
	"go-layered/routes"
)

func main() {
	cfg, err := config.Load()
	if (err != nil) {
		fmt.Println(err)
		return	
	}
	
	db, err := config.Connect(cfg.DatabaseURL)
	if (err != nil) {
		fmt.Println(err)
		return	
	}
	
	repo := repository.NewUserRepository(db)
	ctrl := controllers.NewUserController(repo)
	router := routes.NewUserRouter(ctrl)
	
	router.Run(":" + cfg.Port)
}
