package main

import (
	"log"
	"net/http"

	_ "database/sql"

	"github.com/daffarmd/gofun/helper"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/app"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/controller"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/exception"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/middleware"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/repository"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/service"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.GetConnection()
	validate := validator.New()
	dataRepository := repository.NewDataRepository()
	dataService := service.NewDataService(dataRepository, db, validate)
	dataController := controller.NewDataController(dataService)

	router := httprouter.New()

	router.GET("/api/data", dataController.FindAll)
	router.GET("/api/data/:dataId", dataController.FindById)
	router.POST("/api/data", dataController.Create)
	router.PUT("/api/data/:dataId", dataController.Update)
	router.DELETE("/api/data/:dataId", dataController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	log.Println("Running at http://localhost:3000")

	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}
