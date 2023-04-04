package main

import (
	"fmt"
	"net/http"

	routes "go_zabbix/app"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load(".env")
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		fmt.Println(err)
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "localhost"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://labstack.com", "https://localhost"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/proc", routes.GetProcessesRoute)

	e.GET("/proc/:pid", routes.GetSingleProcessRoute)
	e.GET("/proc/:pid/stop", routes.ProcessStopRoute)
	e.GET("/proc/:pid/run", routes.ProcessResumeRoute)
	e.GET("/proc/:pid/kill", routes.ProcessKillRoute)

	e.GET("/mem", routes.GetMemoryUsageRoute)
	e.GET("/cpu", routes.GetCpuStatRoute)
	e.GET("/host", routes.GetHostInfoRoute)
	e.GET("/log/bash", routes.GetBashHistoryRoute)
	e.POST("/run", routes.RunBashCommandRoute)
	e.Logger.Fatal(e.Start(":8081"))
}
