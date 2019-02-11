package main

import (
	"database/sql"
	"fmt"
	dbinfo_repo "github.com/goodrain/go-demo/dbinfo/repository"
	"github.com/goodrain/go-demo/middleware"
	"github.com/labstack/echo"
	"net/url"
	"os"
)

func main() {
	dbuser := os.Getenv("MYSQL_USER")
	dbpw := os.Getenv("MYSQL_PASSWORD")
	dbhost := os.Getenv("MYSQL_HOST")
	dbport := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DATABASE")
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpw, dbhost, dbport, dbname)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Shanghai")
	dsn := fmt.Sprintf("%s?%s", conn, val.Encode())
	dbconn, _ := sql.Open(`mysql`, dsn)
	defer dbconn.Close()

	_ = dbinfo_repo.NewMysqlDBInfoRepository(dbconn)

	e := echo.New()

	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	e.Static("/", "public")

	e.Logger.Fatal(e.Start(":5000"))
}
