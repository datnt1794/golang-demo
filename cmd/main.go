package main

import (
	"bookstore/db"
	"bookstore/model"
	"bookstore/router"
	"github.com/labstack/echo"
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"strconv"
)

func init() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
	pgHost := os.Getenv("POSTGRES_HOST")
	pgUserName := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDbName := os.Getenv("POSTGRES_DB")
	pgPort := os.Getenv("POSTGRES_PORT")
	pgPortNumber, _ := strconv.Atoi(pgPort)
	fmt.Printf("%s uses %s\n", pgHost, pgUserName)
	sql := &db.Sql{
		Host:     pgHost,
		Port:     pgPortNumber,
		UserName: pgUserName,
		Password: pgPassword,
		DbName:   pgDbName,
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		fmt.Printf("Server start")
		return c.JSON(200, model.Response{
			StatusCode: 200,
			Message:    "Home Page",
		})
	})

	router.UserRouter(e, sql)
	router.CateRouter(e, sql)
	router.ProductRouter(e, sql)
	router.OrderRouter(e, sql)

	e.Logger.Fatal(e.Start(":8000"))
}
