package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

func main() {
	db := db.NewDB()
	gachaValidator := validator.NewGachaValidator()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	gachaRepository := repository.NewGachaRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	gachaUsecase := usecase.NewGachaUsecase(gachaRepository, gachaValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	gachaController := controller.NewGachaController(gachaUsecase)
	e := router.NewRouter(userController, taskController, gachaController)
	e.Logger.Fatal(e.Start(":8080"))
}
