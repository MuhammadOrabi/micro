//go:generate goagen bootstrap -d micro/design

package main

import (
	"database/sql"
	"fmt"
	"micro/app"
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	// DBName Database name
	DBName = "micro"
	// DBUser Database user name
	DBUser = "micro"
	// DBPassword Database password
	DBPassword = "micro"
	// DBPort Port for database connection
	DBPort = "5432"
	// DBHost Hostname for database or IP
	DBHost = "db"
	// DBSSLMode SSL mode for connection
	DBSSLMode = "disable"

	//SERVER PORT
	listenPort = "3000"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	resetVariables()

	// Create service
	service := goa.New("adder")

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", DBHost, DBPort, DBUser, DBPassword, DBName, DBSSLMode)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "operands" controller
	c := NewOperandsController(service, db)
	app.MountOperandsController(service, c)

	// Start service
	if err := service.ListenAndServe(":3000"); err != nil {
		service.LogError("startup", "err", err)
	}

}

func resetVariables() {
	if len(os.Getenv("DB_NAME")) > 0 {
		DBName = os.Getenv("DB_NAME")
	}

	if len(os.Getenv("DB_USER")) > 0 {
		DBUser = os.Getenv("DB_USER")
	}

	if len(os.Getenv("DB_PASSWORD")) > 0 {
		DBPassword = os.Getenv("DB_PASSWORD")
	}

	if len(os.Getenv("DB_PORT")) > 0 {
		DBPort = os.Getenv("DB_PORT")
	}

	if len(os.Getenv("DB_HOST")) > 0 {
		os.Getenv("DB_HOST")
	}

	if len(os.Getenv("DB_SSL_MODE")) > 0 {
		DBSSLMode = os.Getenv("DB_SSL_MODE")
	}
	if len(os.Getenv("PORT")) > 0 {
		listenPort = os.Getenv("PORT")
	}
}
