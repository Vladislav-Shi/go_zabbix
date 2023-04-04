package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func init() {

}

const ()

func GetProcessesRoute(c echo.Context) error {
	/*
		Ручка для получения списка процессов
		?exec=<список_исключаемых_полей_из_набора_через_запятую>
	*/
	executesParams, _ := QueryToArray(c.QueryParam("exec"))
	processes := GetProcessesInfo()
	var filteredProcesses []map[string]string
	for _, v := range processes {
		// filteredProcesses = append(filteredProcesses, v.ToMap([]string{"memoryPercent", "exe"}))
		filteredProcesses = append(filteredProcesses, v.ToMap(executesParams))
	}
	json_data, err := json.Marshal(filteredProcesses)

	if err != nil {
		fmt.Println(err)
	}
	ansver := c.JSONBlob(http.StatusOK, json_data)
	// return c.String(http.StatusOK, string(json_data))
	return ansver
}

func GetSingleProcessRoute(c echo.Context) error {
	pidStr := c.Param("pid")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		fmt.Println(err)
	}
	process := GetProcessInfo(int32(pid))
	var execute []string
	json_data, err := json.Marshal(process.ToMap(execute))
	if err != nil {
		fmt.Println(err)
	}
	ansver := c.JSONBlob(http.StatusOK, json_data)
	return ansver
}

func GetMemoryUsageRoute(c echo.Context) error {
	/*
		Вызвращает использование памяти
	*/
	json_data, err := json.Marshal(GetMemoryInfo())
	if err != nil {
		fmt.Println(err)
	}
	ansver := c.JSONBlob(http.StatusOK, json_data)
	return ansver
}

func GetCpuStatRoute(c echo.Context) error {
	/*
		Вызвращает использование цп
	*/
	json_data, err := json.Marshal(GetCpuStat())
	if err != nil {
		fmt.Println(err)
	}
	ansver := c.JSONBlob(http.StatusOK, json_data)
	return ansver
}

func GetHostInfoRoute(c echo.Context) error {
	/*
		Вызвращает использование цп
	*/
	json_data, err := json.Marshal(GetHostInfo())
	if err != nil {
		fmt.Println(err)
	}
	ansver := c.JSONBlob(http.StatusOK, json_data)
	return ansver
}

func RunBashCommandRoute(c echo.Context) error {
	u := new(BashCommand)
	if err := c.Bind(u); err != nil {
		return err
	}
	output := u.RunBashCommand()
	return c.String(http.StatusOK, output)
}

func GetBashHistoryRoute(c echo.Context) error {
	json_data, err := json.Marshal(GetBashHistory())
	if err != nil {
		fmt.Println(err)
	}
	ansver := c.JSONBlob(http.StatusOK, json_data)
	return ansver
}

func ProcessStopRoute(c echo.Context) error {
	pidStr := c.Param("pid")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		fmt.Println(err)
	}
	status := ProcessStop(int32(pid))
	if status {
		return c.String(http.StatusBadRequest, "Error")
	}
	return c.String(http.StatusOK, "Ok")
}

func ProcessResumeRoute(c echo.Context) error {
	pidStr := c.Param("pid")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		fmt.Println(err)
	}
	status := ProcessResume(int32(pid))
	if status {
		return c.String(http.StatusBadRequest, "Error")
	}
	return c.String(http.StatusOK, "Ok")
}

func ProcessKillRoute(c echo.Context) error {
	pidStr := c.Param("pid")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		fmt.Println(err)
	}
	status := ProcessKill(int32(pid))
	if status {
		return c.String(http.StatusBadRequest, "Error")
	}
	return c.String(http.StatusOK, "Ok")
}
