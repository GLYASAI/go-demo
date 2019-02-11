package main

import (
	"database/sql"
	"fmt"
	dbinfo_http "github.com/goodrain/go-demo/dbinfo/delivery/http"
	dbinfo_repo "github.com/goodrain/go-demo/dbinfo/repository"
	dbinfo_ucase "github.com/goodrain/go-demo/dbinfo/usecase"
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

	dbinfoRepo := dbinfo_repo.NewMysqlDBInfoRepository(dbconn)

	e := echo.New()

	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	e.Static("/", "public")

	dbinfoUcaser := dbinfo_ucase.NewDBInfoUsecase(dbinfoRepo)
	dbinfo_http.NewDBInfoHTTPHandler(e, dbinfoUcaser)

	e.Logger.Fatal(e.Start(":5000"))
}
