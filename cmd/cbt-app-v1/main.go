package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/thesamueliyeh/cbt-app-v1/internals/routes"
	"github.com/thesamueliyeh/cbt-app-v1/internals/services"
	"github.com/thesamueliyeh/cbt-app-v1/internals/utils"
)

func main() {
	godotenv.Load()
	// initialise Supabase
	services.InitSupabase()

	// initialise db
	// db, err := services.InitDB()
	// if err != nil {
	// 	panic("failed to connect to database")
	// }

	// // temp delete db tables
	// err = services.DeleteTables(db)
	// if err != nil {
	// 	panic("failed to drop existing tables")
	// }

	// // automigrate models
	// err = services.AutoMigrate(db)
	// if err != nil {
	// 	panic("failed to automigrate models")
	// }

	// initialise echo
	e := echo.New()
	e.Static("/static", "static")

	// initialise validator
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	// initialise router
	routes.InitRouter(e)

	e.Logger.Fatal(e.Start(":3000"))
}
